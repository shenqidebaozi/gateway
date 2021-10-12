package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"

	config "github.com/go-kratos/gateway/api/gateway/config/v1"
	"github.com/go-kratos/gateway/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
)

// Client is a proxy client.
type Client interface {
	Invoke(ctx context.Context, req *http.Request, opts ...CallOption) (middleware.Response, error)
}

type clientImpl struct {
	selector selector.Selector
	nodes    atomic.Value
	log      *log.Helper
}

func (c *clientImpl) Invoke(ctx context.Context, r *http.Request, opts ...CallOption) (middleware.Response, error) {
	callInfo := defaultCallInfo()
	for _, o := range opts {
		if err := o.before(&callInfo); err != nil {
			return nil, err
		}
	}
	selected, done, err := c.selector.Select(ctx, selector.WithFilter(callInfo.filters...))
	if err != nil {
		return nil, err
	}

	defer done(ctx, selector.DoneInfo{Err: err})
	node := c.nodes.Load().(map[string]*node)[selected.Address()]
	/*req, err := http.NewRequest(r.Method(), fmt.Sprintf("http://%s%s", selected.Address(), r.Endpoint().String()), r.Body())
	if err != nil {
		return nil, err
	}
	c.log.Infof("client invoke %s", req.URL.String())*/
	r.URL.Scheme = "http"
	r.URL.Host = selected.Address()
	r.Host = selected.Address()
	r.RequestURI = ""
	c.log.Infof("client invoke %s", r.URL.String())
	resp, err := node.client.Do(r)
	if err != nil {
		return nil, err
	}
	return &Response{protocol: node.protocol, Response: resp}, nil
}

// NewFactory new a client factory.
func NewFactory(logger log.Logger, r registry.Discovery) func(endpoint *config.Endpoint) (Client, error) {
	log := log.NewHelper(logger)
	return func(endpoint *config.Endpoint) (Client, error) {
		c := &clientImpl{
			selector: wrr.New(),
			log:      log,
		}
		var nodes []selector.Node
		atomicNodes := make(map[string]*node, 0)
		for _, backend := range endpoint.Backends {
			target, err := parseTarget(backend.Target)
			if err != nil {
				return nil, err
			}
			switch target.Scheme {
			case "direct":
				node := buildNode(backend.Target, endpoint.Protocol, backend.Weight, endpoint.Timeout.AsDuration())
				nodes = append(nodes, node)
				atomicNodes[backend.Target] = node
			case "discovery":
				w, err := r.Watch(context.Background(), target.Endpoint)
				if err != nil {
					return nil, err
				}
				go func() {
					// TODO: goroutine leak
					// only one backend configuration allowed when using service discovery
					for {
						services, err := w.Next()
						if err != nil && errors.Is(err, context.Canceled) {
							return
						}
						if len(services) == 0 {
							continue
						}
						var nodes []selector.Node
						atomicNodes := make(map[string]*node, 0)
						for _, ser := range services {
							scheme := strings.ToLower(endpoint.Protocol.String())
							addr, err := parseEndpoint(ser.Endpoints, scheme, false)
							if err != nil {
								log.Errorf("failed to parse endpoint: %v", err)
								continue
							}
							node := buildNode(addr, endpoint.Protocol, backend.Weight, endpoint.Timeout.AsDuration())
							nodes = append(nodes, node)
							atomicNodes[addr] = node
						}
						c.selector.Apply(nodes)
						c.nodes.Store(atomicNodes)
					}
				}()
			default:
				return nil, fmt.Errorf("unknown scheme: %s", target.Scheme)
			}
		}
		c.selector.Apply(nodes)
		c.nodes.Store(atomicNodes)
		return c, nil
	}
}

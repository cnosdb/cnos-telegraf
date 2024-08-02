//go:generate ../../../tools/readme_config_includer/generator
package cnosdb_subscription

import (
	_ "embed"
	"fmt"
	"net"
	"sync"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/config"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb"
	"github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription/cnosdb/generated/kv_service"
	"google.golang.org/grpc"
)

func init() {
	inputs.Add("cnosdb_subscription", func() telegraf.Input {
		return &CnosDbSubscription{
			ServiceAddress: ":8803",
		}
	})
}

//go:embed sample.conf
var sampleConfig string

type CnosDbSubscription struct {
	ServiceAddress string          `toml:"service_address"`
	ServiceVersion string          `toml:"service_version"`
	Timeout        config.Duration `toml:"timeout"`

	Log telegraf.Logger `toml:"-"`

	wg sync.WaitGroup `toml:"-"`

	listener   net.Listener `toml:"-"`
	grpcServer *grpc.Server `toml:"-"`
}

func (*CnosDbSubscription) SampleConfig() string {
	return sampleConfig
}

func (c *CnosDbSubscription) Init() error {
	c.Log.Info("Initialization completed.")
	return nil
}

func (c *CnosDbSubscription) Gather(_ telegraf.Accumulator) error {
	return nil
}

func (c *CnosDbSubscription) Start(acc telegraf.Accumulator) error {
	c.grpcServer = grpc.NewServer(grpc.MaxRecvMsgSize(10 * 1024 * 1024))
	kv_service.RegisterTSKVServiceServer(c.grpcServer, cnosdb.NewTSKVService(acc))

	if c.listener == nil {
		listener, err := net.Listen("tcp", c.ServiceAddress)
		if err != nil {
			return err
		}
		c.listener = listener
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		if err := c.grpcServer.Serve(c.listener); err != nil {
			acc.AddError(fmt.Errorf("failed to stop CnosDbSubscription gRPC service: %w", err))
		}
	}()

	c.Log.Infof("Listening on %s", c.listener.Addr().String())

	return nil
}

func (c *CnosDbSubscription) Stop() {
	if c.grpcServer != nil {
		c.grpcServer.Stop()
	}
	c.wg.Wait()
}

func (c *CnosDbSubscription) MarkHighPriority() {
	// Do nothing
}

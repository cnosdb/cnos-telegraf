//go:generate ../../../tools/config_includer/generator
//go:generate ../../../tools/readme_config_includer/generator
package socket_listener

import (
	_ "embed"
	"fmt"
	"net"
	"sync"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal"
	"github.com/influxdata/telegraf/plugins/common/socket"
	"github.com/influxdata/telegraf/plugins/inputs"
)

//go:embed sample.conf
var sampleConfig string

var once sync.Once

type SocketListener struct {
	ServiceAddress string          `toml:"service_address"`
	Log            telegraf.Logger `toml:"-"`
	socket.Config
	socket.SplitConfig

	socket *socket.Socket
	parser telegraf.Parser
}

func (*SocketListener) SampleConfig() string {
	return sampleConfig
}

func (sl *SocketListener) Init() error {
	sock, err := sl.Config.NewSocket(sl.ServiceAddress, &sl.SplitConfig, sl.Log)
	if err != nil {
		return err
	}
	sl.socket = sock

	return nil
}

func (sl *SocketListener) Gather(_ telegraf.Accumulator) error {
	return nil
}

func (sl *SocketListener) SetParser(parser telegraf.Parser) {
	sl.parser = parser
}

func (sl *SocketListener) Start(acc telegraf.Accumulator) error {
	// Create the callbacks for parsing the data and recording issues
	onData := func(_ net.Addr, data []byte) {
		metrics, err := sl.parser.Parse(data)
		if err != nil {
			acc.AddError(err)
			return
		}
		if len(metrics) == 0 {
			once.Do(func() {
				sl.Log.Debug(internal.NoMetricsCreatedMsg)
			})
		}
		switch ac := acc.(type) {
		case telegraf.HighPriorityAccumulator:
			for _, m := range metrics {
				if err = ac.AddMetricHighPriority(m); err != nil {
					sl.Log.Tracef("input socket_listener got error from high-priority-IO")
					acc.AddError(fmt.Errorf("writing data to output failed: %w", err))
				}
			}
		default:
			for _, m := range metrics {
				ac.AddMetric(m)
			}
		}
	}
	onError := func(err error) {
		acc.AddError(err)
	}

	// Start the listener
	if err := sl.socket.Setup(); err != nil {
		return err
	}
	sl.socket.Listen(onData, onError)
	addr := sl.socket.Address()
	sl.Log.Infof("Listening on %s://%s", addr.Network(), addr.String())

	return nil
}

func (sl *SocketListener) Stop() {
	if sl.socket != nil {
		sl.socket.Close()
	}
}

func (sl *SocketListener) MarkHighPriority() {
	// Do nothing
}

func init() {
	inputs.Add("socket_listener", func() telegraf.Input {
		return &SocketListener{}
	})
}

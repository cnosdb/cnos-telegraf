//go:build !custom || inputs || inputs.bcache

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/cnosdb_subscription" // register plugin

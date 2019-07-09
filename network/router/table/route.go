package table

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var (
	// DefaultLink is default network link
	DefaultLink = "local"
	// DefaultLocalMetric is default route cost metric for the local network
	DefaultLocalMetric = 1
	// DefaultNetworkMetric is default route cost metric for the micro network
	DefaultNetworkMetric = 10
)

// Route is network route
type Route struct {
	// Service is destination service name
	Service string
	// Address is service node address
	Address string
	// Gateway is route gateway
	Gateway string
	// Network is network address
	Network string
	// Link is network link
	Link string
	// Metric is the route cost metric
	Metric int
}

// Hash returns route hash sum.
func (r *Route) Hash() uint64 {
	h := fnv.New64()
	h.Reset()
	h.Write([]byte(r.Service + r.Address + r.Gateway + r.Network + r.Link))

	return h.Sum64()
}

// String returns human readable route
func (r Route) String() string {
	// this will help us build routing table string
	sb := &strings.Builder{}

	// create nice table printing structure
	table := tablewriter.NewWriter(sb)
	table.SetHeader([]string{"Service", "Address", "Gateway", "Network", "Link", "Metric"})

	strRoute := []string{
		r.Service,
		r.Address,
		r.Gateway,
		r.Network,
		r.Link,
		fmt.Sprintf("%d", r.Metric),
	}
	table.Append(strRoute)

	// render table into sb
	table.Render()

	return sb.String()
}

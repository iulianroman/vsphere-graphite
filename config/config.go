package config

import (
	"github.com/cblomart/vsphere-graphite/backend"
	"github.com/cblomart/vsphere-graphite/vsphere"
)

// Configuration : configuration base
type Configuration struct {
	VCenters     []*vsphere.VCenter
	Metrics      []*vsphere.Metric
	Interval     int
	Domain       string
	Properties   []string
	Backend      *backend.Config
	CPUProfiling bool
	MEMProfiling bool
	FlushSize    int
	ReplacePoint bool
	// VCenterResultLimit is the maximum amount of results to fetch back in one query
	VCenterResultLimit int
	// VCenterInstanceRatio is the number of effective result in function of the metrics.
	// This is necessary due to the possibility to retrieve instances with wildcards
	VCenterInstanceRatio float64
}

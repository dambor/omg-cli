package cloudfoundry

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/plugins/products/cf-mysql/enaml-gen/proxy"
	"github.com/pivotal-golang/yaml"
)

//NewMySQLProxyPartition -
func NewMySQLProxyPartition(c *cli.Context) (igf InstanceGrouper, err error) {

	igf = &MySQLProxy{
		AZs:          c.StringSlice("az"),
		StemcellName: c.String("stemcell-name"),
		NetworkIPs:   c.StringSlice("mysql-proxy-ip"),
		NetworkName:  c.String("mysql-proxy-network"),
		VMTypeName:   c.String("mysql-proxy-vm-type"),
		ApiUsername:  c.String("mysql-proxy-api-username"),
		ApiPassword:  c.String("mysql-proxy-api-password"),
		ExternalHost: c.String("mysql-proxy-external-host"),
		ClusterIPs:   c.StringSlice("mysql-ip"),
		SyslogAggregator: &proxy.SyslogAggregator{
			Address:   c.String("syslog-address"),
			Port:      c.Int("syslog-port"),
			Transport: c.String("syslog-transport"),
		},
		Nats: &proxy.Nats{
			User:     c.String("nats-user"),
			Password: c.String("nats-pass"),
			Machines: c.StringSlice("nats-machine-ip"),
			Port:     4222,
		},
	}
	if !igf.HasValidValues() {
		b, _ := yaml.Marshal(igf)
		err = fmt.Errorf("invalid values in MySQLProxy: %v", string(b))
		igf = nil
	}
	return
}

//ToInstanceGroup -
func (s *MySQLProxy) ToInstanceGroup() (ig *enaml.InstanceGroup) {
	ig = &enaml.InstanceGroup{
		Name:      "mysql_proxy-partition",
		Instances: len(s.NetworkIPs),
		VMType:    s.VMTypeName,
		AZs:       s.AZs,
		Stemcell:  s.StemcellName,
		Jobs: []enaml.InstanceJob{
			s.newMySQLProxyJob(),
		},
		Networks: []enaml.Network{
			enaml.Network{Name: s.NetworkName, StaticIPs: s.NetworkIPs},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
	}
	return
}

func (s *MySQLProxy) newMySQLProxyJob() enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "proxy",
		Release: "cf-mysql",
		Properties: &proxy.Proxy{
			ApiUsername:      s.ApiUsername,
			ApiPassword:      s.ApiPassword,
			ExternalHost:     s.ExternalHost,
			ClusterIps:       s.ClusterIPs,
			SyslogAggregator: s.SyslogAggregator,
			Nats:             s.Nats,
		},
	}
}

//HasValidValues -
func (s *MySQLProxy) HasValidValues() bool {
	return (len(s.AZs) > 0 &&
		s.StemcellName != "" &&
		s.VMTypeName != "" &&
		s.NetworkName != "" &&
		len(s.NetworkIPs) > 0 &&
		s.ExternalHost != "" &&
		s.ApiPassword != "" &&
		s.ApiUsername != "" &&
		len(s.ClusterIPs) > 0)
}

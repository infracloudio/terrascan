package config

import (
	"github.com/awslabs/goformation/v4/cloudformation/elasticloadbalancing"
)

// ElasticLoadBalancingLoadBalancerConfig holds config for aws_elb
type ElasticLoadBalancingLoadBalancerConfig struct {
	Config
	Listeners           interface{} `json:"listener"`
	AccessLoggingPolicy interface{} `json:"access_logs,omitempty"`
}

type ELBAccessLoggingPolicyConfig struct {
	Enabled bool `json:"enabled"`
}

type ELBListenerConfig struct {
	LBProtocol       string `json:"lb_protocol"`
	InstanceProtocol string `json:"instance_protocol"`
}

// GetElasticLoadBalancingLoadBalancerConfig returns config for aws_elb
func GetElasticLoadBalancingLoadBalancerConfig(e *elasticloadbalancing.LoadBalancer) []AWSResourceConfig {
	cf := ElasticLoadBalancingLoadBalancerConfig{
		Config: Config{
			Name: e.LoadBalancerName,
			Tags: e.Tags,
		},
	}

	if e.AccessLoggingPolicy != nil {
		cf.AccessLoggingPolicy = ELBAccessLoggingPolicyConfig{
			Enabled: e.AccessLoggingPolicy.Enabled,
		}
	}

	if e.Listeners != nil {
		lc := make([]ELBListenerConfig, 0)
		for _, listener := range e.Listeners {
			lc = append(lc, ELBListenerConfig{
				InstanceProtocol: listener.InstanceProtocol,
				LBProtocol:       listener.Protocol,
			})
		}
		cf.Listeners = lc
	}

	return []AWSResourceConfig{{Resource: cf}}
}

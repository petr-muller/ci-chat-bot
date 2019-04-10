package apiv1

import (
	"k8s.io/apimachinery/pkg/labels"
)

func ProwSpecForPeriodicConfig(config *Periodic, decorationConfig *DecorationConfig) *ProwJobSpec {
	spec := &ProwJobSpec{
		Type:  PeriodicJob,
		Job:   config.Name,
		Agent: KubernetesAgent,

		Refs: &Refs{},

		PodSpec: config.Spec.DeepCopy(),
	}

	if decorationConfig != nil {
		spec.DecorationConfig = decorationConfig.DeepCopy()
	} else {
		spec.DecorationConfig = &DecorationConfig{}
	}
	isTrue := true
	spec.DecorationConfig.SkipCloning = &isTrue
	return spec
}

func HasProwJob(config *Config, name string) (*Periodic, bool) {
	for i := range config.Periodics {
		if config.Periodics[i].Name == name {
			return &config.Periodics[i], true
		}
	}
	return nil, false
}

func HasProwJobWithLabels(config *Config, selector labels.Selector) (*Periodic, bool) {
	for i := range config.Periodics {
		if selector.Matches(labels.Set(config.Periodics[i].Labels)) {
			return &config.Periodics[i], true
		}
	}
	return nil, false
}

package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"kusionstack.io/kusion/pkg/apis/project"
	"kusionstack.io/kusion/pkg/apis/stack"
	appmodel "kusionstack.io/kusion/pkg/modules/inputs"
	"kusionstack.io/kusion/pkg/modules/inputs/workload"
	"kusionstack.io/kusion/pkg/modules/inputs/workload/network"
)

func TestAppsConfigBuilder_Build(t *testing.T) {
	p, s := buildMockProjectAndStack()
	appName, app := buildMockApp()
	acg := &AppsConfigBuilder{
		Apps: map[string]appmodel.AppConfiguration{
			appName: *app,
		},
	}

	spec, err := acg.Build(&Options{}, p, s)
	assert.NoError(t, err)
	assert.NotNil(t, spec)
}

func buildMockApp() (string, *appmodel.AppConfiguration) {
	return "app1", &appmodel.AppConfiguration{
		Workload: &workload.Workload{
			Header: workload.Header{
				Type: "Service",
			},
			Service: &workload.Service{
				Base: workload.Base{},
				Type: "Deployment",
				Ports: []network.Port{
					{
						Type:     network.CSPAliyun,
						Port:     80,
						Protocol: "TCP",
						Public:   true,
					},
				},
			},
		},
	}
}

func buildMockProjectAndStack() (*project.Project, *stack.Stack) {
	p := &project.Project{
		ProjectConfiguration: project.ProjectConfiguration{
			Name: "test-project",
		},
	}

	s := &stack.Stack{
		Configuration: stack.Configuration{
			Name: "test-stack",
		},
	}

	return p, s
}
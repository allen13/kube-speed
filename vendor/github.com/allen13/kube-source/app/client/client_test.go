package client

import (
	"testing"
	"k8s.io/client-go/pkg/api/v1"
	"io/ioutil"
)

func TestKubeSourceClient(t *testing.T) {
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		t.Error(err)
	}

	client, err := NewClientWithToken("integration-containers", string(token))
	if err != nil{
		t.Error(err)
	}

	ports := []v1.ServicePort{
		{
			Name: "6379-tcp",
			Protocol: "tcp",
			Port: int32(6379),
		},
	}

	createRequest := ContainerCreateRequest{
		DockerImage: "redis:alpine",
		Ports: ports,
	}

	containerResource, err := client.CreateContainerResource(createRequest)
	if err != nil {
		t.Error(err)
	}

	nodePort := containerResource.Ports[0]
	if !(nodePort > 30000 && nodePort < 32767){
		t.Errorf("Failed to get back valid not port. Got %d", nodePort)
	}

	err = client.DeleteContainerResource(containerResource.Name)
	if err != nil {
		t.Error(err)
	}
}

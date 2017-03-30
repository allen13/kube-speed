package kubernetesjob

import (
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	batchV1 "k8s.io/client-go/pkg/apis/batch/v1"
	"k8s.io/client-go/rest"
)

type Client struct {
	clientset *kubernetes.Clientset
	namespace string
}

type ContainerCreateRequest struct {
	DockerImage string           `json:"image"`
	Env         []v1.EnvVar      `json:"env"`
	Ports       []v1.ServicePort `json:"ports"`
}

type ContainerResponse struct {
	Name  string           `json:"name"`
	Ip    string           `json:"ip"`
	Ports []v1.ServicePort `json:"ports"`
}

func New() (client *Client, err error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return
	}

	if err != nil {
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	namespace, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		return
	}
	client = &Client{clientset: clientset, namespace: string(namespace)}

	return
}

func (c *Client) CreateKubeSpeedJob(requestID string, completionURL string, kubeSpeedImage string) error {
	kubeSpeedJob := &batchV1.Job{
		ObjectMeta: v1.ObjectMeta{
			Name: requestID,
			Labels: map[string]string{
				"name": requestID,
			},
		},
		Spec: batchV1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    requestID,
							Image:   kubeSpeedImage,
							Command: []string{"./kube-speed", "job", requestID, "--completion-url", completionURL},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}

	_, err := c.clientset.BatchV1Client.Jobs(c.namespace).Create(kubeSpeedJob)
	return err
}

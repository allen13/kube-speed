package server

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/pborman/uuid"
	"github.com/allen13/kube-speed/pkg/kubernetesjob"
)

//Request that captures container start time
type Request struct {
	ID             string
	ContainerStart time.Time
}

//Start the kube-speed server
func Start(completionURL string, kubeSpeedImage string) {
	ticker := time.NewTicker(time.Second * 1)

	requests := map[string]time.Time{}
	kubeClient, err := kubernetesjob.New()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(time.Second * 2)
		for range ticker.C {
			requestID := uuid.New()
			requests[requestID] = time.Now()
			kubeClient.CreateKubeSpeedJob(requestID, completionURL, kubeSpeedImage)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	e := echo.New()

	e.POST("/request", func(c echo.Context) error {
		var request Request
		c.Bind(&request)

		requestSent := requests[request.ID]
		delete(requests, request.ID)

		requestDuration := time.Since(requestSent)
		timeToContainerStartDuration := request.ContainerStart.Sub(requestSent)

		fmt.Printf("request took %v\ncontainer start took %v\n\n", requestDuration, timeToContainerStartDuration)
		return nil
	})

	e.Start(":1595")
}

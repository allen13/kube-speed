package job

import (
	"bytes"
	"encoding/json"
	"github.com/allen13/kube-speed/pkg/server"
	"log"
	"net/http"
	"time"
)

func Run(completionURL string, requestId string) {
	startTime := time.Now()
	request := server.Request{
		ID:             requestId,
		ContainerStart: startTime,
	}

	jsonValue, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	_, err = http.Post(completionURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
}

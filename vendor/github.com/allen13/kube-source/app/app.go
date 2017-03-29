package app

import (
	"github.com/allen13/kube-source/app/client"
	"github.com/allen13/kube-source/app/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunServer() (err error) {
	config.Load()

	server := buildServer()

	bindAddr := config.Get("address")

	if config.Get("tls_enabled") == "true" {
		err = server.RunTLS(bindAddr, config.Get("tls_cert"), config.Get("tls_key"))
	} else {
		err = server.Run(bindAddr)
	}

	return
}

func buildServer() (g *gin.Engine) {
	g = gin.Default()
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	g.POST("/resource", createContainerResource)
	g.DELETE("/resource/:name", deleteContainerResource)

	return
}

func createContainerResource(ctx *gin.Context) {
	kubeSourceClient, err := buildClientFromRequest(ctx.Request)
	if err != nil {
		failedResponse(ctx, err)
		return
	}

	var containerCreateRequest client.ContainerCreateRequest
	err = ctx.Bind(&containerCreateRequest)
	if err != nil {
		failedResponse(ctx, err)
		return
	}

	containerResponse, err := kubeSourceClient.CreateContainerResource(&containerCreateRequest)
	if err != nil {
		failedResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, containerResponse)
}

func deleteContainerResource(ctx *gin.Context) {
	kubeSourceClient, err := buildClientFromRequest(ctx.Request)
	if err != nil {
		return
	}

	name := ctx.Params.ByName("name")

	err = kubeSourceClient.DeleteContainerResource(name)
	if err != nil {
		failedResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"delete": "success"})
}

func failedResponse(ctx *gin.Context, err error){
	ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
}

func buildClientFromRequest(request *http.Request) (kubeSourceClient *client.Client, err error) {
	bearerToken := request.Header.Get("Authorization")
	token := bearerToken[7:]

	kubeSourceClient, err = client.NewClientWithToken(config.Get("container_namespace"), token)
	if err != nil {
		return
	}

	return
}

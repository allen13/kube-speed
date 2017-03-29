package config

import (
	"github.com/spf13/viper"
)

var EnviornmentPrefix = "KUBE_SOURCE"
var conf map[string]string = map[string]string{}

func buildConfig(){
	conf["address"] = "0.0.0.0:5606"
	conf["request_logging"] = "false"
	conf["container_namespace"] = "integration-containers"
	conf["container_ip"] = "127.0.0.1"
	conf["tls_enabled"] = "false"
	conf["tls_cert"] = "/etc/kube-source/ssl/kube-source.crt"
	conf["tls_key"] = "/etc/kube-source/ssl/kube-source.key"
}

func Load()(err error) {
	viper.SetEnvPrefix(EnviornmentPrefix)
	buildConfig()

	for field,fieldDefault := range conf {
		viper.SetDefault(field, fieldDefault)
	}

	for field := range conf {
		err = viper.BindEnv(field)
		if err != nil {
			return
		}
		conf[field] = viper.GetString(field)
	}

	return
}

func Get(field string)(string){
	return conf[field]
}

package service

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Register(consulAddress string,
	consulPort string,
	advertiseAddress string,
	advertisePort string) (registar sd.Registrar) {
	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	// Service discovery domain. In this example we use Consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = consulAddress + ":" + consulPort
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	check := api.AgentServiceCheck{
		HTTP:     "http://" + advertiseAddress + ":" + advertisePort + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Basic health checks",
	}

	port, _ := strconv.Atoi(advertisePort)
	num := rand.Intn(100)
	asr := api.AgentServiceRegistration{
		ID:      "organization" + string(num),
		Name:    "organization",
		Address: advertiseAddress,
		Port:    port,
		Tags:    []string{"org"},
		Check:   &check,
	}

	registar = consulsd.NewRegistrar(client, &asr, logger)
	return
}

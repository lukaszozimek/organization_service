package service

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	out, err := exec.Command("awk", "END{print $1}",
		"/etc/hosts").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	ipAddress := strings.Replace(string(out), "\n", "", -1)
	if ipAddress == "#" {
		ipAddress = "127.0.0.1"
	}
	check := api.AgentServiceCheck{
		HTTP:     "http://" + ipAddress + advertisePort + "/api/v1/organization/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Basic health checks",
	}

	port, _ := strconv.Atoi(advertisePort[1:])
	num := rand.Intn(100)
	asr := api.AgentServiceRegistration{
		ID:      "organization" + string(num),
		Name:    "organization",
		Address: ipAddress,
		Port:    port,
		Tags:    []string{"org"},
		Check:   &check,
	}

	registar = consulsd.NewRegistrar(client, &asr, logger)
	return
}

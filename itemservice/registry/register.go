package registry

import (
	"fmt"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

type (
	ConsulClient struct {
		C         *consulapi.Client
		ServiceId string
	}
	Service struct {
		Address string
		Port    int
		ID      string
	}
)

func NewConsulClient() *ConsulClient {
	consul, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ConsulClient{consul, ""}
}

func (client *ConsulClient) RegisterService(serviceId string) error {
	reg := new(consulapi.AgentServiceRegistration)

	reg.ID = hostname()
	reg.Name = serviceId
	reg.Address = hostname()
	port, err := strconv.Atoi(os.Getenv("PORT")[1:len(os.Getenv("PORT"))])
	if err != nil {
		return err
	}
	reg.Port = port
	reg.Check = new(consulapi.AgentServiceCheck)
	reg.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", hostname(), port)
	reg.Check.Interval = "5s"
	reg.Check.Timeout = "3s"
	reg.Tags = []string{"web"}
	return client.C.Agent().ServiceRegister(reg)

}

func (client *ConsulClient) DeregisterService(serviceName string) error {
	return client.C.Agent().ServiceDeregister(serviceName)
}

func (client *ConsulClient) LookUpService(serviceId string) (*Service, error) {
	services, err := client.C.Agent().Services()
	if err != nil {
		return nil, err
	}
	for _, k := range services {
		if k.Service == serviceId {
			return &Service{k.Address, k.Port, k.ID}, nil
		}
	}
	return nil, fmt.Errorf("No service found")

}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hn
}

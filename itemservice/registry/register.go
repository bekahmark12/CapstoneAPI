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
	_, _, err := client.LookUpService(serviceId)
	if err == nil {
		serviceId = appendId(serviceId)
	}
	client.ServiceId = serviceId
	reg.ID = serviceId
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
	return client.C.Agent().ServiceRegister(reg)

}

func (client *ConsulClient) DeregisterService(serviceName string) error {
	return client.C.Agent().ServiceDeregister(serviceName)
}

func (client *ConsulClient) LookUpService(serviceId string) (string, int, error) {
	services, err := client.C.Agent().Services()
	if err != nil {
		return "", -1, err
	}
	if srvc, ok := services[serviceId]; ok {
		address := srvc.Address
		port := srvc.Port
		return address, port, nil
	}

	return "", -1, fmt.Errorf("No service found")

}

func appendId(serviceId string) string {
	lastChar := serviceId[len(serviceId)-1:]
	i, err := strconv.Atoi(lastChar)
	if err != nil {
		return fmt.Sprintf("%s%v", serviceId, 1)
	}
	return fmt.Sprintf("%s%v", serviceId, i+1)
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hn
}

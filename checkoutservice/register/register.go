package register

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

type (
	ConsulClient struct {
		c         *consulapi.Client
		ServiceId string
	}
	Service struct {
		Address string
		Port    int
		ID      string
	}
)

func NewConsulClient(serviceId string) *ConsulClient {
	consul, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ConsulClient{consul, serviceId}
}

func (client *ConsulClient) RegisterService() error {
	reg := new(consulapi.AgentServiceRegistration)
	reg.Name = client.ServiceId
	ser, err := client.LookUpService(client.ServiceId)
	if err == nil {
		client.ServiceId = appendId(ser.ID)
	}
	reg.ID = client.ServiceId
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
	return client.c.Agent().ServiceRegister(reg)

}

func (client *ConsulClient) DeregisterService() error {
	return client.c.Agent().ServiceDeregister(client.ServiceId)
}

func (client *ConsulClient) LookUpService(serviceId string) (*Service, error) {
	services, err := client.c.Agent().Services()
	if err != nil {
		return nil, err
	}
	rx := regexp.MustCompile(fmt.Sprintf("^%s[0-9]{0,}$", serviceId))
	for k := range services {
		if rx.Match([]byte(k)) {
			return &Service{services[k].Address, services[k].Port, services[k].ID}, nil
		}
	}
	return nil, fmt.Errorf("No service found")

}

func (service *Service) GetHTTP() string {
	return fmt.Sprintf("http://%s:%v/", service.Address, service.Port)
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

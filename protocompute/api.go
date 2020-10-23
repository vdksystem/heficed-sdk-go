package protocompute

import (
	"encoding/json"
	"fmt"
)

type Instance struct {
	ID           int    `json:"id"`
	Status       string `json:"status"`
	Hostname     string `json:"hostname"`
	InstanceType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"instanceType"`
	Location struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Continent string `json:"continent"`
	} `json:"location"`
	Template interface{} `json:"template"`
	Network  struct {
		V4 struct {
			Ipaddress     string        `json:"ipaddress"`
			Netmask       string        `json:"netmask"`
			Gateway       string        `json:"gateway"`
			AdditionalIps []interface{} `json:"additionalIps"`
			Resolvers     []string      `json:"resolvers"`
		} `json:"v4"`
		V6 struct {
			Ipaddress     interface{} `json:"ipaddress"`
			Netmask       interface{} `json:"netmask"`
			Gateway       interface{} `json:"gateway"`
			AdditionalIps []string    `json:"additionalIps"`
			Resolvers     []string    `json:"resolvers"`
		} `json:"v6"`
	} `json:"network"`
	Billing struct {
		Product string `json:"product"`
		Type    struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		Status             string      `json:"status"`
		HourlySpendingRate interface{} `json:"hourlySpendingRate"`
		Price              float64     `json:"price"`
		StartDate          int         `json:"startDate"`
		EndDate            int         `json:"endDate"`
		SLA                struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"sla"`
		CancellationRequest interface{} `json:"cancellationRequest"`
	} `json:"billing"`
	Flavor struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Cpus []struct {
			Count int    `json:"count"`
			Type  string `json:"type"`
		} `json:"cpus"`
		Memory []struct {
			Total string `json:"total"`
			Count int    `json:"count"`
			Type  string `json:"type"`
		} `json:"memory"`
		Disks []struct {
			Count int    `json:"count"`
			Size  string `json:"size"`
			Type  string `json:"type"`
			Label string `json:"label"`
		} `json:"disks"`
		Nics    []interface{} `json:"nics"`
		Pricing interface{}   `json:"pricing"`
	} `json:"flavor"`
	Password  string `json:"password"`
	IpmiProxy struct {
		Hostname    string `json:"hostname"`
		Port        int    `json:"port"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		ReleaseDate string `json:"releaseDate"`
	} `json:"ipmiProxy"`
	CPULimited     bool   `json:"cpuLimited"`
	NetworkLimited bool   `json:"networkLimited"`
	Eta            string `json:"eta"`
}

func (p *protos) ListInstances() []Instance {
	path := "/instances/premium"
	resp, err := p.Request("GET", path, "")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var responseData map[string]json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return nil
	}
	var instances []Instance
	err = json.Unmarshal(responseData["data"], &instances)

	return instances
}

func (p *protos) GetInstance(id int) Instance {
	var instance Instance
	path := fmt.Sprintf("/instances/premium/%d", id)
	resp, err := p.Request("GET", path, "")
	if err != nil {
		fmt.Println(err)
		return instance
	}
	err = json.NewDecoder(resp.Body).Decode(&instance)
	if err != nil {
		fmt.Println(err)
	}

	return instance
}

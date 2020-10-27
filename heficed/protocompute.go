package heficed

import (
	"encoding/json"
	"fmt"
)

const APIPath = "protocompute"

type Protos struct {
	*Client
}

func (p *Protos) ListInstances() []ProtoCompute {
	path := APIPath + "/instances/premium"
	resp, err := p.request("GET", path, []byte{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var responseData map[string]json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return nil
	}
	var instances []ProtoCompute
	err = json.Unmarshal(responseData["data"], &instances)

	return instances
}

func (p *Protos) GetInstance(id int) ProtoCompute {
	var instance ProtoCompute
	path := fmt.Sprintf("%s/instances/premium/%d", APIPath, id)
	resp, err := p.request("GET", path, []byte{})
	if err != nil {
		fmt.Println("Error in GetInstance method", err)
		return instance
	}
	err = json.NewDecoder(resp.Body).Decode(&instance)
	if err != nil {
		fmt.Println(err)
	}

	return instance
}

type NewInstance struct {
	LocationId    string
	FlavorId      string
	Hostname      string
	BillingTypeId int
}

type createInstanceResponse struct {
	Items []item `json:"items"`
}

type item struct {
	InstanceId int `json:"instanceId"`
}

func (p *Protos) CreateInstance(cfg NewInstance) (int, error) {
	payload, _ := json.Marshal(cfg)
	path := fmt.Sprintf("%s/instances/premium/order", APIPath)
	resp, err := p.mock_request("POST", path, payload)
	if err != nil {
		fmt.Println("Error in CreateInstance method", err)
		return 0, err
	}

	var r createInstanceResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return 0, err
	}
	for _, i := range r.Items {
		return i.InstanceId, nil
	}

	return 0, nil
}

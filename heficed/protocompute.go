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
	resp, err := p.request("GET", path, "")
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
	resp, err := p.request("GET", path, "")
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

package heficed

type ProtoCompute struct {
	Id             int64                     `json:"id,omitempty"`
	Status         string                    `json:"status,omitempty"`
	Hostname       string                    `json:"hostname,omitempty"`
	InstanceType   *InstanceType             `json:"instanceType,omitempty"`
	Location       *Location                 `json:"location,omitempty"`
	Template       *KronosCloudTemplate      `json:"template,omitempty"`
	Network        *Network                  `json:"network,omitempty"`
	Billing        *ServiceBilling           `json:"billing,omitempty"`
	Flavor         *ProtoComputeFlavorSimple `json:"flavor,omitempty"`
	Password       string                    `json:"password,omitempty"`
	IpmiProxy      *IpmiProxy                `json:"ipmiProxy,omitempty"`
	CpuLimited     bool                      `json:"cpuLimited,omitempty"`
	NetworkLimited bool                      `json:"networkLimited,omitempty"`
	Eta            string                    `json:"eta,omitempty"`
}

type InstanceType struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Location struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Continent string `json:"continent,omitempty"`
}

type KronosCloudTemplate struct {
	Id      string `json:"id,omitempty"`
	Created int32  `json:"created,omitempty"`
	// Size in GB
	Size           float32 `json:"size,omitempty"`
	Version        string  `json:"version,omitempty"`
	Name           string  `json:"name,omitempty"`
	LocationId     string  `json:"locationId,omitempty"`
	InstanceTypeId string  `json:"instanceTypeId,omitempty"`
}

type ServiceBilling struct {
	Product            string       `json:"product,omitempty"`
	Type_              *BillingType `json:"type,omitempty"`
	Status             string       `json:"status,omitempty"`
	HourlySpendingRate float32      `json:"hourlySpendingRate,omitempty"`
	Price              float32      `json:"price,omitempty"`
	StartDate          int32        `json:"startDate,omitempty"`
	EndDate            int32        `json:"endDate,omitempty"`
	Sla                struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"sla,omitempty"`
	CancellationRequest *CancellationRequest `json:"cancellationRequest,omitempty"`
}

type BillingType struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProtoComputeFlavor struct {
	Id         string              `json:"id,omitempty"`
	LocationId string              `json:"locationId,omitempty"`
	Pricing    []ProtoComputePrice `json:"pricing,omitempty"`
	Cpus       []string            `json:"cpus,omitempty"`
	Memory     []string            `json:"memory,omitempty"`
	Disks      []string            `json:"disks,omitempty"`
}

type ProtoComputeFlavorSimple struct {
	Id     string               `json:"id,omitempty"`
	Name   string               `json:"name,omitempty"`
	Cpus   *[]ProtoComputeCpu   `json:"cpus,omitempty"`
	Memory []ProtoComputeMemory `json:"memory,omitempty"`
	Disks  *[]ProtoComputeDisk  `json:"disks,omitempty"`
	Nics   *[]ProtoComputeNic   `json:"nics,omitempty"`
}

type ProtoComputeCpu struct {
	Count int32  `json:"count,omitempty"`
	Type_ string `json:"type,omitempty"`
}

type ProtoComputeDisk struct {
	Count int32  `json:"count,omitempty"`
	Size  string `json:"size,omitempty"`
	Type_ string `json:"type,omitempty"`
	Label string `json:"label,omitempty"`
}

type ProtoComputeMemory struct {
	Total string `json:"total,omitempty"`
	Count int32  `json:"count,omitempty"`
	Type_ string `json:"type,omitempty"`
}

type ProtoComputeNic struct {
	Count int32  `json:"count,omitempty"`
	Type_ string `json:"type,omitempty"`
}

type ProtoComputePrice struct {
	BillingTypeId int32   `json:"billingTypeId,omitempty"`
	Amount        float32 `json:"amount,omitempty"`
	SetupFee      float32 `json:"setupFee,omitempty"`
}

type CancellationRequest struct {
	Type_       string `json:"type,omitempty"`
	Created     int32  `json:"created,omitempty"`
	Termination int32  `json:"termination,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

type Network struct {
	V4 NetworkV4
	V6 NetworkV6
}

type NetworkV4 struct {
	IpAddress     string   `json:"ipaddress,omitempty"`
	Netmask       string   `json:"netmask,omitempty"`
	Gateway       string   `json:"gateway,omitempty"`
	AdditionalIps []string `json:"additionalIps,omitempty"`
	Resolvers     []string `json:"resolvers,omitempty"`
}

type NetworkV6 struct {
	IpAddress     string   `json:"ipaddress,omitempty"`
	Netmask       string   `json:"netmask,omitempty"`
	Gateway       string   `json:"gateway,omitempty"`
	AdditionalIps []string `json:"additionalIps,omitempty"`
	Resolvers     []string `json:"resolvers,omitempty"`
}

type IpmiProxy struct {
	Hostname    string `json:"hostname,omitempty"`
	Port        int    `json:"port,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
}

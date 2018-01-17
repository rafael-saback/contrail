package models

// LoadbalancerMemberType

import "encoding/json"

// LoadbalancerMemberType
type LoadbalancerMemberType struct {
	Address           IpAddressType `json:"address,omitempty"`
	ProtocolPort      int           `json:"protocol_port,omitempty"`
	Status            string        `json:"status,omitempty"`
	StatusDescription string        `json:"status_description,omitempty"`
	Weight            int           `json:"weight,omitempty"`
	AdminState        bool          `json:"admin_state,omitempty"`
}

// String returns json representation of the object
func (model *LoadbalancerMemberType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerMemberType makes LoadbalancerMemberType
func MakeLoadbalancerMemberType() *LoadbalancerMemberType {
	return &LoadbalancerMemberType{
		//TODO(nati): Apply default
		StatusDescription: "",
		Weight:            0,
		AdminState:        false,
		Address:           MakeIpAddressType(),
		ProtocolPort:      0,
		Status:            "",
	}
}

// MakeLoadbalancerMemberTypeSlice() makes a slice of LoadbalancerMemberType
func MakeLoadbalancerMemberTypeSlice() []*LoadbalancerMemberType {
	return []*LoadbalancerMemberType{}
}

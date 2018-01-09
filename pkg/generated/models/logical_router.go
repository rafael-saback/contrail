package models

// LogicalRouter

import "encoding/json"

// LogicalRouter
type LogicalRouter struct {
	ParentUUID                string           `json:"parent_uuid"`
	ParentType                string           `json:"parent_type"`
	Perms2                    *PermType2       `json:"perms2"`
	UUID                      string           `json:"uuid"`
	Annotations               *KeyValuePairs   `json:"annotations"`
	FQName                    []string         `json:"fq_name"`
	IDPerms                   *IdPermsType     `json:"id_perms"`
	DisplayName               string           `json:"display_name"`
	VxlanNetworkIdentifier    string           `json:"vxlan_network_identifier"`
	ConfiguredRouteTargetList *RouteTargetList `json:"configured_route_target_list"`

	ServiceInstanceRefs         []*LogicalRouterServiceInstanceRef         `json:"service_instance_refs"`
	RouteTableRefs              []*LogicalRouterRouteTableRef              `json:"route_table_refs"`
	VirtualNetworkRefs          []*LogicalRouterVirtualNetworkRef          `json:"virtual_network_refs"`
	PhysicalRouterRefs          []*LogicalRouterPhysicalRouterRef          `json:"physical_router_refs"`
	BGPVPNRefs                  []*LogicalRouterBGPVPNRef                  `json:"bgpvpn_refs"`
	RouteTargetRefs             []*LogicalRouterRouteTargetRef             `json:"route_target_refs"`
	VirtualMachineInterfaceRefs []*LogicalRouterVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// LogicalRouterVirtualMachineInterfaceRef references each other
type LogicalRouterVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterServiceInstanceRef references each other
type LogicalRouterServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterRouteTableRef references each other
type LogicalRouterRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterVirtualNetworkRef references each other
type LogicalRouterVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterPhysicalRouterRef references each other
type LogicalRouterPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterBGPVPNRef references each other
type LogicalRouterBGPVPNRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterRouteTargetRef references each other
type LogicalRouterRouteTargetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *LogicalRouter) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLogicalRouter makes LogicalRouter
func MakeLogicalRouter() *LogicalRouter {
	return &LogicalRouter{
		//TODO(nati): Apply default
		ParentUUID:                "",
		ParentType:                "",
		Perms2:                    MakePermType2(),
		UUID:                      "",
		Annotations:               MakeKeyValuePairs(),
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		VxlanNetworkIdentifier:    "",
		ConfiguredRouteTargetList: MakeRouteTargetList(),
	}
}

// InterfaceToLogicalRouter makes LogicalRouter from interface
func InterfaceToLogicalRouter(iData interface{}) *LogicalRouter {
	data := iData.(map[string]interface{})
	return &LogicalRouter{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ConfiguredRouteTargetList: InterfaceToRouteTargetList(data["configured_route_target_list"]),

		//{"description":"List of route targets that represent this logical router, all virtual networks connected to this logical router will have this as their route target list.","type":"object","properties":{"route_target":{"type":"array","item":{"type":"string"}}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		VxlanNetworkIdentifier: data["vxlan_network_identifier"].(string),

		//{"description":"The VNI that needs to be associated with the internal VN if vxlan_routing mode is enabled.","type":"string"}

	}
}

// InterfaceToLogicalRouterSlice makes a slice of LogicalRouter from interface
func InterfaceToLogicalRouterSlice(data interface{}) []*LogicalRouter {
	list := data.([]interface{})
	result := MakeLogicalRouterSlice()
	for _, item := range list {
		result = append(result, InterfaceToLogicalRouter(item))
	}
	return result
}

// MakeLogicalRouterSlice() makes a slice of LogicalRouter
func MakeLogicalRouterSlice() []*LogicalRouter {
	return []*LogicalRouter{}
}

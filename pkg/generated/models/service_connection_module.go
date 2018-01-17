package models

// ServiceConnectionModule

import "encoding/json"

// ServiceConnectionModule
type ServiceConnectionModule struct {
	E2Service   E2servicetype         `json:"e2_service,omitempty"`
	UUID        string                `json:"uuid,omitempty"`
	ParentUUID  string                `json:"parent_uuid,omitempty"`
	ParentType  string                `json:"parent_type,omitempty"`
	IDPerms     *IdPermsType          `json:"id_perms,omitempty"`
	ServiceType ServiceConnectionType `json:"service_type,omitempty"`
	FQName      []string              `json:"fq_name,omitempty"`
	DisplayName string                `json:"display_name,omitempty"`
	Annotations *KeyValuePairs        `json:"annotations,omitempty"`
	Perms2      *PermType2            `json:"perms2,omitempty"`

	ServiceObjectRefs []*ServiceConnectionModuleServiceObjectRef `json:"service_object_refs,omitempty"`
}

// ServiceConnectionModuleServiceObjectRef references each other
type ServiceConnectionModuleServiceObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ServiceConnectionModule) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceConnectionModule makes ServiceConnectionModule
func MakeServiceConnectionModule() *ServiceConnectionModule {
	return &ServiceConnectionModule{
		//TODO(nati): Apply default
		Perms2:      MakePermType2(),
		ServiceType: MakeServiceConnectionType(),
		FQName:      []string{},
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		IDPerms:     MakeIdPermsType(),
		E2Service:   MakeE2servicetype(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
	}
}

// MakeServiceConnectionModuleSlice() makes a slice of ServiceConnectionModule
func MakeServiceConnectionModuleSlice() []*ServiceConnectionModule {
	return []*ServiceConnectionModule{}
}

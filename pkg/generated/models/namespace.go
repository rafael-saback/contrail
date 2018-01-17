package models

// Namespace

import "encoding/json"

// Namespace
type Namespace struct {
	NamespaceCidr *SubnetType    `json:"namespace_cidr,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
}

// String returns json representation of the object
func (model *Namespace) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNamespace makes Namespace
func MakeNamespace() *Namespace {
	return &Namespace{
		//TODO(nati): Apply default
		NamespaceCidr: MakeSubnetType(),
		ParentType:    "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		UUID:          "",
		ParentUUID:    "",
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
	}
}

// MakeNamespaceSlice() makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
	return []*Namespace{}
}

// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Description *string `json:",omitempty"`
	Id          *string `json:",omitempty"`
	Links       []Link  `json:",omitempty"`
	OrgId       *string `json:",omitempty"`
	Profile     *string `json:",omitempty"`
	PublicKey   *string `json:",omitempty"`
	PrivateKey  *string `json:",omitempty"`
	Roles       []Role  `json:",omitempty"`
}

// Link is autogenerated from the json schema
type Link struct {
	Href *string `json:",omitempty"`
	Rel  *string `json:",omitempty"`
}

// Role is autogenerated from the json schema
type Role struct {
	ProjectId *string `json:",omitempty"`
	OrgId     *string `json:",omitempty"`
	RoleName  *string `json:",omitempty"`
}

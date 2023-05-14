package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AdministrativeUnitsClient *msgraph.AdministrativeUnitsClient
	DirectoryObjectsClient    *msgraph.DirectoryObjectsClient
	GroupsClient              *msgraph.GroupsClient
}

func NewClient(o *common.ClientOptions) *Client {
	administrativeUnitsClient := msgraph.NewAdministrativeUnitsClient()
	o.ConfigureClient(&administrativeUnitsClient.BaseClient)

	// SDK uses wrong endpoint for v1.0 API, see https://github.com/manicminer/hamilton/issues/222
	administrativeUnitsClient.BaseClient.ApiVersion = msgraph.VersionBeta

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient()
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	groupsClient := msgraph.NewGroupsClient()
	o.ConfigureClient(&groupsClient.BaseClient)

	// Group members not returned in full when using v1.0 API, see https://github.com/hashicorp/terraform-provider-azuread/issues/1018
	groupsClient.BaseClient.ApiVersion = msgraph.Version10

	return &Client{
		AdministrativeUnitsClient: administrativeUnitsClient,
		DirectoryObjectsClient:    directoryObjectsClient,
		GroupsClient:              groupsClient,
	}
}

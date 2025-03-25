// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armpostgresqlflexibleservers.ClientFactory type.
type ServerFactory struct {
	// AdministratorsServer contains the fakes for client AdministratorsClient
	AdministratorsServer AdministratorsServer

	// BackupsServer contains the fakes for client BackupsClient
	BackupsServer BackupsServer

	// CheckNameAvailabilityServer contains the fakes for client CheckNameAvailabilityClient
	CheckNameAvailabilityServer CheckNameAvailabilityServer

	// CheckNameAvailabilityWithLocationServer contains the fakes for client CheckNameAvailabilityWithLocationClient
	CheckNameAvailabilityWithLocationServer CheckNameAvailabilityWithLocationServer

	// ConfigurationsServer contains the fakes for client ConfigurationsClient
	ConfigurationsServer ConfigurationsServer

	// DatabasesServer contains the fakes for client DatabasesClient
	DatabasesServer DatabasesServer

	// FirewallRulesServer contains the fakes for client FirewallRulesClient
	FirewallRulesServer FirewallRulesServer

	// FlexibleServerServer contains the fakes for client FlexibleServerClient
	FlexibleServerServer FlexibleServerServer

	// GetPrivateDNSZoneSuffixServer contains the fakes for client GetPrivateDNSZoneSuffixClient
	GetPrivateDNSZoneSuffixServer GetPrivateDNSZoneSuffixServer

	// LocationBasedCapabilitiesServer contains the fakes for client LocationBasedCapabilitiesClient
	LocationBasedCapabilitiesServer LocationBasedCapabilitiesServer

	// LogFilesServer contains the fakes for client LogFilesClient
	LogFilesServer LogFilesServer

	// LtrBackupOperationsServer contains the fakes for client LtrBackupOperationsClient
	LtrBackupOperationsServer LtrBackupOperationsServer

	// MigrationsServer contains the fakes for client MigrationsClient
	MigrationsServer MigrationsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PostgreSQLServerManagementServer contains the fakes for client PostgreSQLServerManagementClient
	PostgreSQLServerManagementServer PostgreSQLServerManagementServer

	// PrivateEndpointConnectionServer contains the fakes for client PrivateEndpointConnectionClient
	PrivateEndpointConnectionServer PrivateEndpointConnectionServer

	// PrivateEndpointConnectionsServer contains the fakes for client PrivateEndpointConnectionsClient
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer

	// PrivateLinkResourcesServer contains the fakes for client PrivateLinkResourcesClient
	PrivateLinkResourcesServer PrivateLinkResourcesServer

	// ReplicasServer contains the fakes for client ReplicasClient
	ReplicasServer ReplicasServer

	// ServerCapabilitiesServer contains the fakes for client ServerCapabilitiesClient
	ServerCapabilitiesServer ServerCapabilitiesServer

	// ServerThreatProtectionSettingsServer contains the fakes for client ServerThreatProtectionSettingsClient
	ServerThreatProtectionSettingsServer ServerThreatProtectionSettingsServer

	// ServersServer contains the fakes for client ServersClient
	ServersServer ServersServer

	// VirtualEndpointsServer contains the fakes for client VirtualEndpointsClient
	VirtualEndpointsServer VirtualEndpointsServer

	// VirtualNetworkSubnetUsageServer contains the fakes for client VirtualNetworkSubnetUsageClient
	VirtualNetworkSubnetUsageServer VirtualNetworkSubnetUsageServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armpostgresqlflexibleservers.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armpostgresqlflexibleservers.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                       *ServerFactory
	trMu                                      sync.Mutex
	trAdministratorsServer                    *AdministratorsServerTransport
	trBackupsServer                           *BackupsServerTransport
	trCheckNameAvailabilityServer             *CheckNameAvailabilityServerTransport
	trCheckNameAvailabilityWithLocationServer *CheckNameAvailabilityWithLocationServerTransport
	trConfigurationsServer                    *ConfigurationsServerTransport
	trDatabasesServer                         *DatabasesServerTransport
	trFirewallRulesServer                     *FirewallRulesServerTransport
	trFlexibleServerServer                    *FlexibleServerServerTransport
	trGetPrivateDNSZoneSuffixServer           *GetPrivateDNSZoneSuffixServerTransport
	trLocationBasedCapabilitiesServer         *LocationBasedCapabilitiesServerTransport
	trLogFilesServer                          *LogFilesServerTransport
	trLtrBackupOperationsServer               *LtrBackupOperationsServerTransport
	trMigrationsServer                        *MigrationsServerTransport
	trOperationsServer                        *OperationsServerTransport
	trPostgreSQLServerManagementServer        *PostgreSQLServerManagementServerTransport
	trPrivateEndpointConnectionServer         *PrivateEndpointConnectionServerTransport
	trPrivateEndpointConnectionsServer        *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer              *PrivateLinkResourcesServerTransport
	trReplicasServer                          *ReplicasServerTransport
	trServerCapabilitiesServer                *ServerCapabilitiesServerTransport
	trServerThreatProtectionSettingsServer    *ServerThreatProtectionSettingsServerTransport
	trServersServer                           *ServersServerTransport
	trVirtualEndpointsServer                  *VirtualEndpointsServerTransport
	trVirtualNetworkSubnetUsageServer         *VirtualNetworkSubnetUsageServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AdministratorsClient":
		initServer(s, &s.trAdministratorsServer, func() *AdministratorsServerTransport {
			return NewAdministratorsServerTransport(&s.srv.AdministratorsServer)
		})
		resp, err = s.trAdministratorsServer.Do(req)
	case "BackupsClient":
		initServer(s, &s.trBackupsServer, func() *BackupsServerTransport { return NewBackupsServerTransport(&s.srv.BackupsServer) })
		resp, err = s.trBackupsServer.Do(req)
	case "CheckNameAvailabilityClient":
		initServer(s, &s.trCheckNameAvailabilityServer, func() *CheckNameAvailabilityServerTransport {
			return NewCheckNameAvailabilityServerTransport(&s.srv.CheckNameAvailabilityServer)
		})
		resp, err = s.trCheckNameAvailabilityServer.Do(req)
	case "CheckNameAvailabilityWithLocationClient":
		initServer(s, &s.trCheckNameAvailabilityWithLocationServer, func() *CheckNameAvailabilityWithLocationServerTransport {
			return NewCheckNameAvailabilityWithLocationServerTransport(&s.srv.CheckNameAvailabilityWithLocationServer)
		})
		resp, err = s.trCheckNameAvailabilityWithLocationServer.Do(req)
	case "ConfigurationsClient":
		initServer(s, &s.trConfigurationsServer, func() *ConfigurationsServerTransport {
			return NewConfigurationsServerTransport(&s.srv.ConfigurationsServer)
		})
		resp, err = s.trConfigurationsServer.Do(req)
	case "DatabasesClient":
		initServer(s, &s.trDatabasesServer, func() *DatabasesServerTransport { return NewDatabasesServerTransport(&s.srv.DatabasesServer) })
		resp, err = s.trDatabasesServer.Do(req)
	case "FirewallRulesClient":
		initServer(s, &s.trFirewallRulesServer, func() *FirewallRulesServerTransport {
			return NewFirewallRulesServerTransport(&s.srv.FirewallRulesServer)
		})
		resp, err = s.trFirewallRulesServer.Do(req)
	case "FlexibleServerClient":
		initServer(s, &s.trFlexibleServerServer, func() *FlexibleServerServerTransport {
			return NewFlexibleServerServerTransport(&s.srv.FlexibleServerServer)
		})
		resp, err = s.trFlexibleServerServer.Do(req)
	case "GetPrivateDNSZoneSuffixClient":
		initServer(s, &s.trGetPrivateDNSZoneSuffixServer, func() *GetPrivateDNSZoneSuffixServerTransport {
			return NewGetPrivateDNSZoneSuffixServerTransport(&s.srv.GetPrivateDNSZoneSuffixServer)
		})
		resp, err = s.trGetPrivateDNSZoneSuffixServer.Do(req)
	case "LocationBasedCapabilitiesClient":
		initServer(s, &s.trLocationBasedCapabilitiesServer, func() *LocationBasedCapabilitiesServerTransport {
			return NewLocationBasedCapabilitiesServerTransport(&s.srv.LocationBasedCapabilitiesServer)
		})
		resp, err = s.trLocationBasedCapabilitiesServer.Do(req)
	case "LogFilesClient":
		initServer(s, &s.trLogFilesServer, func() *LogFilesServerTransport { return NewLogFilesServerTransport(&s.srv.LogFilesServer) })
		resp, err = s.trLogFilesServer.Do(req)
	case "LtrBackupOperationsClient":
		initServer(s, &s.trLtrBackupOperationsServer, func() *LtrBackupOperationsServerTransport {
			return NewLtrBackupOperationsServerTransport(&s.srv.LtrBackupOperationsServer)
		})
		resp, err = s.trLtrBackupOperationsServer.Do(req)
	case "MigrationsClient":
		initServer(s, &s.trMigrationsServer, func() *MigrationsServerTransport { return NewMigrationsServerTransport(&s.srv.MigrationsServer) })
		resp, err = s.trMigrationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PostgreSQLServerManagementClient":
		initServer(s, &s.trPostgreSQLServerManagementServer, func() *PostgreSQLServerManagementServerTransport {
			return NewPostgreSQLServerManagementServerTransport(&s.srv.PostgreSQLServerManagementServer)
		})
		resp, err = s.trPostgreSQLServerManagementServer.Do(req)
	case "PrivateEndpointConnectionClient":
		initServer(s, &s.trPrivateEndpointConnectionServer, func() *PrivateEndpointConnectionServerTransport {
			return NewPrivateEndpointConnectionServerTransport(&s.srv.PrivateEndpointConnectionServer)
		})
		resp, err = s.trPrivateEndpointConnectionServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "ReplicasClient":
		initServer(s, &s.trReplicasServer, func() *ReplicasServerTransport { return NewReplicasServerTransport(&s.srv.ReplicasServer) })
		resp, err = s.trReplicasServer.Do(req)
	case "ServerCapabilitiesClient":
		initServer(s, &s.trServerCapabilitiesServer, func() *ServerCapabilitiesServerTransport {
			return NewServerCapabilitiesServerTransport(&s.srv.ServerCapabilitiesServer)
		})
		resp, err = s.trServerCapabilitiesServer.Do(req)
	case "ServerThreatProtectionSettingsClient":
		initServer(s, &s.trServerThreatProtectionSettingsServer, func() *ServerThreatProtectionSettingsServerTransport {
			return NewServerThreatProtectionSettingsServerTransport(&s.srv.ServerThreatProtectionSettingsServer)
		})
		resp, err = s.trServerThreatProtectionSettingsServer.Do(req)
	case "ServersClient":
		initServer(s, &s.trServersServer, func() *ServersServerTransport { return NewServersServerTransport(&s.srv.ServersServer) })
		resp, err = s.trServersServer.Do(req)
	case "VirtualEndpointsClient":
		initServer(s, &s.trVirtualEndpointsServer, func() *VirtualEndpointsServerTransport {
			return NewVirtualEndpointsServerTransport(&s.srv.VirtualEndpointsServer)
		})
		resp, err = s.trVirtualEndpointsServer.Do(req)
	case "VirtualNetworkSubnetUsageClient":
		initServer(s, &s.trVirtualNetworkSubnetUsageServer, func() *VirtualNetworkSubnetUsageServerTransport {
			return NewVirtualNetworkSubnetUsageServerTransport(&s.srv.VirtualNetworkSubnetUsageServer)
		})
		resp, err = s.trVirtualNetworkSubnetUsageServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}

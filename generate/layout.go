//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

var layout = apiInfo{
	"NetworkService": {
		"addNetworkDevice",
		"addNetworkServiceProvider",
		"addOpenDaylightController",
		"createNetwork",
		"createNetworkACL",
		"createNetworkACLList",
		"createNetworkOffering",
		"createPhysicalNetwork",
		"createServiceInstance",
		"createStorageNetworkIpRange",
		"dedicatePublicIpRange",
		"deleteNetwork",
		"deleteNetworkACL",
		"deleteNetworkACLList",
		"deleteNetworkDevice",
		"deleteNetworkOffering",
		"deleteNetworkServiceProvider",
		"deleteOpenDaylightController",
		"deletePhysicalNetwork",
		"deleteStorageNetworkIpRange",
		"listBrocadeVcsDeviceNetworks",
		"listF5LoadBalancerNetworks",
		"listNetscalerLoadBalancerNetworks",
		"listNetworkACLLists",
		"listNetworkACLs",
		"listNetworkDevice",
		"listNetworkIsolationMethods",
		"listNetworkOfferings",
		"listNetworkServiceProviders",
		"listNetworks",
		"listNiciraNvpDeviceNetworks",
		"listOpenDaylightControllers",
		"listPaloAltoFirewallNetworks",
		"listPhysicalNetworks",
		"listSrxFirewallNetworks",
		"listStorageNetworkIpRange",
		"listSupportedNetworkServices",
		"releasePublicIpRange",
		"replaceNetworkACLList",
		"restartNetwork",
		"updateNetwork",
		"updateNetworkACLItem",
		"updateNetworkACLList",
		"updateNetworkOffering",
		"updateNetworkServiceProvider",
		"updatePhysicalNetwork",
		"updateStorageNetworkIpRange",
	},
	"LoadBalancerService": {
		"addF5LoadBalancer",
		"addNetscalerLoadBalancer",
		"assignCertToLoadBalancer",
		"assignToGlobalLoadBalancerRule",
		"assignToLoadBalancerRule",
		"configureF5LoadBalancer",
		"configureNetscalerLoadBalancer",
		"createGlobalLoadBalancerRule",
		"createLBHealthCheckPolicy",
		"createLBStickinessPolicy",
		"createLoadBalancer",
		"createLoadBalancerRule",
		"deleteF5LoadBalancer",
		"deleteGlobalLoadBalancerRule",
		"deleteLBHealthCheckPolicy",
		"deleteLBStickinessPolicy",
		"deleteLoadBalancer",
		"deleteLoadBalancerRule",
		"deleteNetscalerLoadBalancer",
		"deleteSslCert",
		"listF5LoadBalancers",
		"listGlobalLoadBalancerRules",
		"listLBHealthCheckPolicies",
		"listLBStickinessPolicies",
		"listLoadBalancerRuleInstances",
		"listLoadBalancerRules",
		"listLoadBalancers",
		"listNetscalerLoadBalancers",
		"listSslCerts",
		"removeCertFromLoadBalancer",
		"removeFromGlobalLoadBalancerRule",
		"removeFromLoadBalancerRule",
		"updateGlobalLoadBalancerRule",
		"updateLBHealthCheckPolicy",
		"updateLBStickinessPolicy",
		"updateLoadBalancer",
		"updateLoadBalancerRule",
		"uploadSslCert",
	},
	"VirtualMachineService": {
		"addNicToVirtualMachine",
		"assignVirtualMachine",
		"changeServiceForVirtualMachine",
		"cleanVMReservations",
		"deployVirtualMachine",
		"destroyVirtualMachine",
		"expungeVirtualMachine",
		"getVMPassword",
		"listInternalLoadBalancerVMs",
		"listVirtualMachines",
		"listVirtualMachinesMetrics",
		"migrateVirtualMachine",
		"migrateVirtualMachineWithVolume",
		"rebootVirtualMachine",
		"recoverVirtualMachine",
		"removeNicFromVirtualMachine",
		"resetPasswordForVirtualMachine",
		"resetSSHKeyForVirtualMachine",
		"restoreVirtualMachine",
		"scaleVirtualMachine",
		"startInternalLoadBalancerVM",
		"startVirtualMachine",
		"stopInternalLoadBalancerVM",
		"stopVirtualMachine",
		"updateDefaultNicForVirtualMachine",
		"updateVirtualMachine",
	},
	"HostService": {
		"addBaremetalHost",
		"addGloboDnsHost",
		"addHost",
		"addSecondaryStorage",
		"cancelHostMaintenance",
		"dedicateHost",
		"deleteHost",
		"disableOutOfBandManagementForHost",
		"enableOutOfBandManagementForHost",
		"findHostsForMigration",
		"listDedicatedHosts",
		"listHostTags",
		"listHosts",
		"listHostsMetrics",
		"prepareHostForMaintenance",
		"reconnectHost",
		"releaseDedicatedHost",
		"releaseHostReservation",
		"updateHost",
		"updateHostPassword",
	},
	"FirewallService": {
		"addPaloAltoFirewall",
		"addSrxFirewall",
		"configurePaloAltoFirewall",
		"configureSrxFirewall",
		"createEgressFirewallRule",
		"createFirewallRule",
		"createPortForwardingRule",
		"deleteEgressFirewallRule",
		"deleteFirewallRule",
		"deletePaloAltoFirewall",
		"deletePortForwardingRule",
		"deleteSrxFirewall",
		"listEgressFirewallRules",
		"listFirewallRules",
		"listPaloAltoFirewalls",
		"listPortForwardingRules",
		"listSrxFirewalls",
		"updateEgressFirewallRule",
		"updateFirewallRule",
		"updatePortForwardingRule",
	},
	"AutoScaleService": {
		"createAutoScalePolicy",
		"createAutoScaleVmGroup",
		"createAutoScaleVmProfile",
		"createCondition",
		"createCounter",
		"deleteAutoScalePolicy",
		"deleteAutoScaleVmGroup",
		"deleteAutoScaleVmProfile",
		"deleteCondition",
		"deleteCounter",
		"disableAutoScaleVmGroup",
		"enableAutoScaleVmGroup",
		"listAutoScalePolicies",
		"listAutoScaleVmGroups",
		"listAutoScaleVmProfiles",
		"listConditions",
		"listCounters",
		"updateAutoScalePolicy",
		"updateAutoScaleVmGroup",
		"updateAutoScaleVmProfile",
	},
	"VolumeService": {
		"attachVolume",
		"createVolume",
		"deleteVolume",
		"detachVolume",
		"extractVolume",
		"getPathForVolume",
		"getSolidFireVolumeAccessGroupId",
		"getSolidFireVolumeSize",
		"getUploadParamsForVolume",
		"getVolumeiScsiName",
		"listElastistorVolume",
		"listVolumes",
		"migrateVolume",
		"resizeVolume",
		"updateVolume",
		"uploadVolume",
	},
	"VPNService": {
		"createRemoteAccessVpn",
		"createVpnConnection",
		"createVpnCustomerGateway",
		"createVpnGateway",
		"deleteRemoteAccessVpn",
		"deleteVpnConnection",
		"deleteVpnCustomerGateway",
		"deleteVpnGateway",
		"listRemoteAccessVpns",
		"listVpnConnections",
		"listVpnCustomerGateways",
		"listVpnGateways",
		"resetVpnConnection",
		"updateRemoteAccessVpn",
		"updateVpnConnection",
		"updateVpnCustomerGateway",
		"updateVpnGateway",
	},
	"UserService": {
		"addVpnUser",
		"createUser",
		"deleteUser",
		"disableUser",
		"enableUser",
		"getUser",
		"getVirtualMachineUserData",
		"importLdapUsers",
		"listLdapUsers",
		"listUsers",
		"listVpnUsers",
		"lockUser",
		"registerUserKeys",
		"removeVpnUser",
		"updateUser",
	},
	"TemplateService": {
		"copyTemplate",
		"createTemplate",
		"deleteTemplate",
		"extractTemplate",
		"getUploadParamsForTemplate",
		"listTemplatePermissions",
		"listTemplates",
		"prepareTemplate",
		"registerTemplate",
		"updateTemplate",
		"updateTemplatePermissions",
		"upgradeRouterTemplate",
	},
	"VPCService": {
		"createPrivateGateway",
		"createStaticRoute",
		"createVPC",
		"createVPCOffering",
		"deletePrivateGateway",
		"deleteStaticRoute",
		"deleteVPC",
		"deleteVPCOffering",
		"listPrivateGateways",
		"listStaticRoutes",
		"listVPCOfferings",
		"listVPCs",
		"restartVPC",
		"updateVPC",
		"updateVPCOffering",
	},
	"ZoneService": {
		"addVmwareDc",
		"createZone",
		"dedicateZone",
		"deleteZone",
		"disableOutOfBandManagementForZone",
		"enableOutOfBandManagementForZone",
		"listDedicatedZones",
		"listVmwareDcs",
		"listZones",
		"releaseDedicatedZone",
		"removeVmwareDc",
		"updateZone",
	},
	"AccountService": {
		"addAccountToProject",
		"createAccount",
		"deleteAccount",
		"deleteAccountFromProject",
		"disableAccount",
		"enableAccount",
		"getSolidFireAccountId",
		"listAccounts",
		"listProjectAccounts",
		"lockAccount",
		"markDefaultZoneForAccount",
		"updateAccount",
	},
	"SnapshotService": {
		"createSnapshot",
		"createSnapshotPolicy",
		"createVMSnapshot",
		"deleteSnapshot",
		"deleteSnapshotPolicies",
		"deleteVMSnapshot",
		"listSnapshotPolicies",
		"listSnapshots",
		"listVMSnapshot",
		"revertSnapshot",
		"revertToVMSnapshot",
		"updateSnapshotPolicy",
	},
	"UsageService": {
		"addTrafficMonitor",
		"addTrafficType",
		"deleteTrafficMonitor",
		"deleteTrafficType",
		"generateUsageRecords",
		"listTrafficMonitors",
		"listTrafficTypeImplementors",
		"listTrafficTypes",
		"listUsageRecords",
		"listUsageTypes",
		"removeRawUsageRecords",
		"updateTrafficType",
	},
	"ClusterService": {
		"addCluster",
		"dedicateCluster",
		"deleteCluster",
		"disableOutOfBandManagementForCluster",
		"enableOutOfBandManagementForCluster",
		"listClusters",
		"listClustersMetrics",
		"listDedicatedClusters",
		"releaseDedicatedCluster",
		"updateCluster",
	},
	"PoolService": {
		"createStoragePool",
		"deleteStoragePool",
		"findStoragePoolsForMigration",
		"listElastistorPool",
		"listStoragePools",
		"updateStoragePool",
	},
	"ISOService": {
		"attachIso",
		"copyIso",
		"deleteIso",
		"detachIso",
		"extractIso",
		"listIsoPermissions",
		"listIsos",
		"registerIso",
		"updateIso",
		"updateIsoPermissions",
	},
	"ExternalDeviceService": {
		"addCiscoAsa1000vResource",
		"addCiscoVnmcResource",
		"deleteCiscoAsa1000vResource",
		"deleteCiscoNexusVSM",
		"deleteCiscoVnmcResource",
		"disableCiscoNexusVSM",
		"enableCiscoNexusVSM",
		"listCiscoAsa1000vResources",
		"listCiscoNexusVSMs",
		"listCiscoVnmcResources",
	},
	"VLANService": {
		"createVlanIpRange",
		"dedicateGuestVlanRange",
		"deleteVlanIpRange",
		"listDedicatedGuestVlanRanges",
		"listVlanIpRanges",
		"releaseDedicatedGuestVlanRange",
	},
	"RouterService": {
		"changeServiceForRouter",
		"configureVirtualRouterElement",
		"createVirtualRouterElement",
		"destroyRouter",
		"listRouters",
		"listVirtualRouterElements",
		"rebootRouter",
		"startRouter",
		"stopRouter",
	},
	"ProjectService": {
		"activateProject",
		"createProject",
		"deleteProject",
		"deleteProjectInvitation",
		"listProjectInvitations",
		"listProjects",
		"suspendProject",
		"updateProject",
		"updateProjectInvitation",
	},
	"GuestOSService": {
		"addGuestOs",
		"addGuestOsMapping",
		"listGuestOsMapping",
		"listOsCategories",
		"listOsTypes",
		"removeGuestOs",
		"removeGuestOsMapping",
		"updateGuestOs",
		"updateGuestOsMapping",
	},
	"BaremetalService": {
		"addBaremetalDhcp",
		"addBaremetalPxeKickStartServer",
		"addBaremetalPxePingServer",
		"addBaremetalRct",
		"deleteBaremetalRct",
		"listBaremetalDhcp",
		"listBaremetalPxeServers",
		"listBaremetalRct",
		"notifyBaremetalProvisionDone",
	},
	"SystemVMService": {
		"changeServiceForSystemVm",
		"destroySystemVm",
		"listSystemVms",
		"migrateSystemVm",
		"rebootSystemVm",
		"scaleSystemVm",
		"startSystemVm",
		"stopSystemVm",
	},
	"RoleService": {
		"createRole",
		"createRolePermission",
		"deleteRole",
		"deleteRolePermission",
		"listRolePermissions",
		"listRoles",
		"updateRole",
		"updateRolePermission",
	},
	"LDAPService": {
		"addLdapConfiguration",
		"deleteLdapConfiguration",
		"ldapConfig",
		"ldapCreateAccount",
		"ldapRemove",
		"linkDomainToLdap",
		"listLdapConfigurations",
		"searchLdap",
	},
	"AuthenticationService": {
		"login",
		"logout",
	},
	"SecurityGroupService": {
		"authorizeSecurityGroupEgress",
		"authorizeSecurityGroupIngress",
		"createSecurityGroup",
		"deleteSecurityGroup",
		"listSecurityGroups",
		"revokeSecurityGroupEgress",
		"revokeSecurityGroupIngress",
	},
	"QuotaService": {
		"quotaIsEnabled",
	},
	"PodService": {
		"createPod",
		"dedicatePod",
		"deletePod",
		"listDedicatedPods",
		"listPods",
		"releaseDedicatedPod",
		"updatePod",
	},
	"ImageStoreService": {
		"addImageStore",
		"createSecondaryStagingStore",
		"deleteImageStore",
		"deleteSecondaryStagingStore",
		"listImageStores",
		"listSecondaryStagingStores",
		"updateCloudToUseObjectStore",
	},
	"CertificateService": {
		"uploadCustomCertificate",
	},
	"UCSService": {
		"addUcsManager",
		"associateUcsProfileToBlade",
		"deleteUcsManager",
		"listUcsBlades",
		"listUcsManagers",
		"listUcsProfiles",
	},
	"AddressService": {
		"associateIpAddress",
		"disassociateIpAddress",
		"listPublicIpAddresses",
		"updateIpAddress",
	},
	"NATService": {
		"createIpForwardingRule",
		"deleteIpForwardingRule",
		"disableStaticNat",
		"enableStaticNat",
		"listIpForwardingRules",
	},
	"DomainService": {
		"createDomain",
		"deleteDomain",
		"listDomainChildren",
		"listDomains",
		"updateDomain",
	},
	"AffinityGroupService": {
		"createAffinityGroup",
		"deleteAffinityGroup",
		"listAffinityGroupTypes",
		"listAffinityGroups",
		"updateVMAffinityGroup",
	},
	"VMGroupService": {
		"createInstanceGroup",
		"deleteInstanceGroup",
		"listInstanceGroups",
		"updateInstanceGroup",
	},
	"ServiceOfferingService": {
		"createServiceOffering",
		"deleteServiceOffering",
		"listServiceOfferings",
		"updateServiceOffering",
	},
	"SSHService": {
		"createSSHKeyPair",
		"deleteSSHKeyPair",
		"listSSHKeyPairs",
		"registerSSHKeyPair",
	},
	"ResourcetagsService": {
		"createTags",
		"deleteTags",
		"listStorageTags",
		"listTags",
	},
	"ResourcemetadataService": {
		"addResourceDetail",
		"getVolumeSnapshotDetails",
		"listResourceDetails",
		"removeResourceDetail",
	},
	"RegionService": {
		"addRegion",
		"listRegions",
		"removeRegion",
		"updateRegion",
	},
	"NuageVSPService": {
		"addNuageVspDevice",
		"deleteNuageVspDevice",
		"listNuageVspDevices",
		"updateNuageVspDevice",
	},
	"NicService": {
		"addIpToNic",
		"listNics",
		"removeIpFromNic",
		"updateVmNicIp",
	},
	"MetricsService": {
		"listInfrastructure",
		"listStoragePoolsMetrics",
		"listVolumesMetrics",
		"listZonesMetrics",
	},
	"EventService": {
		"archiveEvents",
		"deleteEvents",
		"listEventTypes",
		"listEvents",
	},
	"DiskOfferingService": {
		"createDiskOffering",
		"deleteDiskOffering",
		"listDiskOfferings",
		"updateDiskOffering",
	},
	"ConfigurationService": {
		"listCapabilities",
		"listConfigurations",
		"listDeploymentPlanners",
		"updateConfiguration",
	},
	"AlertService": {
		"archiveAlerts",
		"deleteAlerts",
		"generateAlert",
		"listAlerts",
	},
	"StoragePoolService": {
		"cancelStorageMaintenance",
		"enableStorageMaintenance",
		"listStorageProviders",
	},
	"PortableIPService": {
		"createPortableIpRange",
		"deletePortableIpRange",
		"listPortableIpRanges",
	},
	"OutofbandManagementService": {
		"changeOutOfBandManagementPassword",
		"configureOutOfBandManagement",
		"issueOutOfBandManagementPowerAction",
	},
	"NiciraNVPService": {
		"addNiciraNvpDevice",
		"deleteNiciraNvpDevice",
		"listNiciraNvpDevices",
	},
	"LimitService": {
		"listResourceLimits",
		"updateResourceCount",
		"updateResourceLimit",
	},
	"InternalLBService": {
		"configureInternalLoadBalancerElement",
		"createInternalLoadBalancerElement",
		"listInternalLoadBalancerElements",
	},
	"HypervisorService": {
		"listHypervisorCapabilities",
		"listHypervisors",
		"updateHypervisorCapabilities",
	},
	"ExtLoadBalancer": {
		"addExternalLoadBalancer",
		"deleteExternalLoadBalancer",
		"listExternalLoadBalancers",
	},
	"ExtFirewallService": {
		"addExternalFirewall",
		"deleteExternalFirewall",
		"listExternalFirewalls",
	},
	"BrocadeVCSService": {
		"addBrocadeVcsDevice",
		"deleteBrocadeVcsDevice",
		"listBrocadeVcsDevices",
	},
	"BigSwitchBCFService": {
		"addBigSwitchBcfDevice",
		"deleteBigSwitchBcfDevice",
		"listBigSwitchBcfDevices",
	},
	"APIDiscoveryService": {
		"getApiLimit",
		"listApis",
		"resetApiLimit",
	},
	"SwiftService": {
		"addSwift",
		"listSwifts",
	},
	"OvsElementService": {
		"configureOvsElement",
		"listOvsElements",
	},
	"AsyncjobService": {
		"listAsyncJobs",
		"queryAsyncJobResult",
	},
	"StratosphereSSP": {
		"addStratosphereSsp",
		"deleteStratosphereSsp",
	},
	"SystemCapacityService": {
		"listCapacity",
	},
	"S3Service": {
		"addImageStoreS3",
	},
	"MiscService": {
		"listElastistorInterface",
	},
	"CloudIdentifierService": {
		"getCloudIdentifier",
	},
}
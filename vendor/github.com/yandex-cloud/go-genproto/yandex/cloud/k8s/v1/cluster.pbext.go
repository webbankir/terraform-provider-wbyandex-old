// Code generated by protoc-gen-goext. DO NOT EDIT.

package k8s

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

type Cluster_InternetGateway = isCluster_InternetGateway

func (m *Cluster) SetInternetGateway(v Cluster_InternetGateway) {
	m.InternetGateway = v
}

func (m *Cluster) SetId(v string) {
	m.Id = v
}

func (m *Cluster) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Cluster) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Cluster) SetName(v string) {
	m.Name = v
}

func (m *Cluster) SetDescription(v string) {
	m.Description = v
}

func (m *Cluster) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Cluster) SetStatus(v Cluster_Status) {
	m.Status = v
}

func (m *Cluster) SetHealth(v Cluster_Health) {
	m.Health = v
}

func (m *Cluster) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Cluster) SetMaster(v *Master) {
	m.Master = v
}

func (m *Cluster) SetIpAllocationPolicy(v *IPAllocationPolicy) {
	m.IpAllocationPolicy = v
}

func (m *Cluster) SetGatewayIpv4Address(v string) {
	m.InternetGateway = &Cluster_GatewayIpv4Address{
		GatewayIpv4Address: v,
	}
}

func (m *Cluster) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Cluster) SetNodeServiceAccountId(v string) {
	m.NodeServiceAccountId = v
}

func (m *Cluster) SetReleaseChannel(v ReleaseChannel) {
	m.ReleaseChannel = v
}

func (m *Cluster) SetNetworkPolicy(v *NetworkPolicy) {
	m.NetworkPolicy = v
}

func (m *Cluster) SetKmsProvider(v *KMSProvider) {
	m.KmsProvider = v
}

func (m *Cluster) SetLogGroupId(v string) {
	m.LogGroupId = v
}

type Master_MasterType = isMaster_MasterType

func (m *Master) SetMasterType(v Master_MasterType) {
	m.MasterType = v
}

func (m *Master) SetZonalMaster(v *ZonalMaster) {
	m.MasterType = &Master_ZonalMaster{
		ZonalMaster: v,
	}
}

func (m *Master) SetRegionalMaster(v *RegionalMaster) {
	m.MasterType = &Master_RegionalMaster{
		RegionalMaster: v,
	}
}

func (m *Master) SetVersion(v string) {
	m.Version = v
}

func (m *Master) SetEndpoints(v *MasterEndpoints) {
	m.Endpoints = v
}

func (m *Master) SetMasterAuth(v *MasterAuth) {
	m.MasterAuth = v
}

func (m *Master) SetVersionInfo(v *VersionInfo) {
	m.VersionInfo = v
}

func (m *Master) SetMaintenancePolicy(v *MasterMaintenancePolicy) {
	m.MaintenancePolicy = v
}

func (m *Master) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *MasterAuth) SetClusterCaCertificate(v string) {
	m.ClusterCaCertificate = v
}

func (m *ZonalMaster) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *ZonalMaster) SetInternalV4Address(v string) {
	m.InternalV4Address = v
}

func (m *ZonalMaster) SetExternalV4Address(v string) {
	m.ExternalV4Address = v
}

func (m *RegionalMaster) SetRegionId(v string) {
	m.RegionId = v
}

func (m *RegionalMaster) SetInternalV4Address(v string) {
	m.InternalV4Address = v
}

func (m *RegionalMaster) SetExternalV4Address(v string) {
	m.ExternalV4Address = v
}

func (m *MasterEndpoints) SetInternalV4Endpoint(v string) {
	m.InternalV4Endpoint = v
}

func (m *MasterEndpoints) SetExternalV4Endpoint(v string) {
	m.ExternalV4Endpoint = v
}

func (m *IPAllocationPolicy) SetClusterIpv4CidrBlock(v string) {
	m.ClusterIpv4CidrBlock = v
}

func (m *IPAllocationPolicy) SetNodeIpv4CidrMaskSize(v int64) {
	m.NodeIpv4CidrMaskSize = v
}

func (m *IPAllocationPolicy) SetServiceIpv4CidrBlock(v string) {
	m.ServiceIpv4CidrBlock = v
}

func (m *IPAllocationPolicy) SetClusterIpv6CidrBlock(v string) {
	m.ClusterIpv6CidrBlock = v
}

func (m *IPAllocationPolicy) SetServiceIpv6CidrBlock(v string) {
	m.ServiceIpv6CidrBlock = v
}

func (m *MasterMaintenancePolicy) SetAutoUpgrade(v bool) {
	m.AutoUpgrade = v
}

func (m *MasterMaintenancePolicy) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *NetworkPolicy) SetProvider(v NetworkPolicy_Provider) {
	m.Provider = v
}

func (m *KMSProvider) SetKeyId(v string) {
	m.KeyId = v
}

package helpers

import "time"

const (
	// IBM PPC Instance

	PPCInstanceName                      = "ppc_instance_name"
	PPCInstanceDate                      = "ppc_creation_date"
	PPCInstanceSSHKeyName                = "ppc_key_pair_name"
	PPCInstanceImageId                   = "ppc_image_id"
	PPCInstanceProcessors                = "ppc_processors"
	PPCInstanceProcType                  = "ppc_proc_type"
	PPCInstanceMemory                    = "ppc_memory"
	PPCInstanceSystemType                = "ppc_sys_type"
	PPCInstanceId                        = "ppc_instance_id"
	PPCInstanceDiskSize                  = "ppc_disk_size"
	PPCInstanceStatus                    = "ppc_instance_status"
	PPCInstanceMinProc                   = "ppc_minproc"
	PPCInstanceVolumeIds                 = "ppc_volume_ids"
	PPCInstanceNetworkIds                = "ppc_network_ids"
	PPCInstancePublicNetwork             = "ppc_public_network"
	PPCInstanceMigratable                = "ppc_migratable"
	PPCCloudInstanceId                   = "ppc_cloud_instance_id"
	PPCCloudInstanceSubnetName           = "ppc_cloud_instance_subnet_name"
	PPCInstanceMimMem                    = "ppc_minmem"
	PPCInstanceMaxProc                   = "ppc_maxproc"
	PPCInstanceMaxMem                    = "ppc_maxmem"
	PPCInstanceReboot                    = "ppc_reboot"
	PPCTenantId                          = "ppc_tenant_id"
	PPCVirtualCoresAssigned              = "ppc_virtual_cores_assigned"
	PPCVirtualCoresMax                   = "ppc_virtual_cores_max"
	PPCVirtualCoresMin                   = "ppc_virtual_cores_min"
	PPCInstancePVMNetwork                = "ppc_instance_pvm_network"
	PPCInstanceStorageType               = "ppc_storage_type"
	PPCInstanceStorageConnection         = "ppc_storage_connection"
	PPCInstanceStoragePool               = "ppc_instance_storage_pool"
	PPCInstanceStorageAffinityPool       = "ppc_instance_storage_affinity_pool"
	PPCInstanceLicenseRepositoryCapacity = "ppc_license_repository_capacity"

	PPCInstanceHealthStatus      = "ppc_health_status"
	PPCInstanceReplicants        = "ppc_replicants"
	PPCInstanceReplicationPolicy = "ppc_replication_policy"
	PPCInstanceReplicationScheme = "ppc_replication_scheme"
	PPCInstanceProgress          = "ppc_progress"
	PPCInstanceUserData          = "ppc_user_data"
	PPCInstancePinPolicy         = "ppc_spn_policy"

	// IBM PPC Volume
	PPCVolumeName              = "ppc_volume_name"
	PPCVolumeSize              = "ppc_volume_size"
	PPCVolumeType              = "ppc_volume_type"
	PPCVolumeShareable         = "ppc_volume_shareable"
	PPCVolumeId                = "ppc_volume_id"
	PPCVolumeStatus            = "ppc_volume_status"
	PPCVolumeWWN               = "ppc_volume_wwn"
	PPCVolumeDeleteOnTerminate = "ppc_volume_delete_on_terminate"
	PPCVolumeCreateDate        = "ppc_volume_create_date"
	PPCVolumeLastUpdate        = "ppc_last_updated_date"
	PPCVolumePool              = "ppc_volume_pool"
	PPCAffinityPolicy          = "ppc_volume_affinity_policy"
	PPCAffinityVolume          = "ppc_volume_affinity"
	PPCAffinityInstance        = "ppc_volume_affinity_instance"
	PPCAffinityDiskCount       = "ppc_volume_disk_count"
	PPCStoragePoolValue        = "ppc_storage_pool_type"
	PPCStoragePoolName         = "ppc_storage_pool_name"
	PPCReplicationEnabled      = "ppc_replication_enabled"

	// IBM PPC Snapshots

	PPCSnapshot         = "ppc_snap_shot_id"
	PPCSnapshotName     = "ppc_snap_shot_name"
	PPCSnapshotStatus   = "ppc_snap_shot_status"
	PPCSnapshotAction   = "ppc_snap_shot_action"
	PPCSnapshotComplete = "ppc_snap_shot_complete"

	// IBM PPC SAP Profile

	PPCSAPProfileID        = "ppc_sap_profile_id"
	PPCSAPProfile          = "ppc_sap_profile"
	PPCSAPProfileMemory    = "ppc_sap_profile_memory"
	PPCSAPProfileCertified = "ppc_sap_profile_certified"
	PPCSAPProfileType      = "ppc_sap_profile_type"
	PPCSAPProfileCores     = "ppc_sap_profile_cores"

	// IBM PPC Clone Volume
	PPCVolumeCloneStatus  = "ppc_volume_clone_status"
	PPCVolumeClonePercent = "ppc_volume_clone_percent"
	PPCVolumeCloneFailure = "ppc_volume_clone_failure"

	// IBM PPC Image

	PPCImageName            = "ppc_image_name"
	PPCImageId              = "ppc_image_id"
	PPCImageAccessKey       = "ppc_image_access_key"
	PPCImageSecretKey       = "ppc_image_secret_key"
	PPCImageSource          = "ppc_image_source"
	PPCImageBucketName      = "ppc_image_bucket_name"
	PPCImageBucketAccess    = "ppc_image_bucket_access"
	PPCImageBucketFileName  = "ppc_image_bucket_file_name"
	PPCImageBucketRegion    = "ppc_image_bucket_region"
	PPCImageStorageAffinity = "ppc_image_storage_affinity"
	PPCImageStoragePool     = "ppc_image_storage_pool"
	PPCImageStorageType     = "ppc_image_storage_type"
	PPCImageCopyID          = "ppc_image_copy_id"
	PPCImagePath            = "ppc_image_path"
	PPCImageOsType          = "ppc_image_os_type"

	// IBM PPC Key

	PPCKeyName = "ppc_key_name"
	PPCKey     = "ppc_ssh_key"
	PPCKeyDate = "ppc_creation_date"
	PPCKeyId   = "ppc_key_id"

	// IBM PPC Network

	PPCNetworkReady           = "ready"
	PPCNetworkID              = "ppc_networkid"
	PPCNetworkName            = "ppc_network_name"
	PPCNetworkCidr            = "ppc_cidr"
	PPCNetworkDNS             = "ppc_dns"
	PPCNetworkType            = "ppc_network_type"
	PPCNetworkGateway         = "ppc_gateway"
	PPCNetworkIPAddressRange  = "ppc_ipaddress_range"
	PPCNetworkVlanId          = "ppc_vlan_id"
	PPCNetworkProvisioning    = "build"
	PPCNetworkJumbo           = "ppc_network_jumbo"
	PPCNetworkPortDescription = "ppc_network_port_description"
	PPCNetworkPortIPAddress   = "ppc_network_port_ipaddress"
	PPCNetworkPortMacAddress  = "ppc_network_port_macaddress"
	PPCNetworkPortStatus      = "ppc_network_port_status"
	PPCNetworkPortPortID      = "ppc_network_port_portid"

	// IBM PPC Operations
	PPCInstanceOperationType       = "ppc_operation"
	PPCInstanceOperationProgress   = "ppc_progress"
	PPCInstanceOperationStatus     = "ppc_status"
	PPCInstanceOperationServerName = "ppc_instance_name"

	// IBM PPC Volume Attach
	PPCVolumeAttachName               = "ppc_volume_attach_name"
	PPCVolumeAllowableAttachStatus    = "in-use"
	PPCVolumeAttachStatus             = "status"
	PowerVolumeAttachDeleting         = "deleting"
	PowerVolumeAttachProvisioning     = "creating"
	PowerVolumeAttachProvisioningDone = "available"

	// IBM PPC Instance Capture
	PPCInstanceCaptureName                  = "ppc_capture_name"
	PPCInstanceCaptureDestination           = "ppc_capture_destination"
	PPCInstanceCaptureVolumeIds             = "ppc_capture_volume_ids"
	PPCInstanceCaptureCloudStorageImagePath = "ppc_capture_storage_image_path"
	PPCInstanceCaptureCloudStorageRegion    = "ppc_capture_cloud_storage_region"
	PPCInstanceCaptureCloudStorageAccessKey = "ppc_capture_cloud_storage_access_key"
	PPCInstanceCaptureCloudStorageSecretKey = "ppc_capture_cloud_storage_secret_key"

	// IBM PPC Cloud Connections

	PPCCloudConnectionName          = "ppc_cloud_connection_name"
	PPCCloudConnectionStatus        = "ppc_cloud_connection_status"
	PPCCloudConnectionMetered       = "ppc_cloud_connection_metered"
	PPCCloudConnectionUserIPAddress = "ppc_cloud_connection_user_ip_address"
	PPCCloudConnectionIBMIPAddress  = "ppc_cloud_connection_ibm_ip_address"
	PPCCloudConnectionSpeed         = "ppc_cloud_connection_speed"
	PPCCloudConnectionPort          = "ppc_cloud_connection_port"
	PPCCloudConnectionGlobalRouting = "ppc_cloud_connection_global_routing"
	PPCCloudConnectionId            = "ppc_cloud_connection_id"
	//PPCCloudConnectionClassic          = "ppc_cloud_connection_classic"
	PPCCloudConnectionClassicEnabled   = "ppc_cloud_connection_classic_enabled"
	PPCCloudConnectionClassicGreCidr   = "ppc_cloud_connection_gre_cidr"
	PPCCloudConnectionClassicGreDest   = "ppc_cloud_connection_gre_destination_address"
	PPCCloudConnectionClassicGreSource = "ppc_cloud_connection_gre_source_address"
	PPCCloudConnectionNetworks         = "ppc_cloud_connection_networks"
	//PPCCloudConnectionVPC              = "ppc_cloud_connection_vpc"
	PPCCloudConnectionVPCEnabled = "ppc_cloud_connection_vpc_enabled"
	PPCCloudConnectionVPCCRNs    = "ppc_cloud_connection_vpc_crns"
	PPCCloudConnectionVPCName    = "ppc_cloud_connection_vpc_name"

	// // IBM PPC VPN Connections

	// PPCVPNConnectionName                       = "ppc_vpn_connection_name"
	// PPCVPNConnectionId                         = "ppc_vpn_connection_id"
	// PPCVPNIKEPolicyId                          = "ppc_ike_policy_id"
	// PPCVPNIPSecPolicyId                        = "ppc_ipsec_policy_id"
	// PPCVPNConnectionLocalGatewayAddress        = "ppc_local_gateway_address"
	// PPCVPNConnectionMode                       = "ppc_vpn_connection_mode"
	// PPCVPNConnectionNetworks                   = "ppc_networks"
	// PPCVPNConnectionPeerGatewayAddress         = "ppc_peer_gateway_address"
	// PPCVPNConnectionPeerSubnets                = "ppc_peer_subnets"
	// PPCVPNConnectionStatus                     = "ppc_vpn_connection_status"
	// PPCVPNConnectionVpnGatewayAddress          = "ppc_gateway_address"
	// PPCVPNConnectionDeadPeerDetection          = "ppc_dead_peer_detections"
	// PPCVPNConnectionDeadPeerDetectionAction    = "ppc_dead_peer_detections_action"
	// PPCVPNConnectionDeadPeerDetectionInterval  = "ppc_dead_peer_detections_interval"
	// PPCVPNConnectionDeadPeerDetectionThreshold = "ppc_dead_peer_detections_threshold"

	// // IBM PPC VPN Policy
	// PPCVPNPolicyId             = "ppc_policy_id"
	// PPCVPNPolicyName           = "ppc_policy_name"
	// PPCVPNPolicyDhGroup        = "ppc_policy_dh_group"
	// PPCVPNPolicyEncryption     = "ppc_policy_encryption"
	// PPCVPNPolicyKeyLifetime    = "ppc_policy_key_lifetime"
	// PPCVPNPolicyPresharedKey   = "ppc_policy_preshared_key"
	// PPCVPNPolicyVersion        = "ppc_policy_version"
	// PPCVPNPolicyAuthentication = "ppc_policy_authentication"
	// PPCVPNPolicyPFS            = "ppc_policy_pfs"

	JobStatusQueued             = "queued"
	JobStatusReadyForProcessing = "readyForProcessing"
	JobStatusInProgress         = "inProgress"
	JobStatusCompleted          = "completed"
	JobStatusFailed             = "failed"
	JobStatusRunning            = "running"
	JobStatusWaiting            = "waiting"

	// IBM PPC DHCP
	PPCDhcpId          = "ppc_dhcp_id"
	PPCDhcpStatus      = "ppc_dhcp_status"
	PPCDhcpNetwork     = "ppc_dhcp_network"
	PPCDhcpLeases      = "ppc_dhcp_leases"
	PPCDhcpInstanceIp  = "ppc_dhcp_instance_ip"
	PPCDhcpInstanceMac = "ppc_dhcp_instance_mac"

	// IBM PPC Placement Groups

	PPCPlacementGroupName   = "ppc_placement_group_name"
	PPCPlacementGroupPolicy = "ppc_placement_group_policy"
	PPCPlacementGroupID     = "ppc_placement_group_id"

	// Status For all the resources

	PPCVolumeDeleting         = "deleting"
	PPCVolumeDeleted          = "done"
	PPCVolumeProvisioning     = "creating"
	PPCVolumeProvisioningDone = "available"
	PPCInstanceAvailable      = "ACTIVE"
	PPCInstanceHealthOk       = "OK"
	PPCInstanceHealthWarning  = "WARNING"
	PPCInstanceBuilding       = "BUILD"
	PPCInstanceDeleting       = "DELETING"
	PPCInstanceNotFound       = "Not Found"
	PPCImageQueStatus         = "queued"
	PPCImageActiveStatus      = "active"

	//Timeout values for Power VS -

	PPCCreateTimeOut = 5 * time.Minute
	PPCUpdateTimeOut = 5 * time.Minute
	PPCDeleteTimeOut = 3 * time.Minute
	PPCGetTimeOut    = 2 * time.Minute
)

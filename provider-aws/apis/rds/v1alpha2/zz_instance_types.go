/*
Copyright 2022 Upbound Inc.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LatestRestorableTime *string `json:"latestRestorableTime,omitempty" tf:"latest_restorable_time,omitempty"`

	Replicas []*string `json:"replicas,omitempty" tf:"replicas,omitempty"`

	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type InstanceParameters struct {

	// +kubebuilder:validation:Optional
	AllocatedStorage *float64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// +kubebuilder:validation:Optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	BackupWindow *string `json:"backupWindow,omitempty" tf:"backup_window,omitempty"`

	// +kubebuilder:validation:Optional
	CACertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	CharacterSetName *string `json:"characterSetName,omitempty" tf:"character_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerOwnedIPEnabled *bool `json:"customerOwnedIpEnabled,omitempty" tf:"customer_owned_ip_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// +crossplane:generate:reference:type=SubnetGroup
	// +kubebuilder:validation:Optional
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	DBSubnetGroupNameRef *v1.Reference `json:"dbSubnetGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DBSubnetGroupNameSelector *v1.Selector `json:"dbSubnetGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups,omitempty" tf:"delete_automated_backups,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	DomainIAMRoleName *string `json:"domainIamRoleName,omitempty" tf:"domain_iam_role_name,omitempty"`

	// +kubebuilder:validation:Optional
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`

	// +kubebuilder:validation:Required
	InstanceClass *string `json:"instanceClass" tf:"instance_class,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage,omitempty" tf:"max_allocated_storage,omitempty"`

	// +kubebuilder:validation:Optional
	MonitoringInterval *float64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`

	// +kubebuilder:validation:Optional
	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NcharCharacterSetName *string `json:"ncharCharacterSetName,omitempty" tf:"nchar_character_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`

	// +crossplane:generate:reference:type=ParameterGroup
	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ParameterGroupNameRef *v1.Reference `json:"parameterGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ParameterGroupNameSelector *v1.Selector `json:"parameterGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyID *string `json:"performanceInsightsKmsKeyId,omitempty" tf:"performance_insights_kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDRef *v1.Reference `json:"performanceInsightsKmsKeyIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDSelector *v1.Selector `json:"performanceInsightsKmsKeyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicaMode *string `json:"replicaMode,omitempty" tf:"replica_mode,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicateSourceDB *string `json:"replicateSourceDb,omitempty" tf:"replicate_source_db,omitempty"`

	// +kubebuilder:validation:Optional
	RestoreToPointInTime []RestoreToPointInTimeParameters `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time,omitempty"`

	// +kubebuilder:validation:Optional
	S3Import []S3ImportParameters `json:"s3Import,omitempty" tf:"s3_import,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupNameRefs []v1.Reference `json:"securityGroupNameRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupNameSelector *v1.Selector `json:"securityGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1alpha2.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupNameRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupNameSelector
	// +kubebuilder:validation:Optional
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1alpha2.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VpcSecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=VpcSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIdRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcSecurityGroupIdSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`
}

type RestoreToPointInTimeObservation struct {
}

type RestoreToPointInTimeParameters struct {

	// +kubebuilder:validation:Optional
	RestoreTime *string `json:"restoreTime,omitempty" tf:"restore_time,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDBInstanceAutomatedBackupsArn *string `json:"sourceDbInstanceAutomatedBackupsArn,omitempty" tf:"source_db_instance_automated_backups_arn,omitempty"`

	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	SourceDBInstanceIdentifier *string `json:"sourceDbInstanceIdentifier,omitempty" tf:"source_db_instance_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDBInstanceIdentifierRef *v1.Reference `json:"sourceDbInstanceIdentifierRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceDBInstanceIdentifierSelector *v1.Selector `json:"sourceDbInstanceIdentifierSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceDbiResourceID *string `json:"sourceDbiResourceId,omitempty" tf:"source_dbi_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time,omitempty"`
}

type S3ImportObservation struct {
}

type S3ImportParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1alpha2.Bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// +kubebuilder:validation:Required
	IngestionRole *string `json:"ingestionRole" tf:"ingestion_role,omitempty"`

	// +kubebuilder:validation:Required
	SourceEngine *string `json:"sourceEngine" tf:"source_engine,omitempty"`

	// +kubebuilder:validation:Required
	SourceEngineVersion *string `json:"sourceEngineVersion" tf:"source_engine_version,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}

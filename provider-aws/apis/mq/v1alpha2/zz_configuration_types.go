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

type ConfigurationObservation_2 struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LatestRevision *float64 `json:"latestRevision,omitempty" tf:"latest_revision,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConfigurationParameters_2 struct {

	// +kubebuilder:validation:Optional
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy,omitempty"`

	// +kubebuilder:validation:Required
	Data *string `json:"data" tf:"data,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	EngineType *string `json:"engineType" tf:"engine_type,omitempty"`

	// +kubebuilder:validation:Required
	EngineVersion *string `json:"engineVersion" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationParameters_2 `json:"forProvider"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Configuration is the Schema for the Configurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationSpec   `json:"spec"`
	Status            ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	Configuration_Kind             = "Configuration"
	Configuration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Configuration_Kind}.String()
	Configuration_KindAPIVersion   = Configuration_Kind + "." + CRDGroupVersion.String()
	Configuration_GroupVersionKind = CRDGroupVersion.WithKind(Configuration_Kind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}

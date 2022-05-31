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

type HostedZoneDNSSECObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostedZoneDNSSECParameters struct {

	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	HostedZoneIDRef *v1.Reference `json:"hostedZoneIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HostedZoneIDSelector *v1.Selector `json:"hostedZoneIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SigningStatus *string `json:"signingStatus,omitempty" tf:"signing_status,omitempty"`
}

// HostedZoneDNSSECSpec defines the desired state of HostedZoneDNSSEC
type HostedZoneDNSSECSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedZoneDNSSECParameters `json:"forProvider"`
}

// HostedZoneDNSSECStatus defines the observed state of HostedZoneDNSSEC.
type HostedZoneDNSSECStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedZoneDNSSECObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedZoneDNSSEC is the Schema for the HostedZoneDNSSECs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HostedZoneDNSSEC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedZoneDNSSECSpec   `json:"spec"`
	Status            HostedZoneDNSSECStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedZoneDNSSECList contains a list of HostedZoneDNSSECs
type HostedZoneDNSSECList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedZoneDNSSEC `json:"items"`
}

// Repository type metadata.
var (
	HostedZoneDNSSEC_Kind             = "HostedZoneDNSSEC"
	HostedZoneDNSSEC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedZoneDNSSEC_Kind}.String()
	HostedZoneDNSSEC_KindAPIVersion   = HostedZoneDNSSEC_Kind + "." + CRDGroupVersion.String()
	HostedZoneDNSSEC_GroupVersionKind = CRDGroupVersion.WithKind(HostedZoneDNSSEC_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedZoneDNSSEC{}, &HostedZoneDNSSECList{})
}

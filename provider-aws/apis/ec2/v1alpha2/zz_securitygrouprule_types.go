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

type SecurityGroupRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityGroupRuleParameters struct {

	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	FromPort *float64 `json:"fromPort" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks,omitempty"`

	// +kubebuilder:validation:Optional
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self,omitempty"`

	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SourceSecurityGroupID *string `json:"sourceSecurityGroupId,omitempty" tf:"source_security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDRef *v1.Reference `json:"sourceSecurityGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupIDSelector *v1.Selector `json:"sourceSecurityGroupIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ToPort *float64 `json:"toPort" tf:"to_port,omitempty"`

	// Type of rule, ingress (inbound) or egress (outbound).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// SecurityGroupRuleSpec defines the desired state of SecurityGroupRule
type SecurityGroupRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupRuleParameters `json:"forProvider"`
}

// SecurityGroupRuleStatus defines the observed state of SecurityGroupRule.
type SecurityGroupRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRule is the Schema for the SecurityGroupRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupRuleSpec   `json:"spec"`
	Status            SecurityGroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRuleList contains a list of SecurityGroupRules
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupRule `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroupRule_Kind             = "SecurityGroupRule"
	SecurityGroupRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroupRule_Kind}.String()
	SecurityGroupRule_KindAPIVersion   = SecurityGroupRule_Kind + "." + CRDGroupVersion.String()
	SecurityGroupRule_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroupRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroupRule{}, &SecurityGroupRuleList{})
}

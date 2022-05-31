/*
Copyright 2022 Upbound Inc.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FirewallRuleParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	BlockOverrideDNSType *string `json:"blockOverrideDnsType,omitempty" tf:"block_override_dns_type,omitempty"`

	// +kubebuilder:validation:Optional
	BlockOverrideDomain *string `json:"blockOverrideDomain,omitempty" tf:"block_override_domain,omitempty"`

	// +kubebuilder:validation:Optional
	BlockOverrideTTL *float64 `json:"blockOverrideTtl,omitempty" tf:"block_override_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	BlockResponse *string `json:"blockResponse,omitempty" tf:"block_response,omitempty"`

	// +kubebuilder:validation:Required
	FirewallDomainListID *string `json:"firewallDomainListId" tf:"firewall_domain_list_id,omitempty"`

	// +kubebuilder:validation:Required
	FirewallRuleGroupID *string `json:"firewallRuleGroupId" tf:"firewall_rule_group_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// FirewallRuleSpec defines the desired state of FirewallRule
type FirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRuleParameters `json:"forProvider"`
}

// FirewallRuleStatus defines the observed state of FirewallRule.
type FirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRule is the Schema for the FirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallRuleSpec   `json:"spec"`
	Status            FirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleList contains a list of FirewallRules
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FirewallRule_Kind             = "FirewallRule"
	FirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRule_Kind}.String()
	FirewallRule_KindAPIVersion   = FirewallRule_Kind + "." + CRDGroupVersion.String()
	FirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRule{}, &FirewallRuleList{})
}

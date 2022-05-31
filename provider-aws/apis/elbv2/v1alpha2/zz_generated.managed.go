/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LB.
func (mg *LB) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LB.
func (mg *LB) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LB.
func (mg *LB) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LB.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LB) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LB.
func (mg *LB) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LB.
func (mg *LB) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LB.
func (mg *LB) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LB.
func (mg *LB) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LB.
func (mg *LB) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LB.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LB) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LB.
func (mg *LB) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LB.
func (mg *LB) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LBListener.
func (mg *LBListener) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LBListener.
func (mg *LBListener) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LBListener.
func (mg *LBListener) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LBListener.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LBListener) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LBListener.
func (mg *LBListener) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LBListener.
func (mg *LBListener) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LBListener.
func (mg *LBListener) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LBListener.
func (mg *LBListener) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LBListener.
func (mg *LBListener) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LBListener.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LBListener) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LBListener.
func (mg *LBListener) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LBListener.
func (mg *LBListener) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LBTargetGroup.
func (mg *LBTargetGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LBTargetGroup.
func (mg *LBTargetGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LBTargetGroup.
func (mg *LBTargetGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LBTargetGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LBTargetGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LBTargetGroup.
func (mg *LBTargetGroup) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LBTargetGroup.
func (mg *LBTargetGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LBTargetGroup.
func (mg *LBTargetGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LBTargetGroup.
func (mg *LBTargetGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LBTargetGroup.
func (mg *LBTargetGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LBTargetGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LBTargetGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LBTargetGroup.
func (mg *LBTargetGroup) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LBTargetGroup.
func (mg *LBTargetGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LBTargetGroupAttachment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LBTargetGroupAttachment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LBTargetGroupAttachment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LBTargetGroupAttachment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LBTargetGroupAttachment.
func (mg *LBTargetGroupAttachment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

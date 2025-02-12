/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AssessmentTargetInitParameters struct {

	// The name of the assessment target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AssessmentTargetObservation struct {

	// The target assessment ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the assessment target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
	ResourceGroupArn *string `json:"resourceGroupArn,omitempty" tf:"resource_group_arn,omitempty"`
}

type AssessmentTargetParameters struct {

	// The name of the assessment target.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/inspector/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceGroupArn *string `json:"resourceGroupArn,omitempty" tf:"resource_group_arn,omitempty"`

	// Reference to a ResourceGroup in inspector to populate resourceGroupArn.
	// +kubebuilder:validation:Optional
	ResourceGroupArnRef *v1.Reference `json:"resourceGroupArnRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in inspector to populate resourceGroupArn.
	// +kubebuilder:validation:Optional
	ResourceGroupArnSelector *v1.Selector `json:"resourceGroupArnSelector,omitempty" tf:"-"`
}

// AssessmentTargetSpec defines the desired state of AssessmentTarget
type AssessmentTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssessmentTargetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AssessmentTargetInitParameters `json:"initProvider,omitempty"`
}

// AssessmentTargetStatus defines the observed state of AssessmentTarget.
type AssessmentTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssessmentTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AssessmentTarget is the Schema for the AssessmentTargets API. Provides an Inspector Classic Assessment Target.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AssessmentTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   AssessmentTargetSpec   `json:"spec"`
	Status AssessmentTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssessmentTargetList contains a list of AssessmentTargets
type AssessmentTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AssessmentTarget `json:"items"`
}

// Repository type metadata.
var (
	AssessmentTarget_Kind             = "AssessmentTarget"
	AssessmentTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AssessmentTarget_Kind}.String()
	AssessmentTarget_KindAPIVersion   = AssessmentTarget_Kind + "." + CRDGroupVersion.String()
	AssessmentTarget_GroupVersionKind = CRDGroupVersion.WithKind(AssessmentTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&AssessmentTarget{}, &AssessmentTargetList{})
}

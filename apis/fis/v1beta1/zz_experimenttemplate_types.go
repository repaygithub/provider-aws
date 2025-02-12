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

type ActionInitParameters struct {

	// ID of the action. To find out what actions are supported see AWS FIS actions reference.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Description of the action.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Friendly name of the action.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Parameter(s) for the action, if applicable. See below.
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Set of action names that must complete before this action can be executed.
	StartAfter []*string `json:"startAfter,omitempty" tf:"start_after,omitempty"`

	// Action's target, if applicable. See below.
	Target []TargetInitParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type ActionObservation struct {

	// ID of the action. To find out what actions are supported see AWS FIS actions reference.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Description of the action.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Friendly name of the action.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Parameter(s) for the action, if applicable. See below.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Set of action names that must complete before this action can be executed.
	StartAfter []*string `json:"startAfter,omitempty" tf:"start_after,omitempty"`

	// Action's target, if applicable. See below.
	Target []TargetObservation `json:"target,omitempty" tf:"target,omitempty"`
}

type ActionParameters struct {

	// ID of the action. To find out what actions are supported see AWS FIS actions reference.
	// +kubebuilder:validation:Optional
	ActionID *string `json:"actionId" tf:"action_id,omitempty"`

	// Description of the action.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Friendly name of the action.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter(s) for the action, if applicable. See below.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Set of action names that must complete before this action can be executed.
	// +kubebuilder:validation:Optional
	StartAfter []*string `json:"startAfter,omitempty" tf:"start_after,omitempty"`

	// Action's target, if applicable. See below.
	// +kubebuilder:validation:Optional
	Target []TargetParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type ExperimentTemplateInitParameters struct {

	// Action to be performed during an experiment. See below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Description for the experiment template.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// When an ongoing experiment should be stopped. See below.
	StopCondition []StopConditionInitParameters `json:"stopCondition,omitempty" tf:"stop_condition,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Target of an action. See below.
	Target []ExperimentTemplateTargetInitParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type ExperimentTemplateObservation struct {

	// Action to be performed during an experiment. See below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// Description for the experiment template.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Experiment Template ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// When an ongoing experiment should be stopped. See below.
	StopCondition []StopConditionObservation `json:"stopCondition,omitempty" tf:"stop_condition,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Target of an action. See below.
	Target []ExperimentTemplateTargetObservation `json:"target,omitempty" tf:"target,omitempty"`
}

type ExperimentTemplateParameters struct {

	// Action to be performed during an experiment. See below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Description for the experiment template.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// When an ongoing experiment should be stopped. See below.
	// +kubebuilder:validation:Optional
	StopCondition []StopConditionParameters `json:"stopCondition,omitempty" tf:"stop_condition,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Target of an action. See below.
	// +kubebuilder:validation:Optional
	Target []ExperimentTemplateTargetParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type ExperimentTemplateTargetInitParameters struct {

	// Filter(s) for the target. Filters can be used to select resources based on specific attributes returned by the respective describe action of the resource type. For more information, see Targets for AWS FIS. See below.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Friendly name given to the target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of ARNs of the resources to target with an action. Conflicts with resource_tag.
	ResourceArns []*string `json:"resourceArns,omitempty" tf:"resource_arns,omitempty"`

	// Tag(s) the resources need to have to be considered a valid target for an action. Conflicts with resource_arns. See below.
	ResourceTag []ResourceTagInitParameters `json:"resourceTag,omitempty" tf:"resource_tag,omitempty"`

	// AWS resource type. The resource type must be supported for the specified action. To find out what resource types are supported, see Targets for AWS FIS.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Scopes the identified resources. Valid values are ALL (all identified resources), COUNT(n) (randomly select n of the identified resources), PERCENT(n) (randomly select n percent of the identified resources).
	SelectionMode *string `json:"selectionMode,omitempty" tf:"selection_mode,omitempty"`
}

type ExperimentTemplateTargetObservation struct {

	// Filter(s) for the target. Filters can be used to select resources based on specific attributes returned by the respective describe action of the resource type. For more information, see Targets for AWS FIS. See below.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// Friendly name given to the target.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of ARNs of the resources to target with an action. Conflicts with resource_tag.
	ResourceArns []*string `json:"resourceArns,omitempty" tf:"resource_arns,omitempty"`

	// Tag(s) the resources need to have to be considered a valid target for an action. Conflicts with resource_arns. See below.
	ResourceTag []ResourceTagObservation `json:"resourceTag,omitempty" tf:"resource_tag,omitempty"`

	// AWS resource type. The resource type must be supported for the specified action. To find out what resource types are supported, see Targets for AWS FIS.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Scopes the identified resources. Valid values are ALL (all identified resources), COUNT(n) (randomly select n of the identified resources), PERCENT(n) (randomly select n percent of the identified resources).
	SelectionMode *string `json:"selectionMode,omitempty" tf:"selection_mode,omitempty"`
}

type ExperimentTemplateTargetParameters struct {

	// Filter(s) for the target. Filters can be used to select resources based on specific attributes returned by the respective describe action of the resource type. For more information, see Targets for AWS FIS. See below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Friendly name given to the target.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Set of ARNs of the resources to target with an action. Conflicts with resource_tag.
	// +kubebuilder:validation:Optional
	ResourceArns []*string `json:"resourceArns,omitempty" tf:"resource_arns,omitempty"`

	// Tag(s) the resources need to have to be considered a valid target for an action. Conflicts with resource_arns. See below.
	// +kubebuilder:validation:Optional
	ResourceTag []ResourceTagParameters `json:"resourceTag,omitempty" tf:"resource_tag,omitempty"`

	// AWS resource type. The resource type must be supported for the specified action. To find out what resource types are supported, see Targets for AWS FIS.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`

	// Scopes the identified resources. Valid values are ALL (all identified resources), COUNT(n) (randomly select n of the identified resources), PERCENT(n) (randomly select n percent of the identified resources).
	// +kubebuilder:validation:Optional
	SelectionMode *string `json:"selectionMode" tf:"selection_mode,omitempty"`
}

type FilterInitParameters struct {

	// Attribute path for the filter.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Set of attribute values for the filter.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterObservation struct {

	// Attribute path for the filter.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Set of attribute values for the filter.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterParameters struct {

	// Attribute path for the filter.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`

	// Set of attribute values for the filter.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ParameterInitParameters struct {

	// Parameter name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterObservation struct {

	// Parameter name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// Parameter name.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Parameter value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type ResourceTagInitParameters struct {

	// Tag key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Tag value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourceTagObservation struct {

	// Tag key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Tag value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourceTagParameters struct {

	// Tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Tag value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type StopConditionInitParameters struct {

	// Source of the condition. One of none, aws:cloudwatch:alarm.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// ARN of the CloudWatch alarm. Required if the source is a CloudWatch alarm.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StopConditionObservation struct {

	// Source of the condition. One of none, aws:cloudwatch:alarm.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// ARN of the CloudWatch alarm. Required if the source is a CloudWatch alarm.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StopConditionParameters struct {

	// Source of the condition. One of none, aws:cloudwatch:alarm.
	// +kubebuilder:validation:Optional
	Source *string `json:"source" tf:"source,omitempty"`

	// ARN of the CloudWatch alarm. Required if the source is a CloudWatch alarm.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TargetInitParameters struct {

	// Tag key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Target name, referencing a corresponding target.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TargetObservation struct {

	// Tag key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Target name, referencing a corresponding target.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TargetParameters struct {

	// Tag key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Target name, referencing a corresponding target.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// ExperimentTemplateSpec defines the desired state of ExperimentTemplate
type ExperimentTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExperimentTemplateParameters `json:"forProvider"`
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
	InitProvider ExperimentTemplateInitParameters `json:"initProvider,omitempty"`
}

// ExperimentTemplateStatus defines the observed state of ExperimentTemplate.
type ExperimentTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExperimentTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExperimentTemplate is the Schema for the ExperimentTemplates API. Provides an FIS Experiment Template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ExperimentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || has(self.initProvider.action)",message="action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || has(self.initProvider.description)",message="description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stopCondition) || has(self.initProvider.stopCondition)",message="stopCondition is a required parameter"
	Spec   ExperimentTemplateSpec   `json:"spec"`
	Status ExperimentTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExperimentTemplateList contains a list of ExperimentTemplates
type ExperimentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExperimentTemplate `json:"items"`
}

// Repository type metadata.
var (
	ExperimentTemplate_Kind             = "ExperimentTemplate"
	ExperimentTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExperimentTemplate_Kind}.String()
	ExperimentTemplate_KindAPIVersion   = ExperimentTemplate_Kind + "." + CRDGroupVersion.String()
	ExperimentTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ExperimentTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ExperimentTemplate{}, &ExperimentTemplateList{})
}

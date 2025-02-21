package awsautoscalingplans

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscalingplans/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AutoScalingPlans::ScalingPlan`.
type CfnScalingPlan interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationSource() interface{}
	SetApplicationSource(val interface{})
	AttrScalingPlanName() *string
	AttrScalingPlanVersion() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ScalingInstructions() interface{}
	SetScalingInstructions(val interface{})
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScalingPlan
type jsiiProxy_CfnScalingPlan struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScalingPlan) ApplicationSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) AttrScalingPlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrScalingPlanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) AttrScalingPlanVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrScalingPlanVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) ScalingInstructions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingInstructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScalingPlans::ScalingPlan`.
func NewCfnScalingPlan(scope constructs.Construct, id *string, props *CfnScalingPlanProps) CfnScalingPlan {
	_init_.Initialize()

	j := jsiiProxy_CfnScalingPlan{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScalingPlans::ScalingPlan`.
func NewCfnScalingPlan_Override(c CfnScalingPlan, scope constructs.Construct, id *string, props *CfnScalingPlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScalingPlan) SetApplicationSource(val interface{}) {
	_jsii_.Set(
		j,
		"applicationSource",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPlan) SetScalingInstructions(val interface{}) {
	_jsii_.Set(
		j,
		"scalingInstructions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnScalingPlan_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnScalingPlan_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnScalingPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalingPlan_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScalingPlan) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnScalingPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnScalingPlan) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnScalingPlan_ApplicationSourceProperty struct {
	// `CfnScalingPlan.ApplicationSourceProperty.CloudFormationStackARN`.
	CloudFormationStackArn *string `json:"cloudFormationStackArn"`
	// `CfnScalingPlan.ApplicationSourceProperty.TagFilters`.
	TagFilters interface{} `json:"tagFilters"`
}

type CfnScalingPlan_CustomizedLoadMetricSpecificationProperty struct {
	// `CfnScalingPlan.CustomizedLoadMetricSpecificationProperty.MetricName`.
	MetricName *string `json:"metricName"`
	// `CfnScalingPlan.CustomizedLoadMetricSpecificationProperty.Namespace`.
	Namespace *string `json:"namespace"`
	// `CfnScalingPlan.CustomizedLoadMetricSpecificationProperty.Statistic`.
	Statistic *string `json:"statistic"`
	// `CfnScalingPlan.CustomizedLoadMetricSpecificationProperty.Dimensions`.
	Dimensions interface{} `json:"dimensions"`
	// `CfnScalingPlan.CustomizedLoadMetricSpecificationProperty.Unit`.
	Unit *string `json:"unit"`
}

type CfnScalingPlan_CustomizedScalingMetricSpecificationProperty struct {
	// `CfnScalingPlan.CustomizedScalingMetricSpecificationProperty.MetricName`.
	MetricName *string `json:"metricName"`
	// `CfnScalingPlan.CustomizedScalingMetricSpecificationProperty.Namespace`.
	Namespace *string `json:"namespace"`
	// `CfnScalingPlan.CustomizedScalingMetricSpecificationProperty.Statistic`.
	Statistic *string `json:"statistic"`
	// `CfnScalingPlan.CustomizedScalingMetricSpecificationProperty.Dimensions`.
	Dimensions interface{} `json:"dimensions"`
	// `CfnScalingPlan.CustomizedScalingMetricSpecificationProperty.Unit`.
	Unit *string `json:"unit"`
}

type CfnScalingPlan_MetricDimensionProperty struct {
	// `CfnScalingPlan.MetricDimensionProperty.Name`.
	Name *string `json:"name"`
	// `CfnScalingPlan.MetricDimensionProperty.Value`.
	Value *string `json:"value"`
}

type CfnScalingPlan_PredefinedLoadMetricSpecificationProperty struct {
	// `CfnScalingPlan.PredefinedLoadMetricSpecificationProperty.PredefinedLoadMetricType`.
	PredefinedLoadMetricType *string `json:"predefinedLoadMetricType"`
	// `CfnScalingPlan.PredefinedLoadMetricSpecificationProperty.ResourceLabel`.
	ResourceLabel *string `json:"resourceLabel"`
}

type CfnScalingPlan_PredefinedScalingMetricSpecificationProperty struct {
	// `CfnScalingPlan.PredefinedScalingMetricSpecificationProperty.PredefinedScalingMetricType`.
	PredefinedScalingMetricType *string `json:"predefinedScalingMetricType"`
	// `CfnScalingPlan.PredefinedScalingMetricSpecificationProperty.ResourceLabel`.
	ResourceLabel *string `json:"resourceLabel"`
}

type CfnScalingPlan_ScalingInstructionProperty struct {
	// `CfnScalingPlan.ScalingInstructionProperty.MaxCapacity`.
	MaxCapacity *float64 `json:"maxCapacity"`
	// `CfnScalingPlan.ScalingInstructionProperty.MinCapacity`.
	MinCapacity *float64 `json:"minCapacity"`
	// `CfnScalingPlan.ScalingInstructionProperty.ResourceId`.
	ResourceId *string `json:"resourceId"`
	// `CfnScalingPlan.ScalingInstructionProperty.ScalableDimension`.
	ScalableDimension *string `json:"scalableDimension"`
	// `CfnScalingPlan.ScalingInstructionProperty.ServiceNamespace`.
	ServiceNamespace *string `json:"serviceNamespace"`
	// `CfnScalingPlan.ScalingInstructionProperty.TargetTrackingConfigurations`.
	TargetTrackingConfigurations interface{} `json:"targetTrackingConfigurations"`
	// `CfnScalingPlan.ScalingInstructionProperty.CustomizedLoadMetricSpecification`.
	CustomizedLoadMetricSpecification interface{} `json:"customizedLoadMetricSpecification"`
	// `CfnScalingPlan.ScalingInstructionProperty.DisableDynamicScaling`.
	DisableDynamicScaling interface{} `json:"disableDynamicScaling"`
	// `CfnScalingPlan.ScalingInstructionProperty.PredefinedLoadMetricSpecification`.
	PredefinedLoadMetricSpecification interface{} `json:"predefinedLoadMetricSpecification"`
	// `CfnScalingPlan.ScalingInstructionProperty.PredictiveScalingMaxCapacityBehavior`.
	PredictiveScalingMaxCapacityBehavior *string `json:"predictiveScalingMaxCapacityBehavior"`
	// `CfnScalingPlan.ScalingInstructionProperty.PredictiveScalingMaxCapacityBuffer`.
	PredictiveScalingMaxCapacityBuffer *float64 `json:"predictiveScalingMaxCapacityBuffer"`
	// `CfnScalingPlan.ScalingInstructionProperty.PredictiveScalingMode`.
	PredictiveScalingMode *string `json:"predictiveScalingMode"`
	// `CfnScalingPlan.ScalingInstructionProperty.ScalingPolicyUpdateBehavior`.
	ScalingPolicyUpdateBehavior *string `json:"scalingPolicyUpdateBehavior"`
	// `CfnScalingPlan.ScalingInstructionProperty.ScheduledActionBufferTime`.
	ScheduledActionBufferTime *float64 `json:"scheduledActionBufferTime"`
}

type CfnScalingPlan_TagFilterProperty struct {
	// `CfnScalingPlan.TagFilterProperty.Key`.
	Key *string `json:"key"`
	// `CfnScalingPlan.TagFilterProperty.Values`.
	Values *[]*string `json:"values"`
}

type CfnScalingPlan_TargetTrackingConfigurationProperty struct {
	// `CfnScalingPlan.TargetTrackingConfigurationProperty.TargetValue`.
	TargetValue *float64 `json:"targetValue"`
	// `CfnScalingPlan.TargetTrackingConfigurationProperty.CustomizedScalingMetricSpecification`.
	CustomizedScalingMetricSpecification interface{} `json:"customizedScalingMetricSpecification"`
	// `CfnScalingPlan.TargetTrackingConfigurationProperty.DisableScaleIn`.
	DisableScaleIn interface{} `json:"disableScaleIn"`
	// `CfnScalingPlan.TargetTrackingConfigurationProperty.EstimatedInstanceWarmup`.
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup"`
	// `CfnScalingPlan.TargetTrackingConfigurationProperty.PredefinedScalingMetricSpecification`.
	PredefinedScalingMetricSpecification interface{} `json:"predefinedScalingMetricSpecification"`
	// `CfnScalingPlan.TargetTrackingConfigurationProperty.ScaleInCooldown`.
	ScaleInCooldown *float64 `json:"scaleInCooldown"`
	// `CfnScalingPlan.TargetTrackingConfigurationProperty.ScaleOutCooldown`.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown"`
}

// Properties for defining a `AWS::AutoScalingPlans::ScalingPlan`.
type CfnScalingPlanProps struct {
	// `AWS::AutoScalingPlans::ScalingPlan.ApplicationSource`.
	ApplicationSource interface{} `json:"applicationSource"`
	// `AWS::AutoScalingPlans::ScalingPlan.ScalingInstructions`.
	ScalingInstructions interface{} `json:"scalingInstructions"`
}


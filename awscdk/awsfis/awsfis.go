package awsfis

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::FIS::ExperimentTemplate`.
type CfnExperimentTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Actions() interface{}
	SetActions(val interface{})
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Stack() awscdk.Stack
	StopConditions() interface{}
	SetStopConditions(val interface{})
	Tags() interface{}
	SetTags(val interface{})
	Targets() interface{}
	SetTargets(val interface{})
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

// The jsii proxy struct for CfnExperimentTemplate
type jsiiProxy_CfnExperimentTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnExperimentTemplate) Actions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) StopConditions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::FIS::ExperimentTemplate`.
func NewCfnExperimentTemplate(scope constructs.Construct, id *string, props *CfnExperimentTemplateProps) CfnExperimentTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnExperimentTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FIS::ExperimentTemplate`.
func NewCfnExperimentTemplate_Override(c CfnExperimentTemplate, scope constructs.Construct, id *string, props *CfnExperimentTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetActions(val interface{}) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetStopConditions(val interface{}) {
	_jsii_.Set(
		j,
		"stopConditions",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
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
func CfnExperimentTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnExperimentTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
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
func CfnExperimentTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExperimentTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnExperimentTemplate) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnExperimentTemplate) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnExperimentTemplate) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnExperimentTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnExperimentTemplate) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnExperimentTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnExperimentTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnExperimentTemplate) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnExperimentTemplate) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnExperimentTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnExperimentTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnExperimentTemplate) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnExperimentTemplate) ToString() *string {
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
func (c *jsiiProxy_CfnExperimentTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnExperimentTemplate_ExperimentTemplateActionItemParameterMapProperty struct {
}

type CfnExperimentTemplate_ExperimentTemplateActionItemStartAfterListProperty struct {
	// `CfnExperimentTemplate.ExperimentTemplateActionItemStartAfterListProperty.ExperimentTemplateActionItemStartAfterList`.
	ExperimentTemplateActionItemStartAfterList *[]*string `json:"experimentTemplateActionItemStartAfterList"`
}

type CfnExperimentTemplate_ExperimentTemplateActionItemTargetMapProperty struct {
}

type CfnExperimentTemplate_ExperimentTemplateActionProperty struct {
	// `CfnExperimentTemplate.ExperimentTemplateActionProperty.actionId`.
	ActionId *string `json:"actionId"`
	// `CfnExperimentTemplate.ExperimentTemplateActionProperty.description`.
	Description *string `json:"description"`
	// `CfnExperimentTemplate.ExperimentTemplateActionProperty.parameters`.
	Parameters interface{} `json:"parameters"`
	// `CfnExperimentTemplate.ExperimentTemplateActionProperty.startAfter`.
	StartAfter interface{} `json:"startAfter"`
	// `CfnExperimentTemplate.ExperimentTemplateActionProperty.targets`.
	Targets interface{} `json:"targets"`
}

type CfnExperimentTemplate_ExperimentTemplateStopConditionProperty struct {
	// `CfnExperimentTemplate.ExperimentTemplateStopConditionProperty.source`.
	Source *string `json:"source"`
	// `CfnExperimentTemplate.ExperimentTemplateStopConditionProperty.value`.
	Value *string `json:"value"`
}

type CfnExperimentTemplate_ExperimentTemplateTargetFilterListProperty struct {
	// `CfnExperimentTemplate.ExperimentTemplateTargetFilterListProperty.ExperimentTemplateTargetFilterList`.
	ExperimentTemplateTargetFilterList interface{} `json:"experimentTemplateTargetFilterList"`
}

type CfnExperimentTemplate_ExperimentTemplateTargetFilterProperty struct {
	// `CfnExperimentTemplate.ExperimentTemplateTargetFilterProperty.path`.
	Path *string `json:"path"`
	// `CfnExperimentTemplate.ExperimentTemplateTargetFilterProperty.values`.
	Values interface{} `json:"values"`
}

type CfnExperimentTemplate_ExperimentTemplateTargetFilterValuesProperty struct {
	// `CfnExperimentTemplate.ExperimentTemplateTargetFilterValuesProperty.ExperimentTemplateTargetFilterValues`.
	ExperimentTemplateTargetFilterValues *[]*string `json:"experimentTemplateTargetFilterValues"`
}

type CfnExperimentTemplate_ExperimentTemplateTargetProperty struct {
	// `CfnExperimentTemplate.ExperimentTemplateTargetProperty.resourceType`.
	ResourceType *string `json:"resourceType"`
	// `CfnExperimentTemplate.ExperimentTemplateTargetProperty.selectionMode`.
	SelectionMode *string `json:"selectionMode"`
	// `CfnExperimentTemplate.ExperimentTemplateTargetProperty.filters`.
	Filters interface{} `json:"filters"`
	// `CfnExperimentTemplate.ExperimentTemplateTargetProperty.resourceArns`.
	ResourceArns interface{} `json:"resourceArns"`
	// `CfnExperimentTemplate.ExperimentTemplateTargetProperty.resourceTags`.
	ResourceTags interface{} `json:"resourceTags"`
}

type CfnExperimentTemplate_ResourceArnListProperty struct {
	// `CfnExperimentTemplate.ResourceArnListProperty.ResourceArnList`.
	ResourceArnList *[]*string `json:"resourceArnList"`
}

type CfnExperimentTemplate_TagMapProperty struct {
}

// Properties for defining a `AWS::FIS::ExperimentTemplate`.
type CfnExperimentTemplateProps struct {
	// `AWS::FIS::ExperimentTemplate.description`.
	Description *string `json:"description"`
	// `AWS::FIS::ExperimentTemplate.roleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::FIS::ExperimentTemplate.stopConditions`.
	StopConditions interface{} `json:"stopConditions"`
	// `AWS::FIS::ExperimentTemplate.tags`.
	Tags interface{} `json:"tags"`
	// `AWS::FIS::ExperimentTemplate.targets`.
	Targets interface{} `json:"targets"`
	// `AWS::FIS::ExperimentTemplate.actions`.
	Actions interface{} `json:"actions"`
}


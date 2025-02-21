package awswafv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::WAFv2::IPSet`.
type CfnIPSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Addresses() *[]*string
	SetAddresses(val *[]*string)
	AttrArn() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	IpAddressVersion() *string
	SetIpAddressVersion(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Scope() *string
	SetScope(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnIPSet
type jsiiProxy_CfnIPSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIPSet) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) IpAddressVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::IPSet`.
func NewCfnIPSet(scope constructs.Construct, id *string, props *CfnIPSetProps) CfnIPSet {
	_init_.Initialize()

	j := jsiiProxy_CfnIPSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnIPSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::IPSet`.
func NewCfnIPSet_Override(c CfnIPSet, scope constructs.Construct, id *string, props *CfnIPSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnIPSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIPSet) SetAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"addresses",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetIpAddressVersion(val *string) {
	_jsii_.Set(
		j,
		"ipAddressVersion",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
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
func CfnIPSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnIPSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIPSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnIPSet",
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
func CfnIPSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnIPSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIPSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafv2.CfnIPSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnIPSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnIPSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnIPSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnIPSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnIPSet) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnIPSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnIPSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnIPSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnIPSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnIPSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnIPSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIPSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnIPSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnIPSet) ToString() *string {
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
func (c *jsiiProxy_CfnIPSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::WAFv2::IPSet`.
type CfnIPSetProps struct {
	// `AWS::WAFv2::IPSet.Addresses`.
	Addresses *[]*string `json:"addresses"`
	// `AWS::WAFv2::IPSet.IPAddressVersion`.
	IpAddressVersion *string `json:"ipAddressVersion"`
	// `AWS::WAFv2::IPSet.Scope`.
	Scope *string `json:"scope"`
	// `AWS::WAFv2::IPSet.Description`.
	Description *string `json:"description"`
	// `AWS::WAFv2::IPSet.Name`.
	Name *string `json:"name"`
	// `AWS::WAFv2::IPSet.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::WAFv2::RegexPatternSet`.
type CfnRegexPatternSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RegularExpressionList() *[]*string
	SetRegularExpressionList(val *[]*string)
	Scope() *string
	SetScope(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnRegexPatternSet
type jsiiProxy_CfnRegexPatternSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRegexPatternSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) RegularExpressionList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regularExpressionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::RegexPatternSet`.
func NewCfnRegexPatternSet(scope constructs.Construct, id *string, props *CfnRegexPatternSetProps) CfnRegexPatternSet {
	_init_.Initialize()

	j := jsiiProxy_CfnRegexPatternSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnRegexPatternSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::RegexPatternSet`.
func NewCfnRegexPatternSet_Override(c CfnRegexPatternSet, scope constructs.Construct, id *string, props *CfnRegexPatternSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnRegexPatternSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetRegularExpressionList(val *[]*string) {
	_jsii_.Set(
		j,
		"regularExpressionList",
		val,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
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
func CfnRegexPatternSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRegexPatternSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRegexPatternSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRegexPatternSet",
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
func CfnRegexPatternSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRegexPatternSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegexPatternSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafv2.CfnRegexPatternSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnRegexPatternSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRegexPatternSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRegexPatternSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRegexPatternSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnRegexPatternSet) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnRegexPatternSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnRegexPatternSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRegexPatternSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRegexPatternSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRegexPatternSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnRegexPatternSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRegexPatternSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRegexPatternSet) ToString() *string {
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
func (c *jsiiProxy_CfnRegexPatternSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::WAFv2::RegexPatternSet`.
type CfnRegexPatternSetProps struct {
	// `AWS::WAFv2::RegexPatternSet.RegularExpressionList`.
	RegularExpressionList *[]*string `json:"regularExpressionList"`
	// `AWS::WAFv2::RegexPatternSet.Scope`.
	Scope *string `json:"scope"`
	// `AWS::WAFv2::RegexPatternSet.Description`.
	Description *string `json:"description"`
	// `AWS::WAFv2::RegexPatternSet.Name`.
	Name *string `json:"name"`
	// `AWS::WAFv2::RegexPatternSet.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::WAFv2::RuleGroup`.
type CfnRuleGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	Capacity() *float64
	SetCapacity(val *float64)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Rules() interface{}
	SetRules(val interface{})
	Scope() *string
	SetScope(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VisibilityConfig() interface{}
	SetVisibilityConfig(val interface{})
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

// The jsii proxy struct for CfnRuleGroup
type jsiiProxy_CfnRuleGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRuleGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) VisibilityConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::RuleGroup`.
func NewCfnRuleGroup(scope constructs.Construct, id *string, props *CfnRuleGroupProps) CfnRuleGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnRuleGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::RuleGroup`.
func NewCfnRuleGroup_Override(c CfnRuleGroup, scope constructs.Construct, id *string, props *CfnRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetCapacity(val *float64) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetVisibilityConfig(val interface{}) {
	_jsii_.Set(
		j,
		"visibilityConfig",
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
func CfnRuleGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRuleGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
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
func CfnRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuleGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnRuleGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRuleGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRuleGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRuleGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnRuleGroup) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnRuleGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnRuleGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRuleGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRuleGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRuleGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnRuleGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRuleGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRuleGroup) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRuleGroup) ToString() *string {
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
func (c *jsiiProxy_CfnRuleGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnRuleGroup_AndStatementOneProperty struct {
	// `CfnRuleGroup.AndStatementOneProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnRuleGroup_AndStatementTwoProperty struct {
	// `CfnRuleGroup.AndStatementTwoProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnRuleGroup_ByteMatchStatementProperty struct {
	// `CfnRuleGroup.ByteMatchStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnRuleGroup.ByteMatchStatementProperty.PositionalConstraint`.
	PositionalConstraint *string `json:"positionalConstraint"`
	// `CfnRuleGroup.ByteMatchStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
	// `CfnRuleGroup.ByteMatchStatementProperty.SearchString`.
	SearchString *string `json:"searchString"`
	// `CfnRuleGroup.ByteMatchStatementProperty.SearchStringBase64`.
	SearchStringBase64 *string `json:"searchStringBase64"`
}

type CfnRuleGroup_FieldToMatchProperty struct {
	// `CfnRuleGroup.FieldToMatchProperty.AllQueryArguments`.
	AllQueryArguments interface{} `json:"allQueryArguments"`
	// `CfnRuleGroup.FieldToMatchProperty.Body`.
	Body interface{} `json:"body"`
	// `CfnRuleGroup.FieldToMatchProperty.Method`.
	Method interface{} `json:"method"`
	// `CfnRuleGroup.FieldToMatchProperty.QueryString`.
	QueryString interface{} `json:"queryString"`
	// `CfnRuleGroup.FieldToMatchProperty.SingleHeader`.
	SingleHeader interface{} `json:"singleHeader"`
	// `CfnRuleGroup.FieldToMatchProperty.SingleQueryArgument`.
	SingleQueryArgument interface{} `json:"singleQueryArgument"`
	// `CfnRuleGroup.FieldToMatchProperty.UriPath`.
	UriPath interface{} `json:"uriPath"`
}

type CfnRuleGroup_ForwardedIPConfigurationProperty struct {
	// `CfnRuleGroup.ForwardedIPConfigurationProperty.FallbackBehavior`.
	FallbackBehavior *string `json:"fallbackBehavior"`
	// `CfnRuleGroup.ForwardedIPConfigurationProperty.HeaderName`.
	HeaderName *string `json:"headerName"`
}

type CfnRuleGroup_GeoMatchStatementProperty struct {
	// `CfnRuleGroup.GeoMatchStatementProperty.CountryCodes`.
	CountryCodes *[]*string `json:"countryCodes"`
	// `CfnRuleGroup.GeoMatchStatementProperty.ForwardedIPConfig`.
	ForwardedIpConfig interface{} `json:"forwardedIpConfig"`
}

type CfnRuleGroup_IPSetForwardedIPConfigurationProperty interface {
	// `CfnRuleGroup.IPSetForwardedIPConfigurationProperty.FallbackBehavior`.
	FallbackBehavior() *string
	// `CfnRuleGroup.IPSetForwardedIPConfigurationProperty.HeaderName`.
	HeaderName() *string
	// `CfnRuleGroup.IPSetForwardedIPConfigurationProperty.Position`.
	Position() *string
}

// The jsii proxy for CfnRuleGroup_IPSetForwardedIPConfigurationProperty
type jsiiProxy_CfnRuleGroup_IPSetForwardedIPConfigurationProperty struct {
	_ byte // padding
}

func (j *jsiiProxy_CfnRuleGroup_IPSetForwardedIPConfigurationProperty) FallbackBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup_IPSetForwardedIPConfigurationProperty) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup_IPSetForwardedIPConfigurationProperty) Position() *string {
	var returns *string
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

type CfnRuleGroup_IPSetReferenceStatementProperty interface {
	// `CfnRuleGroup.IPSetReferenceStatementProperty.Arn`.
	Arn() *string
	// `CfnRuleGroup.IPSetReferenceStatementProperty.IPSetForwardedIPConfig`.
	IpSetForwardedIpConfig() interface{}
}

// The jsii proxy for CfnRuleGroup_IPSetReferenceStatementProperty
type jsiiProxy_CfnRuleGroup_IPSetReferenceStatementProperty struct {
	_ byte // padding
}

func (j *jsiiProxy_CfnRuleGroup_IPSetReferenceStatementProperty) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup_IPSetReferenceStatementProperty) IpSetForwardedIpConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetForwardedIpConfig",
		&returns,
	)
	return returns
}

type CfnRuleGroup_NotStatementOneProperty struct {
	// `CfnRuleGroup.NotStatementOneProperty.Statement`.
	Statement interface{} `json:"statement"`
}

type CfnRuleGroup_NotStatementTwoProperty struct {
	// `CfnRuleGroup.NotStatementTwoProperty.Statement`.
	Statement interface{} `json:"statement"`
}

type CfnRuleGroup_OrStatementOneProperty struct {
	// `CfnRuleGroup.OrStatementOneProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnRuleGroup_OrStatementTwoProperty struct {
	// `CfnRuleGroup.OrStatementTwoProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnRuleGroup_RateBasedStatementOneProperty struct {
	// `CfnRuleGroup.RateBasedStatementOneProperty.AggregateKeyType`.
	AggregateKeyType *string `json:"aggregateKeyType"`
	// `CfnRuleGroup.RateBasedStatementOneProperty.Limit`.
	Limit *float64 `json:"limit"`
	// `CfnRuleGroup.RateBasedStatementOneProperty.ForwardedIPConfig`.
	ForwardedIpConfig interface{} `json:"forwardedIpConfig"`
	// `CfnRuleGroup.RateBasedStatementOneProperty.ScopeDownStatement`.
	ScopeDownStatement interface{} `json:"scopeDownStatement"`
}

type CfnRuleGroup_RateBasedStatementTwoProperty struct {
	// `CfnRuleGroup.RateBasedStatementTwoProperty.AggregateKeyType`.
	AggregateKeyType *string `json:"aggregateKeyType"`
	// `CfnRuleGroup.RateBasedStatementTwoProperty.Limit`.
	Limit *float64 `json:"limit"`
	// `CfnRuleGroup.RateBasedStatementTwoProperty.ForwardedIPConfig`.
	ForwardedIpConfig interface{} `json:"forwardedIpConfig"`
	// `CfnRuleGroup.RateBasedStatementTwoProperty.ScopeDownStatement`.
	ScopeDownStatement interface{} `json:"scopeDownStatement"`
}

type CfnRuleGroup_RegexPatternSetReferenceStatementProperty struct {
	// `CfnRuleGroup.RegexPatternSetReferenceStatementProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnRuleGroup.RegexPatternSetReferenceStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnRuleGroup.RegexPatternSetReferenceStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

type CfnRuleGroup_RuleActionProperty struct {
	// `CfnRuleGroup.RuleActionProperty.Allow`.
	Allow interface{} `json:"allow"`
	// `CfnRuleGroup.RuleActionProperty.Block`.
	Block interface{} `json:"block"`
	// `CfnRuleGroup.RuleActionProperty.Count`.
	Count interface{} `json:"count"`
}

type CfnRuleGroup_RuleProperty struct {
	// `CfnRuleGroup.RuleProperty.Name`.
	Name *string `json:"name"`
	// `CfnRuleGroup.RuleProperty.Priority`.
	Priority *float64 `json:"priority"`
	// `CfnRuleGroup.RuleProperty.Statement`.
	Statement interface{} `json:"statement"`
	// `CfnRuleGroup.RuleProperty.VisibilityConfig`.
	VisibilityConfig interface{} `json:"visibilityConfig"`
	// `CfnRuleGroup.RuleProperty.Action`.
	Action interface{} `json:"action"`
}

type CfnRuleGroup_SizeConstraintStatementProperty struct {
	// `CfnRuleGroup.SizeConstraintStatementProperty.ComparisonOperator`.
	ComparisonOperator *string `json:"comparisonOperator"`
	// `CfnRuleGroup.SizeConstraintStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnRuleGroup.SizeConstraintStatementProperty.Size`.
	Size *float64 `json:"size"`
	// `CfnRuleGroup.SizeConstraintStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

type CfnRuleGroup_SqliMatchStatementProperty struct {
	// `CfnRuleGroup.SqliMatchStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnRuleGroup.SqliMatchStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

type CfnRuleGroup_StatementOneProperty struct {
	// `CfnRuleGroup.StatementOneProperty.AndStatement`.
	AndStatement interface{} `json:"andStatement"`
	// `CfnRuleGroup.StatementOneProperty.ByteMatchStatement`.
	ByteMatchStatement interface{} `json:"byteMatchStatement"`
	// `CfnRuleGroup.StatementOneProperty.GeoMatchStatement`.
	GeoMatchStatement interface{} `json:"geoMatchStatement"`
	// `CfnRuleGroup.StatementOneProperty.IPSetReferenceStatement`.
	IpSetReferenceStatement interface{} `json:"ipSetReferenceStatement"`
	// `CfnRuleGroup.StatementOneProperty.NotStatement`.
	NotStatement interface{} `json:"notStatement"`
	// `CfnRuleGroup.StatementOneProperty.OrStatement`.
	OrStatement interface{} `json:"orStatement"`
	// `CfnRuleGroup.StatementOneProperty.RateBasedStatement`.
	RateBasedStatement interface{} `json:"rateBasedStatement"`
	// `CfnRuleGroup.StatementOneProperty.RegexPatternSetReferenceStatement`.
	RegexPatternSetReferenceStatement interface{} `json:"regexPatternSetReferenceStatement"`
	// `CfnRuleGroup.StatementOneProperty.SizeConstraintStatement`.
	SizeConstraintStatement interface{} `json:"sizeConstraintStatement"`
	// `CfnRuleGroup.StatementOneProperty.SqliMatchStatement`.
	SqliMatchStatement interface{} `json:"sqliMatchStatement"`
	// `CfnRuleGroup.StatementOneProperty.XssMatchStatement`.
	XssMatchStatement interface{} `json:"xssMatchStatement"`
}

type CfnRuleGroup_StatementThreeProperty struct {
	// `CfnRuleGroup.StatementThreeProperty.ByteMatchStatement`.
	ByteMatchStatement interface{} `json:"byteMatchStatement"`
	// `CfnRuleGroup.StatementThreeProperty.GeoMatchStatement`.
	GeoMatchStatement interface{} `json:"geoMatchStatement"`
	// `CfnRuleGroup.StatementThreeProperty.IPSetReferenceStatement`.
	IpSetReferenceStatement interface{} `json:"ipSetReferenceStatement"`
	// `CfnRuleGroup.StatementThreeProperty.RegexPatternSetReferenceStatement`.
	RegexPatternSetReferenceStatement interface{} `json:"regexPatternSetReferenceStatement"`
	// `CfnRuleGroup.StatementThreeProperty.SizeConstraintStatement`.
	SizeConstraintStatement interface{} `json:"sizeConstraintStatement"`
	// `CfnRuleGroup.StatementThreeProperty.SqliMatchStatement`.
	SqliMatchStatement interface{} `json:"sqliMatchStatement"`
	// `CfnRuleGroup.StatementThreeProperty.XssMatchStatement`.
	XssMatchStatement interface{} `json:"xssMatchStatement"`
}

type CfnRuleGroup_StatementTwoProperty struct {
	// `CfnRuleGroup.StatementTwoProperty.AndStatement`.
	AndStatement interface{} `json:"andStatement"`
	// `CfnRuleGroup.StatementTwoProperty.ByteMatchStatement`.
	ByteMatchStatement interface{} `json:"byteMatchStatement"`
	// `CfnRuleGroup.StatementTwoProperty.GeoMatchStatement`.
	GeoMatchStatement interface{} `json:"geoMatchStatement"`
	// `CfnRuleGroup.StatementTwoProperty.IPSetReferenceStatement`.
	IpSetReferenceStatement interface{} `json:"ipSetReferenceStatement"`
	// `CfnRuleGroup.StatementTwoProperty.NotStatement`.
	NotStatement interface{} `json:"notStatement"`
	// `CfnRuleGroup.StatementTwoProperty.OrStatement`.
	OrStatement interface{} `json:"orStatement"`
	// `CfnRuleGroup.StatementTwoProperty.RateBasedStatement`.
	RateBasedStatement interface{} `json:"rateBasedStatement"`
	// `CfnRuleGroup.StatementTwoProperty.RegexPatternSetReferenceStatement`.
	RegexPatternSetReferenceStatement interface{} `json:"regexPatternSetReferenceStatement"`
	// `CfnRuleGroup.StatementTwoProperty.SizeConstraintStatement`.
	SizeConstraintStatement interface{} `json:"sizeConstraintStatement"`
	// `CfnRuleGroup.StatementTwoProperty.SqliMatchStatement`.
	SqliMatchStatement interface{} `json:"sqliMatchStatement"`
	// `CfnRuleGroup.StatementTwoProperty.XssMatchStatement`.
	XssMatchStatement interface{} `json:"xssMatchStatement"`
}

type CfnRuleGroup_TextTransformationProperty struct {
	// `CfnRuleGroup.TextTransformationProperty.Priority`.
	Priority *float64 `json:"priority"`
	// `CfnRuleGroup.TextTransformationProperty.Type`.
	Type *string `json:"type"`
}

type CfnRuleGroup_VisibilityConfigProperty struct {
	// `CfnRuleGroup.VisibilityConfigProperty.CloudWatchMetricsEnabled`.
	CloudWatchMetricsEnabled interface{} `json:"cloudWatchMetricsEnabled"`
	// `CfnRuleGroup.VisibilityConfigProperty.MetricName`.
	MetricName *string `json:"metricName"`
	// `CfnRuleGroup.VisibilityConfigProperty.SampledRequestsEnabled`.
	SampledRequestsEnabled interface{} `json:"sampledRequestsEnabled"`
}

type CfnRuleGroup_XssMatchStatementProperty struct {
	// `CfnRuleGroup.XssMatchStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnRuleGroup.XssMatchStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

// Properties for defining a `AWS::WAFv2::RuleGroup`.
type CfnRuleGroupProps struct {
	// `AWS::WAFv2::RuleGroup.Capacity`.
	Capacity *float64 `json:"capacity"`
	// `AWS::WAFv2::RuleGroup.Scope`.
	Scope *string `json:"scope"`
	// `AWS::WAFv2::RuleGroup.VisibilityConfig`.
	VisibilityConfig interface{} `json:"visibilityConfig"`
	// `AWS::WAFv2::RuleGroup.Description`.
	Description *string `json:"description"`
	// `AWS::WAFv2::RuleGroup.Name`.
	Name *string `json:"name"`
	// `AWS::WAFv2::RuleGroup.Rules`.
	Rules interface{} `json:"rules"`
	// `AWS::WAFv2::RuleGroup.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::WAFv2::WebACL`.
type CfnWebACL interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCapacity() *float64
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DefaultAction() interface{}
	SetDefaultAction(val interface{})
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Rules() interface{}
	SetRules(val interface{})
	Scope() *string
	SetScope(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VisibilityConfig() interface{}
	SetVisibilityConfig(val interface{})
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

// The jsii proxy struct for CfnWebACL
type jsiiProxy_CfnWebACL struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWebACL) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) AttrCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) DefaultAction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) VisibilityConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::WebACL`.
func NewCfnWebACL(scope constructs.Construct, id *string, props *CfnWebACLProps) CfnWebACL {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACL{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnWebACL",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::WebACL`.
func NewCfnWebACL_Override(c CfnWebACL, scope constructs.Construct, id *string, props *CfnWebACLProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnWebACL",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWebACL) SetDefaultAction(val interface{}) {
	_jsii_.Set(
		j,
		"defaultAction",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetVisibilityConfig(val interface{}) {
	_jsii_.Set(
		j,
		"visibilityConfig",
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
func CfnWebACL_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnWebACL",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWebACL_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnWebACL",
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
func CfnWebACL_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnWebACL",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACL_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafv2.CfnWebACL",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnWebACL) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnWebACL) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnWebACL) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnWebACL) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnWebACL) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnWebACL) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnWebACL) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnWebACL) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnWebACL) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnWebACL) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnWebACL) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWebACL) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnWebACL) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnWebACL) ToString() *string {
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
func (c *jsiiProxy_CfnWebACL) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnWebACL_AndStatementOneProperty struct {
	// `CfnWebACL.AndStatementOneProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnWebACL_AndStatementTwoProperty struct {
	// `CfnWebACL.AndStatementTwoProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnWebACL_ByteMatchStatementProperty struct {
	// `CfnWebACL.ByteMatchStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnWebACL.ByteMatchStatementProperty.PositionalConstraint`.
	PositionalConstraint *string `json:"positionalConstraint"`
	// `CfnWebACL.ByteMatchStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
	// `CfnWebACL.ByteMatchStatementProperty.SearchString`.
	SearchString *string `json:"searchString"`
	// `CfnWebACL.ByteMatchStatementProperty.SearchStringBase64`.
	SearchStringBase64 *string `json:"searchStringBase64"`
}

type CfnWebACL_DefaultActionProperty struct {
	// `CfnWebACL.DefaultActionProperty.Allow`.
	Allow interface{} `json:"allow"`
	// `CfnWebACL.DefaultActionProperty.Block`.
	Block interface{} `json:"block"`
}

type CfnWebACL_ExcludedRuleProperty struct {
	// `CfnWebACL.ExcludedRuleProperty.Name`.
	Name *string `json:"name"`
}

type CfnWebACL_FieldToMatchProperty struct {
	// `CfnWebACL.FieldToMatchProperty.AllQueryArguments`.
	AllQueryArguments interface{} `json:"allQueryArguments"`
	// `CfnWebACL.FieldToMatchProperty.Body`.
	Body interface{} `json:"body"`
	// `CfnWebACL.FieldToMatchProperty.Method`.
	Method interface{} `json:"method"`
	// `CfnWebACL.FieldToMatchProperty.QueryString`.
	QueryString interface{} `json:"queryString"`
	// `CfnWebACL.FieldToMatchProperty.SingleHeader`.
	SingleHeader interface{} `json:"singleHeader"`
	// `CfnWebACL.FieldToMatchProperty.SingleQueryArgument`.
	SingleQueryArgument interface{} `json:"singleQueryArgument"`
	// `CfnWebACL.FieldToMatchProperty.UriPath`.
	UriPath interface{} `json:"uriPath"`
}

type CfnWebACL_ForwardedIPConfigurationProperty struct {
	// `CfnWebACL.ForwardedIPConfigurationProperty.FallbackBehavior`.
	FallbackBehavior *string `json:"fallbackBehavior"`
	// `CfnWebACL.ForwardedIPConfigurationProperty.HeaderName`.
	HeaderName *string `json:"headerName"`
}

type CfnWebACL_GeoMatchStatementProperty struct {
	// `CfnWebACL.GeoMatchStatementProperty.CountryCodes`.
	CountryCodes *[]*string `json:"countryCodes"`
	// `CfnWebACL.GeoMatchStatementProperty.ForwardedIPConfig`.
	ForwardedIpConfig interface{} `json:"forwardedIpConfig"`
}

type CfnWebACL_IPSetForwardedIPConfigurationProperty interface {
	// `CfnWebACL.IPSetForwardedIPConfigurationProperty.FallbackBehavior`.
	FallbackBehavior() *string
	// `CfnWebACL.IPSetForwardedIPConfigurationProperty.HeaderName`.
	HeaderName() *string
	// `CfnWebACL.IPSetForwardedIPConfigurationProperty.Position`.
	Position() *string
}

// The jsii proxy for CfnWebACL_IPSetForwardedIPConfigurationProperty
type jsiiProxy_CfnWebACL_IPSetForwardedIPConfigurationProperty struct {
	_ byte // padding
}

func (j *jsiiProxy_CfnWebACL_IPSetForwardedIPConfigurationProperty) FallbackBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL_IPSetForwardedIPConfigurationProperty) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL_IPSetForwardedIPConfigurationProperty) Position() *string {
	var returns *string
	_jsii_.Get(
		j,
		"position",
		&returns,
	)
	return returns
}

type CfnWebACL_IPSetReferenceStatementProperty interface {
	// `CfnWebACL.IPSetReferenceStatementProperty.Arn`.
	Arn() *string
	// `CfnWebACL.IPSetReferenceStatementProperty.IPSetForwardedIPConfig`.
	IpSetForwardedIpConfig() interface{}
}

// The jsii proxy for CfnWebACL_IPSetReferenceStatementProperty
type jsiiProxy_CfnWebACL_IPSetReferenceStatementProperty struct {
	_ byte // padding
}

func (j *jsiiProxy_CfnWebACL_IPSetReferenceStatementProperty) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL_IPSetReferenceStatementProperty) IpSetForwardedIpConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetForwardedIpConfig",
		&returns,
	)
	return returns
}

type CfnWebACL_ManagedRuleGroupStatementProperty struct {
	// `CfnWebACL.ManagedRuleGroupStatementProperty.Name`.
	Name *string `json:"name"`
	// `CfnWebACL.ManagedRuleGroupStatementProperty.VendorName`.
	VendorName *string `json:"vendorName"`
	// `CfnWebACL.ManagedRuleGroupStatementProperty.ExcludedRules`.
	ExcludedRules interface{} `json:"excludedRules"`
}

type CfnWebACL_NotStatementOneProperty struct {
	// `CfnWebACL.NotStatementOneProperty.Statement`.
	Statement interface{} `json:"statement"`
}

type CfnWebACL_NotStatementTwoProperty struct {
	// `CfnWebACL.NotStatementTwoProperty.Statement`.
	Statement interface{} `json:"statement"`
}

type CfnWebACL_OrStatementOneProperty struct {
	// `CfnWebACL.OrStatementOneProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnWebACL_OrStatementTwoProperty struct {
	// `CfnWebACL.OrStatementTwoProperty.Statements`.
	Statements interface{} `json:"statements"`
}

type CfnWebACL_OverrideActionProperty struct {
	// `CfnWebACL.OverrideActionProperty.Count`.
	Count interface{} `json:"count"`
	// `CfnWebACL.OverrideActionProperty.None`.
	None interface{} `json:"none"`
}

type CfnWebACL_RateBasedStatementOneProperty struct {
	// `CfnWebACL.RateBasedStatementOneProperty.AggregateKeyType`.
	AggregateKeyType *string `json:"aggregateKeyType"`
	// `CfnWebACL.RateBasedStatementOneProperty.Limit`.
	Limit *float64 `json:"limit"`
	// `CfnWebACL.RateBasedStatementOneProperty.ForwardedIPConfig`.
	ForwardedIpConfig interface{} `json:"forwardedIpConfig"`
	// `CfnWebACL.RateBasedStatementOneProperty.ScopeDownStatement`.
	ScopeDownStatement interface{} `json:"scopeDownStatement"`
}

type CfnWebACL_RateBasedStatementTwoProperty struct {
	// `CfnWebACL.RateBasedStatementTwoProperty.AggregateKeyType`.
	AggregateKeyType *string `json:"aggregateKeyType"`
	// `CfnWebACL.RateBasedStatementTwoProperty.Limit`.
	Limit *float64 `json:"limit"`
	// `CfnWebACL.RateBasedStatementTwoProperty.ForwardedIPConfig`.
	ForwardedIpConfig interface{} `json:"forwardedIpConfig"`
	// `CfnWebACL.RateBasedStatementTwoProperty.ScopeDownStatement`.
	ScopeDownStatement interface{} `json:"scopeDownStatement"`
}

type CfnWebACL_RegexPatternSetReferenceStatementProperty struct {
	// `CfnWebACL.RegexPatternSetReferenceStatementProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnWebACL.RegexPatternSetReferenceStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnWebACL.RegexPatternSetReferenceStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

type CfnWebACL_RuleActionProperty struct {
	// `CfnWebACL.RuleActionProperty.Allow`.
	Allow interface{} `json:"allow"`
	// `CfnWebACL.RuleActionProperty.Block`.
	Block interface{} `json:"block"`
	// `CfnWebACL.RuleActionProperty.Count`.
	Count interface{} `json:"count"`
}

type CfnWebACL_RuleGroupReferenceStatementProperty struct {
	// `CfnWebACL.RuleGroupReferenceStatementProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnWebACL.RuleGroupReferenceStatementProperty.ExcludedRules`.
	ExcludedRules interface{} `json:"excludedRules"`
}

type CfnWebACL_RuleProperty struct {
	// `CfnWebACL.RuleProperty.Name`.
	Name *string `json:"name"`
	// `CfnWebACL.RuleProperty.Priority`.
	Priority *float64 `json:"priority"`
	// `CfnWebACL.RuleProperty.Statement`.
	Statement interface{} `json:"statement"`
	// `CfnWebACL.RuleProperty.VisibilityConfig`.
	VisibilityConfig interface{} `json:"visibilityConfig"`
	// `CfnWebACL.RuleProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnWebACL.RuleProperty.OverrideAction`.
	OverrideAction interface{} `json:"overrideAction"`
}

type CfnWebACL_SizeConstraintStatementProperty struct {
	// `CfnWebACL.SizeConstraintStatementProperty.ComparisonOperator`.
	ComparisonOperator *string `json:"comparisonOperator"`
	// `CfnWebACL.SizeConstraintStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnWebACL.SizeConstraintStatementProperty.Size`.
	Size *float64 `json:"size"`
	// `CfnWebACL.SizeConstraintStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

type CfnWebACL_SqliMatchStatementProperty struct {
	// `CfnWebACL.SqliMatchStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnWebACL.SqliMatchStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

type CfnWebACL_StatementOneProperty struct {
	// `CfnWebACL.StatementOneProperty.AndStatement`.
	AndStatement interface{} `json:"andStatement"`
	// `CfnWebACL.StatementOneProperty.ByteMatchStatement`.
	ByteMatchStatement interface{} `json:"byteMatchStatement"`
	// `CfnWebACL.StatementOneProperty.GeoMatchStatement`.
	GeoMatchStatement interface{} `json:"geoMatchStatement"`
	// `CfnWebACL.StatementOneProperty.IPSetReferenceStatement`.
	IpSetReferenceStatement interface{} `json:"ipSetReferenceStatement"`
	// `CfnWebACL.StatementOneProperty.ManagedRuleGroupStatement`.
	ManagedRuleGroupStatement interface{} `json:"managedRuleGroupStatement"`
	// `CfnWebACL.StatementOneProperty.NotStatement`.
	NotStatement interface{} `json:"notStatement"`
	// `CfnWebACL.StatementOneProperty.OrStatement`.
	OrStatement interface{} `json:"orStatement"`
	// `CfnWebACL.StatementOneProperty.RateBasedStatement`.
	RateBasedStatement interface{} `json:"rateBasedStatement"`
	// `CfnWebACL.StatementOneProperty.RegexPatternSetReferenceStatement`.
	RegexPatternSetReferenceStatement interface{} `json:"regexPatternSetReferenceStatement"`
	// `CfnWebACL.StatementOneProperty.RuleGroupReferenceStatement`.
	RuleGroupReferenceStatement interface{} `json:"ruleGroupReferenceStatement"`
	// `CfnWebACL.StatementOneProperty.SizeConstraintStatement`.
	SizeConstraintStatement interface{} `json:"sizeConstraintStatement"`
	// `CfnWebACL.StatementOneProperty.SqliMatchStatement`.
	SqliMatchStatement interface{} `json:"sqliMatchStatement"`
	// `CfnWebACL.StatementOneProperty.XssMatchStatement`.
	XssMatchStatement interface{} `json:"xssMatchStatement"`
}

type CfnWebACL_StatementThreeProperty struct {
	// `CfnWebACL.StatementThreeProperty.ByteMatchStatement`.
	ByteMatchStatement interface{} `json:"byteMatchStatement"`
	// `CfnWebACL.StatementThreeProperty.GeoMatchStatement`.
	GeoMatchStatement interface{} `json:"geoMatchStatement"`
	// `CfnWebACL.StatementThreeProperty.IPSetReferenceStatement`.
	IpSetReferenceStatement interface{} `json:"ipSetReferenceStatement"`
	// `CfnWebACL.StatementThreeProperty.ManagedRuleGroupStatement`.
	ManagedRuleGroupStatement interface{} `json:"managedRuleGroupStatement"`
	// `CfnWebACL.StatementThreeProperty.RegexPatternSetReferenceStatement`.
	RegexPatternSetReferenceStatement interface{} `json:"regexPatternSetReferenceStatement"`
	// `CfnWebACL.StatementThreeProperty.RuleGroupReferenceStatement`.
	RuleGroupReferenceStatement interface{} `json:"ruleGroupReferenceStatement"`
	// `CfnWebACL.StatementThreeProperty.SizeConstraintStatement`.
	SizeConstraintStatement interface{} `json:"sizeConstraintStatement"`
	// `CfnWebACL.StatementThreeProperty.SqliMatchStatement`.
	SqliMatchStatement interface{} `json:"sqliMatchStatement"`
	// `CfnWebACL.StatementThreeProperty.XssMatchStatement`.
	XssMatchStatement interface{} `json:"xssMatchStatement"`
}

type CfnWebACL_StatementTwoProperty struct {
	// `CfnWebACL.StatementTwoProperty.AndStatement`.
	AndStatement interface{} `json:"andStatement"`
	// `CfnWebACL.StatementTwoProperty.ByteMatchStatement`.
	ByteMatchStatement interface{} `json:"byteMatchStatement"`
	// `CfnWebACL.StatementTwoProperty.GeoMatchStatement`.
	GeoMatchStatement interface{} `json:"geoMatchStatement"`
	// `CfnWebACL.StatementTwoProperty.IPSetReferenceStatement`.
	IpSetReferenceStatement interface{} `json:"ipSetReferenceStatement"`
	// `CfnWebACL.StatementTwoProperty.ManagedRuleGroupStatement`.
	ManagedRuleGroupStatement interface{} `json:"managedRuleGroupStatement"`
	// `CfnWebACL.StatementTwoProperty.NotStatement`.
	NotStatement interface{} `json:"notStatement"`
	// `CfnWebACL.StatementTwoProperty.OrStatement`.
	OrStatement interface{} `json:"orStatement"`
	// `CfnWebACL.StatementTwoProperty.RateBasedStatement`.
	RateBasedStatement interface{} `json:"rateBasedStatement"`
	// `CfnWebACL.StatementTwoProperty.RegexPatternSetReferenceStatement`.
	RegexPatternSetReferenceStatement interface{} `json:"regexPatternSetReferenceStatement"`
	// `CfnWebACL.StatementTwoProperty.RuleGroupReferenceStatement`.
	RuleGroupReferenceStatement interface{} `json:"ruleGroupReferenceStatement"`
	// `CfnWebACL.StatementTwoProperty.SizeConstraintStatement`.
	SizeConstraintStatement interface{} `json:"sizeConstraintStatement"`
	// `CfnWebACL.StatementTwoProperty.SqliMatchStatement`.
	SqliMatchStatement interface{} `json:"sqliMatchStatement"`
	// `CfnWebACL.StatementTwoProperty.XssMatchStatement`.
	XssMatchStatement interface{} `json:"xssMatchStatement"`
}

type CfnWebACL_TextTransformationProperty struct {
	// `CfnWebACL.TextTransformationProperty.Priority`.
	Priority *float64 `json:"priority"`
	// `CfnWebACL.TextTransformationProperty.Type`.
	Type *string `json:"type"`
}

type CfnWebACL_VisibilityConfigProperty struct {
	// `CfnWebACL.VisibilityConfigProperty.CloudWatchMetricsEnabled`.
	CloudWatchMetricsEnabled interface{} `json:"cloudWatchMetricsEnabled"`
	// `CfnWebACL.VisibilityConfigProperty.MetricName`.
	MetricName *string `json:"metricName"`
	// `CfnWebACL.VisibilityConfigProperty.SampledRequestsEnabled`.
	SampledRequestsEnabled interface{} `json:"sampledRequestsEnabled"`
}

type CfnWebACL_XssMatchStatementProperty struct {
	// `CfnWebACL.XssMatchStatementProperty.FieldToMatch`.
	FieldToMatch interface{} `json:"fieldToMatch"`
	// `CfnWebACL.XssMatchStatementProperty.TextTransformations`.
	TextTransformations interface{} `json:"textTransformations"`
}

// A CloudFormation `AWS::WAFv2::WebACLAssociation`.
type CfnWebACLAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ResourceArn() *string
	SetResourceArn(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	WebAclArn() *string
	SetWebAclArn(val *string)
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

// The jsii proxy struct for CfnWebACLAssociation
type jsiiProxy_CfnWebACLAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) WebAclArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclArn",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::WebACLAssociation`.
func NewCfnWebACLAssociation(scope constructs.Construct, id *string, props *CfnWebACLAssociationProps) CfnWebACLAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACLAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnWebACLAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::WebACLAssociation`.
func NewCfnWebACLAssociation_Override(c CfnWebACLAssociation, scope constructs.Construct, id *string, props *CfnWebACLAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnWebACLAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWebACLAssociation) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnWebACLAssociation) SetWebAclArn(val *string) {
	_jsii_.Set(
		j,
		"webAclArn",
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
func CfnWebACLAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnWebACLAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWebACLAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnWebACLAssociation",
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
func CfnWebACLAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnWebACLAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACLAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafv2.CfnWebACLAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnWebACLAssociation) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnWebACLAssociation) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnWebACLAssociation) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnWebACLAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnWebACLAssociation) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnWebACLAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnWebACLAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnWebACLAssociation) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnWebACLAssociation) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnWebACLAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnWebACLAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnWebACLAssociation) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnWebACLAssociation) ToString() *string {
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
func (c *jsiiProxy_CfnWebACLAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::WAFv2::WebACLAssociation`.
type CfnWebACLAssociationProps struct {
	// `AWS::WAFv2::WebACLAssociation.ResourceArn`.
	ResourceArn *string `json:"resourceArn"`
	// `AWS::WAFv2::WebACLAssociation.WebACLArn`.
	WebAclArn *string `json:"webAclArn"`
}

// Properties for defining a `AWS::WAFv2::WebACL`.
type CfnWebACLProps struct {
	// `AWS::WAFv2::WebACL.DefaultAction`.
	DefaultAction interface{} `json:"defaultAction"`
	// `AWS::WAFv2::WebACL.Scope`.
	Scope *string `json:"scope"`
	// `AWS::WAFv2::WebACL.VisibilityConfig`.
	VisibilityConfig interface{} `json:"visibilityConfig"`
	// `AWS::WAFv2::WebACL.Description`.
	Description *string `json:"description"`
	// `AWS::WAFv2::WebACL.Name`.
	Name *string `json:"name"`
	// `AWS::WAFv2::WebACL.Rules`.
	Rules interface{} `json:"rules"`
	// `AWS::WAFv2::WebACL.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}


package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::QuickSight::Analysis`.
type CfnAnalysis interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AnalysisId() *string
	SetAnalysisId(val *string)
	AttrArn() *string
	AttrCreatedTime() *string
	AttrDataSetArns() *[]*string
	AttrLastUpdatedTime() *string
	AttrSheets() awscdk.IResolvable
	AttrStatus() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Errors() interface{}
	SetErrors(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Parameters() interface{}
	SetParameters(val interface{})
	Permissions() interface{}
	SetPermissions(val interface{})
	Ref() *string
	SourceEntity() interface{}
	SetSourceEntity(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	ThemeArn() *string
	SetThemeArn(val *string)
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

// The jsii proxy struct for CfnAnalysis
type jsiiProxy_CfnAnalysis struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAnalysis) AnalysisId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrDataSetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDataSetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrSheets() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrSheets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Errors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) SourceEntity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) ThemeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Analysis`.
func NewCfnAnalysis(scope constructs.Construct, id *string, props *CfnAnalysisProps) CfnAnalysis {
	_init_.Initialize()

	j := jsiiProxy_CfnAnalysis{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Analysis`.
func NewCfnAnalysis_Override(c CfnAnalysis, scope constructs.Construct, id *string, props *CfnAnalysisProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetAnalysisId(val *string) {
	_jsii_.Set(
		j,
		"analysisId",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetErrors(val interface{}) {
	_jsii_.Set(
		j,
		"errors",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetSourceEntity(val interface{}) {
	_jsii_.Set(
		j,
		"sourceEntity",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetThemeArn(val *string) {
	_jsii_.Set(
		j,
		"themeArn",
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
func CfnAnalysis_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAnalysis_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
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
func CfnAnalysis_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnalysis_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnAnalysis) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAnalysis) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAnalysis) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAnalysis) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnAnalysis) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnAnalysis) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnAnalysis) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAnalysis) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAnalysis) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAnalysis) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAnalysis) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAnalysis) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAnalysis) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAnalysis) ToString() *string {
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
func (c *jsiiProxy_CfnAnalysis) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnAnalysis_AnalysisErrorProperty struct {
	// `CfnAnalysis.AnalysisErrorProperty.Message`.
	Message *string `json:"message"`
	// `CfnAnalysis.AnalysisErrorProperty.Type`.
	Type *string `json:"type"`
}

type CfnAnalysis_AnalysisSourceEntityProperty struct {
	// `CfnAnalysis.AnalysisSourceEntityProperty.SourceTemplate`.
	SourceTemplate interface{} `json:"sourceTemplate"`
}

type CfnAnalysis_AnalysisSourceTemplateProperty struct {
	// `CfnAnalysis.AnalysisSourceTemplateProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnAnalysis.AnalysisSourceTemplateProperty.DataSetReferences`.
	DataSetReferences interface{} `json:"dataSetReferences"`
}

type CfnAnalysis_DataSetReferenceProperty struct {
	// `CfnAnalysis.DataSetReferenceProperty.DataSetArn`.
	DataSetArn *string `json:"dataSetArn"`
	// `CfnAnalysis.DataSetReferenceProperty.DataSetPlaceholder`.
	DataSetPlaceholder *string `json:"dataSetPlaceholder"`
}

type CfnAnalysis_DateTimeParameterProperty struct {
	// `CfnAnalysis.DateTimeParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnAnalysis.DateTimeParameterProperty.Values`.
	Values *[]*string `json:"values"`
}

type CfnAnalysis_DecimalParameterProperty struct {
	// `CfnAnalysis.DecimalParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnAnalysis.DecimalParameterProperty.Values`.
	Values interface{} `json:"values"`
}

type CfnAnalysis_IntegerParameterProperty struct {
	// `CfnAnalysis.IntegerParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnAnalysis.IntegerParameterProperty.Values`.
	Values interface{} `json:"values"`
}

type CfnAnalysis_ParametersProperty struct {
	// `CfnAnalysis.ParametersProperty.DateTimeParameters`.
	DateTimeParameters interface{} `json:"dateTimeParameters"`
	// `CfnAnalysis.ParametersProperty.DecimalParameters`.
	DecimalParameters interface{} `json:"decimalParameters"`
	// `CfnAnalysis.ParametersProperty.IntegerParameters`.
	IntegerParameters interface{} `json:"integerParameters"`
	// `CfnAnalysis.ParametersProperty.StringParameters`.
	StringParameters interface{} `json:"stringParameters"`
}

type CfnAnalysis_ResourcePermissionProperty struct {
	// `CfnAnalysis.ResourcePermissionProperty.Actions`.
	Actions *[]*string `json:"actions"`
	// `CfnAnalysis.ResourcePermissionProperty.Principal`.
	Principal *string `json:"principal"`
}

type CfnAnalysis_SheetProperty struct {
	// `CfnAnalysis.SheetProperty.Name`.
	Name *string `json:"name"`
	// `CfnAnalysis.SheetProperty.SheetId`.
	SheetId *string `json:"sheetId"`
}

type CfnAnalysis_StringParameterProperty struct {
	// `CfnAnalysis.StringParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnAnalysis.StringParameterProperty.Values`.
	Values *[]*string `json:"values"`
}

// Properties for defining a `AWS::QuickSight::Analysis`.
type CfnAnalysisProps struct {
	// `AWS::QuickSight::Analysis.AnalysisId`.
	AnalysisId *string `json:"analysisId"`
	// `AWS::QuickSight::Analysis.AwsAccountId`.
	AwsAccountId *string `json:"awsAccountId"`
	// `AWS::QuickSight::Analysis.Errors`.
	Errors interface{} `json:"errors"`
	// `AWS::QuickSight::Analysis.Name`.
	Name *string `json:"name"`
	// `AWS::QuickSight::Analysis.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `AWS::QuickSight::Analysis.Permissions`.
	Permissions interface{} `json:"permissions"`
	// `AWS::QuickSight::Analysis.SourceEntity`.
	SourceEntity interface{} `json:"sourceEntity"`
	// `AWS::QuickSight::Analysis.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::QuickSight::Analysis.ThemeArn`.
	ThemeArn *string `json:"themeArn"`
}

// A CloudFormation `AWS::QuickSight::Dashboard`.
type CfnDashboard interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreatedTime() *string
	AttrLastPublishedTime() *string
	AttrLastUpdatedTime() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DashboardId() *string
	SetDashboardId(val *string)
	DashboardPublishOptions() interface{}
	SetDashboardPublishOptions(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Parameters() interface{}
	SetParameters(val interface{})
	Permissions() interface{}
	SetPermissions(val interface{})
	Ref() *string
	SourceEntity() interface{}
	SetSourceEntity(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	ThemeArn() *string
	SetThemeArn(val *string)
	UpdatedProperites() *map[string]interface{}
	VersionDescription() *string
	SetVersionDescription(val *string)
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

// The jsii proxy struct for CfnDashboard
type jsiiProxy_CfnDashboard struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDashboard) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AttrLastPublishedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastPublishedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) DashboardPublishOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashboardPublishOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) SourceEntity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) ThemeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Dashboard`.
func NewCfnDashboard(scope constructs.Construct, id *string, props *CfnDashboardProps) CfnDashboard {
	_init_.Initialize()

	j := jsiiProxy_CfnDashboard{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Dashboard`.
func NewCfnDashboard_Override(c CfnDashboard, scope constructs.Construct, id *string, props *CfnDashboardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDashboard) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetDashboardId(val *string) {
	_jsii_.Set(
		j,
		"dashboardId",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetDashboardPublishOptions(val interface{}) {
	_jsii_.Set(
		j,
		"dashboardPublishOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetSourceEntity(val interface{}) {
	_jsii_.Set(
		j,
		"sourceEntity",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetThemeArn(val *string) {
	_jsii_.Set(
		j,
		"themeArn",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
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
func CfnDashboard_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDashboard_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
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
func CfnDashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDashboard_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDashboard) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDashboard) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDashboard) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDashboard) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDashboard) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDashboard) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDashboard) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDashboard) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDashboard) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDashboard) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDashboard) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDashboard) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDashboard) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDashboard) ToString() *string {
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
func (c *jsiiProxy_CfnDashboard) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDashboard_AdHocFilteringOptionProperty struct {
	// `CfnDashboard.AdHocFilteringOptionProperty.AvailabilityStatus`.
	AvailabilityStatus *string `json:"availabilityStatus"`
}

type CfnDashboard_DashboardPublishOptionsProperty struct {
	// `CfnDashboard.DashboardPublishOptionsProperty.AdHocFilteringOption`.
	AdHocFilteringOption interface{} `json:"adHocFilteringOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.ExportToCSVOption`.
	ExportToCsvOption interface{} `json:"exportToCsvOption"`
	// `CfnDashboard.DashboardPublishOptionsProperty.SheetControlsOption`.
	SheetControlsOption interface{} `json:"sheetControlsOption"`
}

type CfnDashboard_DashboardSourceEntityProperty struct {
	// `CfnDashboard.DashboardSourceEntityProperty.SourceTemplate`.
	SourceTemplate interface{} `json:"sourceTemplate"`
}

type CfnDashboard_DashboardSourceTemplateProperty struct {
	// `CfnDashboard.DashboardSourceTemplateProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnDashboard.DashboardSourceTemplateProperty.DataSetReferences`.
	DataSetReferences interface{} `json:"dataSetReferences"`
}

type CfnDashboard_DataSetReferenceProperty struct {
	// `CfnDashboard.DataSetReferenceProperty.DataSetArn`.
	DataSetArn *string `json:"dataSetArn"`
	// `CfnDashboard.DataSetReferenceProperty.DataSetPlaceholder`.
	DataSetPlaceholder *string `json:"dataSetPlaceholder"`
}

type CfnDashboard_DateTimeParameterProperty struct {
	// `CfnDashboard.DateTimeParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnDashboard.DateTimeParameterProperty.Values`.
	Values *[]*string `json:"values"`
}

type CfnDashboard_DecimalParameterProperty struct {
	// `CfnDashboard.DecimalParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnDashboard.DecimalParameterProperty.Values`.
	Values interface{} `json:"values"`
}

type CfnDashboard_ExportToCSVOptionProperty struct {
	// `CfnDashboard.ExportToCSVOptionProperty.AvailabilityStatus`.
	AvailabilityStatus *string `json:"availabilityStatus"`
}

type CfnDashboard_IntegerParameterProperty struct {
	// `CfnDashboard.IntegerParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnDashboard.IntegerParameterProperty.Values`.
	Values interface{} `json:"values"`
}

type CfnDashboard_ParametersProperty struct {
	// `CfnDashboard.ParametersProperty.DateTimeParameters`.
	DateTimeParameters interface{} `json:"dateTimeParameters"`
	// `CfnDashboard.ParametersProperty.DecimalParameters`.
	DecimalParameters interface{} `json:"decimalParameters"`
	// `CfnDashboard.ParametersProperty.IntegerParameters`.
	IntegerParameters interface{} `json:"integerParameters"`
	// `CfnDashboard.ParametersProperty.StringParameters`.
	StringParameters interface{} `json:"stringParameters"`
}

type CfnDashboard_ResourcePermissionProperty struct {
	// `CfnDashboard.ResourcePermissionProperty.Actions`.
	Actions *[]*string `json:"actions"`
	// `CfnDashboard.ResourcePermissionProperty.Principal`.
	Principal *string `json:"principal"`
}

type CfnDashboard_SheetControlsOptionProperty struct {
	// `CfnDashboard.SheetControlsOptionProperty.VisibilityState`.
	VisibilityState *string `json:"visibilityState"`
}

type CfnDashboard_StringParameterProperty struct {
	// `CfnDashboard.StringParameterProperty.Name`.
	Name *string `json:"name"`
	// `CfnDashboard.StringParameterProperty.Values`.
	Values *[]*string `json:"values"`
}

// Properties for defining a `AWS::QuickSight::Dashboard`.
type CfnDashboardProps struct {
	// `AWS::QuickSight::Dashboard.AwsAccountId`.
	AwsAccountId *string `json:"awsAccountId"`
	// `AWS::QuickSight::Dashboard.DashboardId`.
	DashboardId *string `json:"dashboardId"`
	// `AWS::QuickSight::Dashboard.DashboardPublishOptions`.
	DashboardPublishOptions interface{} `json:"dashboardPublishOptions"`
	// `AWS::QuickSight::Dashboard.Name`.
	Name *string `json:"name"`
	// `AWS::QuickSight::Dashboard.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `AWS::QuickSight::Dashboard.Permissions`.
	Permissions interface{} `json:"permissions"`
	// `AWS::QuickSight::Dashboard.SourceEntity`.
	SourceEntity interface{} `json:"sourceEntity"`
	// `AWS::QuickSight::Dashboard.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::QuickSight::Dashboard.ThemeArn`.
	ThemeArn *string `json:"themeArn"`
	// `AWS::QuickSight::Dashboard.VersionDescription`.
	VersionDescription *string `json:"versionDescription"`
}

// A CloudFormation `AWS::QuickSight::Template`.
type CfnTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreatedTime() *string
	AttrLastUpdatedTime() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Permissions() interface{}
	SetPermissions(val interface{})
	Ref() *string
	SourceEntity() interface{}
	SetSourceEntity(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TemplateId() *string
	SetTemplateId(val *string)
	UpdatedProperites() *map[string]interface{}
	VersionDescription() *string
	SetVersionDescription(val *string)
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

// The jsii proxy struct for CfnTemplate
type jsiiProxy_CfnTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) SourceEntity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) TemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Template`.
func NewCfnTemplate(scope constructs.Construct, id *string, props *CfnTemplateProps) CfnTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Template`.
func NewCfnTemplate_Override(c CfnTemplate, scope constructs.Construct, id *string, props *CfnTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTemplate) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetSourceEntity(val interface{}) {
	_jsii_.Set(
		j,
		"sourceEntity",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetTemplateId(val *string) {
	_jsii_.Set(
		j,
		"templateId",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
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
func CfnTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
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
func CfnTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnTemplate) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTemplate) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTemplate) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnTemplate) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTemplate) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTemplate) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTemplate) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTemplate) ToString() *string {
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
func (c *jsiiProxy_CfnTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnTemplate_DataSetReferenceProperty struct {
	// `CfnTemplate.DataSetReferenceProperty.DataSetArn`.
	DataSetArn *string `json:"dataSetArn"`
	// `CfnTemplate.DataSetReferenceProperty.DataSetPlaceholder`.
	DataSetPlaceholder *string `json:"dataSetPlaceholder"`
}

type CfnTemplate_ResourcePermissionProperty struct {
	// `CfnTemplate.ResourcePermissionProperty.Actions`.
	Actions *[]*string `json:"actions"`
	// `CfnTemplate.ResourcePermissionProperty.Principal`.
	Principal *string `json:"principal"`
}

type CfnTemplate_TemplateSourceAnalysisProperty struct {
	// `CfnTemplate.TemplateSourceAnalysisProperty.Arn`.
	Arn *string `json:"arn"`
	// `CfnTemplate.TemplateSourceAnalysisProperty.DataSetReferences`.
	DataSetReferences interface{} `json:"dataSetReferences"`
}

type CfnTemplate_TemplateSourceEntityProperty struct {
	// `CfnTemplate.TemplateSourceEntityProperty.SourceAnalysis`.
	SourceAnalysis interface{} `json:"sourceAnalysis"`
	// `CfnTemplate.TemplateSourceEntityProperty.SourceTemplate`.
	SourceTemplate interface{} `json:"sourceTemplate"`
}

type CfnTemplate_TemplateSourceTemplateProperty struct {
	// `CfnTemplate.TemplateSourceTemplateProperty.Arn`.
	Arn *string `json:"arn"`
}

// Properties for defining a `AWS::QuickSight::Template`.
type CfnTemplateProps struct {
	// `AWS::QuickSight::Template.AwsAccountId`.
	AwsAccountId *string `json:"awsAccountId"`
	// `AWS::QuickSight::Template.TemplateId`.
	TemplateId *string `json:"templateId"`
	// `AWS::QuickSight::Template.Name`.
	Name *string `json:"name"`
	// `AWS::QuickSight::Template.Permissions`.
	Permissions interface{} `json:"permissions"`
	// `AWS::QuickSight::Template.SourceEntity`.
	SourceEntity interface{} `json:"sourceEntity"`
	// `AWS::QuickSight::Template.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::QuickSight::Template.VersionDescription`.
	VersionDescription *string `json:"versionDescription"`
}

// A CloudFormation `AWS::QuickSight::Theme`.
type CfnTheme interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrCreatedTime() *string
	AttrLastUpdatedTime() *string
	AttrType() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	BaseThemeId() *string
	SetBaseThemeId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Configuration() interface{}
	SetConfiguration(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Permissions() interface{}
	SetPermissions(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	ThemeId() *string
	SetThemeId(val *string)
	UpdatedProperites() *map[string]interface{}
	VersionDescription() *string
	SetVersionDescription(val *string)
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

// The jsii proxy struct for CfnTheme
type jsiiProxy_CfnTheme struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTheme) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) BaseThemeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseThemeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) ThemeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Theme`.
func NewCfnTheme(scope constructs.Construct, id *string, props *CfnThemeProps) CfnTheme {
	_init_.Initialize()

	j := jsiiProxy_CfnTheme{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Theme`.
func NewCfnTheme_Override(c CfnTheme, scope constructs.Construct, id *string, props *CfnThemeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTheme) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetBaseThemeId(val *string) {
	_jsii_.Set(
		j,
		"baseThemeId",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetThemeId(val *string) {
	_jsii_.Set(
		j,
		"themeId",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
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
func CfnTheme_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTheme_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
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
func CfnTheme_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTheme_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnTheme) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTheme) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTheme) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTheme) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnTheme) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnTheme) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnTheme) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTheme) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTheme) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTheme) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnTheme) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTheme) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTheme) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTheme) ToString() *string {
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
func (c *jsiiProxy_CfnTheme) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnTheme_BorderStyleProperty struct {
	// `CfnTheme.BorderStyleProperty.Show`.
	Show interface{} `json:"show"`
}

type CfnTheme_DataColorPaletteProperty struct {
	// `CfnTheme.DataColorPaletteProperty.Colors`.
	Colors *[]*string `json:"colors"`
	// `CfnTheme.DataColorPaletteProperty.EmptyFillColor`.
	EmptyFillColor *string `json:"emptyFillColor"`
	// `CfnTheme.DataColorPaletteProperty.MinMaxGradient`.
	MinMaxGradient *[]*string `json:"minMaxGradient"`
}

type CfnTheme_FontProperty struct {
	// `CfnTheme.FontProperty.FontFamily`.
	FontFamily *string `json:"fontFamily"`
}

type CfnTheme_GutterStyleProperty struct {
	// `CfnTheme.GutterStyleProperty.Show`.
	Show interface{} `json:"show"`
}

type CfnTheme_MarginStyleProperty struct {
	// `CfnTheme.MarginStyleProperty.Show`.
	Show interface{} `json:"show"`
}

type CfnTheme_ResourcePermissionProperty struct {
	// `CfnTheme.ResourcePermissionProperty.Actions`.
	Actions *[]*string `json:"actions"`
	// `CfnTheme.ResourcePermissionProperty.Principal`.
	Principal *string `json:"principal"`
}

type CfnTheme_SheetStyleProperty struct {
	// `CfnTheme.SheetStyleProperty.Tile`.
	Tile interface{} `json:"tile"`
	// `CfnTheme.SheetStyleProperty.TileLayout`.
	TileLayout interface{} `json:"tileLayout"`
}

type CfnTheme_ThemeConfigurationProperty struct {
	// `CfnTheme.ThemeConfigurationProperty.DataColorPalette`.
	DataColorPalette interface{} `json:"dataColorPalette"`
	// `CfnTheme.ThemeConfigurationProperty.Sheet`.
	Sheet interface{} `json:"sheet"`
	// `CfnTheme.ThemeConfigurationProperty.Typography`.
	Typography interface{} `json:"typography"`
	// `CfnTheme.ThemeConfigurationProperty.UIColorPalette`.
	UiColorPalette interface{} `json:"uiColorPalette"`
}

type CfnTheme_TileLayoutStyleProperty struct {
	// `CfnTheme.TileLayoutStyleProperty.Gutter`.
	Gutter interface{} `json:"gutter"`
	// `CfnTheme.TileLayoutStyleProperty.Margin`.
	Margin interface{} `json:"margin"`
}

type CfnTheme_TileStyleProperty struct {
	// `CfnTheme.TileStyleProperty.Border`.
	Border interface{} `json:"border"`
}

type CfnTheme_TypographyProperty struct {
	// `CfnTheme.TypographyProperty.FontFamilies`.
	FontFamilies interface{} `json:"fontFamilies"`
}

type CfnTheme_UIColorPaletteProperty struct {
	// `CfnTheme.UIColorPaletteProperty.Accent`.
	Accent *string `json:"accent"`
	// `CfnTheme.UIColorPaletteProperty.AccentForeground`.
	AccentForeground *string `json:"accentForeground"`
	// `CfnTheme.UIColorPaletteProperty.Danger`.
	Danger *string `json:"danger"`
	// `CfnTheme.UIColorPaletteProperty.DangerForeground`.
	DangerForeground *string `json:"dangerForeground"`
	// `CfnTheme.UIColorPaletteProperty.Dimension`.
	Dimension *string `json:"dimension"`
	// `CfnTheme.UIColorPaletteProperty.DimensionForeground`.
	DimensionForeground *string `json:"dimensionForeground"`
	// `CfnTheme.UIColorPaletteProperty.Measure`.
	Measure *string `json:"measure"`
	// `CfnTheme.UIColorPaletteProperty.MeasureForeground`.
	MeasureForeground *string `json:"measureForeground"`
	// `CfnTheme.UIColorPaletteProperty.PrimaryBackground`.
	PrimaryBackground *string `json:"primaryBackground"`
	// `CfnTheme.UIColorPaletteProperty.PrimaryForeground`.
	PrimaryForeground *string `json:"primaryForeground"`
	// `CfnTheme.UIColorPaletteProperty.SecondaryBackground`.
	SecondaryBackground *string `json:"secondaryBackground"`
	// `CfnTheme.UIColorPaletteProperty.SecondaryForeground`.
	SecondaryForeground *string `json:"secondaryForeground"`
	// `CfnTheme.UIColorPaletteProperty.Success`.
	Success *string `json:"success"`
	// `CfnTheme.UIColorPaletteProperty.SuccessForeground`.
	SuccessForeground *string `json:"successForeground"`
	// `CfnTheme.UIColorPaletteProperty.Warning`.
	Warning *string `json:"warning"`
	// `CfnTheme.UIColorPaletteProperty.WarningForeground`.
	WarningForeground *string `json:"warningForeground"`
}

// Properties for defining a `AWS::QuickSight::Theme`.
type CfnThemeProps struct {
	// `AWS::QuickSight::Theme.AwsAccountId`.
	AwsAccountId *string `json:"awsAccountId"`
	// `AWS::QuickSight::Theme.ThemeId`.
	ThemeId *string `json:"themeId"`
	// `AWS::QuickSight::Theme.BaseThemeId`.
	BaseThemeId *string `json:"baseThemeId"`
	// `AWS::QuickSight::Theme.Configuration`.
	Configuration interface{} `json:"configuration"`
	// `AWS::QuickSight::Theme.Name`.
	Name *string `json:"name"`
	// `AWS::QuickSight::Theme.Permissions`.
	Permissions interface{} `json:"permissions"`
	// `AWS::QuickSight::Theme.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::QuickSight::Theme.VersionDescription`.
	VersionDescription *string `json:"versionDescription"`
}


package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Kendra::DataSource`.
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataSourceConfiguration() interface{}
	SetDataSourceConfiguration(val interface{})
	Description() *string
	SetDescription(val *string)
	IndexId() *string
	SetIndexId(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Schedule() *string
	SetSchedule(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSource) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::DataSource`.
func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDataSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"dataSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetIndexId(val *string) {
	_jsii_.Set(
		j,
		"indexId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
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
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
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
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDataSource) ToString() *string {
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
func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDataSource_AccessControlListConfigurationProperty struct {
	// `CfnDataSource.AccessControlListConfigurationProperty.KeyPath`.
	KeyPath *string `json:"keyPath"`
}

type CfnDataSource_AclConfigurationProperty struct {
	// `CfnDataSource.AclConfigurationProperty.AllowedGroupsColumnName`.
	AllowedGroupsColumnName *string `json:"allowedGroupsColumnName"`
}

type CfnDataSource_ChangeDetectingColumnsProperty struct {
	// `CfnDataSource.ChangeDetectingColumnsProperty.ChangeDetectingColumns`.
	ChangeDetectingColumns *[]*string `json:"changeDetectingColumns"`
}

type CfnDataSource_ColumnConfigurationProperty struct {
	// `CfnDataSource.ColumnConfigurationProperty.ChangeDetectingColumns`.
	ChangeDetectingColumns interface{} `json:"changeDetectingColumns"`
	// `CfnDataSource.ColumnConfigurationProperty.DocumentDataColumnName`.
	DocumentDataColumnName *string `json:"documentDataColumnName"`
	// `CfnDataSource.ColumnConfigurationProperty.DocumentIdColumnName`.
	DocumentIdColumnName *string `json:"documentIdColumnName"`
	// `CfnDataSource.ColumnConfigurationProperty.DocumentTitleColumnName`.
	DocumentTitleColumnName *string `json:"documentTitleColumnName"`
	// `CfnDataSource.ColumnConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
}

type CfnDataSource_ConfluenceAttachmentConfigurationProperty struct {
	// `CfnDataSource.ConfluenceAttachmentConfigurationProperty.AttachmentFieldMappings`.
	AttachmentFieldMappings interface{} `json:"attachmentFieldMappings"`
	// `CfnDataSource.ConfluenceAttachmentConfigurationProperty.CrawlAttachments`.
	CrawlAttachments interface{} `json:"crawlAttachments"`
}

type CfnDataSource_ConfluenceAttachmentFieldMappingsListProperty struct {
	// `CfnDataSource.ConfluenceAttachmentFieldMappingsListProperty.ConfluenceAttachmentFieldMappingsList`.
	ConfluenceAttachmentFieldMappingsList interface{} `json:"confluenceAttachmentFieldMappingsList"`
}

type CfnDataSource_ConfluenceAttachmentToIndexFieldMappingProperty struct {
	// `CfnDataSource.ConfluenceAttachmentToIndexFieldMappingProperty.DataSourceFieldName`.
	DataSourceFieldName *string `json:"dataSourceFieldName"`
	// `CfnDataSource.ConfluenceAttachmentToIndexFieldMappingProperty.IndexFieldName`.
	IndexFieldName *string `json:"indexFieldName"`
	// `CfnDataSource.ConfluenceAttachmentToIndexFieldMappingProperty.DateFieldFormat`.
	DateFieldFormat *string `json:"dateFieldFormat"`
}

type CfnDataSource_ConfluenceBlogConfigurationProperty struct {
	// `CfnDataSource.ConfluenceBlogConfigurationProperty.BlogFieldMappings`.
	BlogFieldMappings interface{} `json:"blogFieldMappings"`
}

type CfnDataSource_ConfluenceBlogFieldMappingsListProperty struct {
	// `CfnDataSource.ConfluenceBlogFieldMappingsListProperty.ConfluenceBlogFieldMappingsList`.
	ConfluenceBlogFieldMappingsList interface{} `json:"confluenceBlogFieldMappingsList"`
}

type CfnDataSource_ConfluenceBlogToIndexFieldMappingProperty struct {
	// `CfnDataSource.ConfluenceBlogToIndexFieldMappingProperty.DataSourceFieldName`.
	DataSourceFieldName *string `json:"dataSourceFieldName"`
	// `CfnDataSource.ConfluenceBlogToIndexFieldMappingProperty.IndexFieldName`.
	IndexFieldName *string `json:"indexFieldName"`
	// `CfnDataSource.ConfluenceBlogToIndexFieldMappingProperty.DateFieldFormat`.
	DateFieldFormat *string `json:"dateFieldFormat"`
}

type CfnDataSource_ConfluenceConfigurationProperty struct {
	// `CfnDataSource.ConfluenceConfigurationProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDataSource.ConfluenceConfigurationProperty.ServerUrl`.
	ServerUrl *string `json:"serverUrl"`
	// `CfnDataSource.ConfluenceConfigurationProperty.Version`.
	Version *string `json:"version"`
	// `CfnDataSource.ConfluenceConfigurationProperty.AttachmentConfiguration`.
	AttachmentConfiguration interface{} `json:"attachmentConfiguration"`
	// `CfnDataSource.ConfluenceConfigurationProperty.BlogConfiguration`.
	BlogConfiguration interface{} `json:"blogConfiguration"`
	// `CfnDataSource.ConfluenceConfigurationProperty.ExclusionPatterns`.
	ExclusionPatterns interface{} `json:"exclusionPatterns"`
	// `CfnDataSource.ConfluenceConfigurationProperty.InclusionPatterns`.
	InclusionPatterns interface{} `json:"inclusionPatterns"`
	// `CfnDataSource.ConfluenceConfigurationProperty.PageConfiguration`.
	PageConfiguration interface{} `json:"pageConfiguration"`
	// `CfnDataSource.ConfluenceConfigurationProperty.SpaceConfiguration`.
	SpaceConfiguration interface{} `json:"spaceConfiguration"`
	// `CfnDataSource.ConfluenceConfigurationProperty.VpcConfiguration`.
	VpcConfiguration interface{} `json:"vpcConfiguration"`
}

type CfnDataSource_ConfluencePageConfigurationProperty struct {
	// `CfnDataSource.ConfluencePageConfigurationProperty.PageFieldMappings`.
	PageFieldMappings interface{} `json:"pageFieldMappings"`
}

type CfnDataSource_ConfluencePageFieldMappingsListProperty struct {
	// `CfnDataSource.ConfluencePageFieldMappingsListProperty.ConfluencePageFieldMappingsList`.
	ConfluencePageFieldMappingsList interface{} `json:"confluencePageFieldMappingsList"`
}

type CfnDataSource_ConfluencePageToIndexFieldMappingProperty struct {
	// `CfnDataSource.ConfluencePageToIndexFieldMappingProperty.DataSourceFieldName`.
	DataSourceFieldName *string `json:"dataSourceFieldName"`
	// `CfnDataSource.ConfluencePageToIndexFieldMappingProperty.IndexFieldName`.
	IndexFieldName *string `json:"indexFieldName"`
	// `CfnDataSource.ConfluencePageToIndexFieldMappingProperty.DateFieldFormat`.
	DateFieldFormat *string `json:"dateFieldFormat"`
}

type CfnDataSource_ConfluenceSpaceConfigurationProperty struct {
	// `CfnDataSource.ConfluenceSpaceConfigurationProperty.CrawlArchivedSpaces`.
	CrawlArchivedSpaces interface{} `json:"crawlArchivedSpaces"`
	// `CfnDataSource.ConfluenceSpaceConfigurationProperty.CrawlPersonalSpaces`.
	CrawlPersonalSpaces interface{} `json:"crawlPersonalSpaces"`
	// `CfnDataSource.ConfluenceSpaceConfigurationProperty.ExcludeSpaces`.
	ExcludeSpaces interface{} `json:"excludeSpaces"`
	// `CfnDataSource.ConfluenceSpaceConfigurationProperty.IncludeSpaces`.
	IncludeSpaces interface{} `json:"includeSpaces"`
	// `CfnDataSource.ConfluenceSpaceConfigurationProperty.SpaceFieldMappings`.
	SpaceFieldMappings interface{} `json:"spaceFieldMappings"`
}

type CfnDataSource_ConfluenceSpaceFieldMappingsListProperty struct {
	// `CfnDataSource.ConfluenceSpaceFieldMappingsListProperty.ConfluenceSpaceFieldMappingsList`.
	ConfluenceSpaceFieldMappingsList interface{} `json:"confluenceSpaceFieldMappingsList"`
}

type CfnDataSource_ConfluenceSpaceListProperty struct {
	// `CfnDataSource.ConfluenceSpaceListProperty.ConfluenceSpaceList`.
	ConfluenceSpaceList *[]*string `json:"confluenceSpaceList"`
}

type CfnDataSource_ConfluenceSpaceToIndexFieldMappingProperty struct {
	// `CfnDataSource.ConfluenceSpaceToIndexFieldMappingProperty.DataSourceFieldName`.
	DataSourceFieldName *string `json:"dataSourceFieldName"`
	// `CfnDataSource.ConfluenceSpaceToIndexFieldMappingProperty.IndexFieldName`.
	IndexFieldName *string `json:"indexFieldName"`
	// `CfnDataSource.ConfluenceSpaceToIndexFieldMappingProperty.DateFieldFormat`.
	DateFieldFormat *string `json:"dateFieldFormat"`
}

type CfnDataSource_ConnectionConfigurationProperty struct {
	// `CfnDataSource.ConnectionConfigurationProperty.DatabaseHost`.
	DatabaseHost *string `json:"databaseHost"`
	// `CfnDataSource.ConnectionConfigurationProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnDataSource.ConnectionConfigurationProperty.DatabasePort`.
	DatabasePort *float64 `json:"databasePort"`
	// `CfnDataSource.ConnectionConfigurationProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDataSource.ConnectionConfigurationProperty.TableName`.
	TableName *string `json:"tableName"`
}

type CfnDataSource_DataSourceConfigurationProperty struct {
	// `CfnDataSource.DataSourceConfigurationProperty.ConfluenceConfiguration`.
	ConfluenceConfiguration interface{} `json:"confluenceConfiguration"`
	// `CfnDataSource.DataSourceConfigurationProperty.DatabaseConfiguration`.
	DatabaseConfiguration interface{} `json:"databaseConfiguration"`
	// `CfnDataSource.DataSourceConfigurationProperty.GoogleDriveConfiguration`.
	GoogleDriveConfiguration interface{} `json:"googleDriveConfiguration"`
	// `CfnDataSource.DataSourceConfigurationProperty.OneDriveConfiguration`.
	OneDriveConfiguration interface{} `json:"oneDriveConfiguration"`
	// `CfnDataSource.DataSourceConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `json:"s3Configuration"`
	// `CfnDataSource.DataSourceConfigurationProperty.SalesforceConfiguration`.
	SalesforceConfiguration interface{} `json:"salesforceConfiguration"`
	// `CfnDataSource.DataSourceConfigurationProperty.ServiceNowConfiguration`.
	ServiceNowConfiguration interface{} `json:"serviceNowConfiguration"`
	// `CfnDataSource.DataSourceConfigurationProperty.SharePointConfiguration`.
	SharePointConfiguration interface{} `json:"sharePointConfiguration"`
}

type CfnDataSource_DataSourceInclusionsExclusionsStringsProperty struct {
	// `CfnDataSource.DataSourceInclusionsExclusionsStringsProperty.DataSourceInclusionsExclusionsStrings`.
	DataSourceInclusionsExclusionsStrings *[]*string `json:"dataSourceInclusionsExclusionsStrings"`
}

type CfnDataSource_DataSourceToIndexFieldMappingListProperty struct {
	// `CfnDataSource.DataSourceToIndexFieldMappingListProperty.DataSourceToIndexFieldMappingList`.
	DataSourceToIndexFieldMappingList interface{} `json:"dataSourceToIndexFieldMappingList"`
}

type CfnDataSource_DataSourceToIndexFieldMappingProperty struct {
	// `CfnDataSource.DataSourceToIndexFieldMappingProperty.DataSourceFieldName`.
	DataSourceFieldName *string `json:"dataSourceFieldName"`
	// `CfnDataSource.DataSourceToIndexFieldMappingProperty.IndexFieldName`.
	IndexFieldName *string `json:"indexFieldName"`
	// `CfnDataSource.DataSourceToIndexFieldMappingProperty.DateFieldFormat`.
	DateFieldFormat *string `json:"dateFieldFormat"`
}

type CfnDataSource_DataSourceVpcConfigurationProperty struct {
	// `CfnDataSource.DataSourceVpcConfigurationProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnDataSource.DataSourceVpcConfigurationProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
}

type CfnDataSource_DatabaseConfigurationProperty struct {
	// `CfnDataSource.DatabaseConfigurationProperty.ColumnConfiguration`.
	ColumnConfiguration interface{} `json:"columnConfiguration"`
	// `CfnDataSource.DatabaseConfigurationProperty.ConnectionConfiguration`.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	// `CfnDataSource.DatabaseConfigurationProperty.DatabaseEngineType`.
	DatabaseEngineType *string `json:"databaseEngineType"`
	// `CfnDataSource.DatabaseConfigurationProperty.AclConfiguration`.
	AclConfiguration interface{} `json:"aclConfiguration"`
	// `CfnDataSource.DatabaseConfigurationProperty.SqlConfiguration`.
	SqlConfiguration interface{} `json:"sqlConfiguration"`
	// `CfnDataSource.DatabaseConfigurationProperty.VpcConfiguration`.
	VpcConfiguration interface{} `json:"vpcConfiguration"`
}

type CfnDataSource_DocumentsMetadataConfigurationProperty struct {
	// `CfnDataSource.DocumentsMetadataConfigurationProperty.S3Prefix`.
	S3Prefix *string `json:"s3Prefix"`
}

type CfnDataSource_ExcludeMimeTypesListProperty struct {
	// `CfnDataSource.ExcludeMimeTypesListProperty.ExcludeMimeTypesList`.
	ExcludeMimeTypesList *[]*string `json:"excludeMimeTypesList"`
}

type CfnDataSource_ExcludeSharedDrivesListProperty struct {
	// `CfnDataSource.ExcludeSharedDrivesListProperty.ExcludeSharedDrivesList`.
	ExcludeSharedDrivesList *[]*string `json:"excludeSharedDrivesList"`
}

type CfnDataSource_ExcludeUserAccountsListProperty struct {
	// `CfnDataSource.ExcludeUserAccountsListProperty.ExcludeUserAccountsList`.
	ExcludeUserAccountsList *[]*string `json:"excludeUserAccountsList"`
}

type CfnDataSource_GoogleDriveConfigurationProperty struct {
	// `CfnDataSource.GoogleDriveConfigurationProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDataSource.GoogleDriveConfigurationProperty.ExcludeMimeTypes`.
	ExcludeMimeTypes interface{} `json:"excludeMimeTypes"`
	// `CfnDataSource.GoogleDriveConfigurationProperty.ExcludeSharedDrives`.
	ExcludeSharedDrives interface{} `json:"excludeSharedDrives"`
	// `CfnDataSource.GoogleDriveConfigurationProperty.ExcludeUserAccounts`.
	ExcludeUserAccounts interface{} `json:"excludeUserAccounts"`
	// `CfnDataSource.GoogleDriveConfigurationProperty.ExclusionPatterns`.
	ExclusionPatterns interface{} `json:"exclusionPatterns"`
	// `CfnDataSource.GoogleDriveConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
	// `CfnDataSource.GoogleDriveConfigurationProperty.InclusionPatterns`.
	InclusionPatterns interface{} `json:"inclusionPatterns"`
}

type CfnDataSource_OneDriveConfigurationProperty struct {
	// `CfnDataSource.OneDriveConfigurationProperty.OneDriveUsers`.
	OneDriveUsers interface{} `json:"oneDriveUsers"`
	// `CfnDataSource.OneDriveConfigurationProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDataSource.OneDriveConfigurationProperty.TenantDomain`.
	TenantDomain *string `json:"tenantDomain"`
	// `CfnDataSource.OneDriveConfigurationProperty.DisableLocalGroups`.
	DisableLocalGroups interface{} `json:"disableLocalGroups"`
	// `CfnDataSource.OneDriveConfigurationProperty.ExclusionPatterns`.
	ExclusionPatterns interface{} `json:"exclusionPatterns"`
	// `CfnDataSource.OneDriveConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
	// `CfnDataSource.OneDriveConfigurationProperty.InclusionPatterns`.
	InclusionPatterns interface{} `json:"inclusionPatterns"`
}

type CfnDataSource_OneDriveUserListProperty struct {
	// `CfnDataSource.OneDriveUserListProperty.OneDriveUserList`.
	OneDriveUserList *[]*string `json:"oneDriveUserList"`
}

type CfnDataSource_OneDriveUsersProperty struct {
	// `CfnDataSource.OneDriveUsersProperty.OneDriveUserList`.
	OneDriveUserList interface{} `json:"oneDriveUserList"`
	// `CfnDataSource.OneDriveUsersProperty.OneDriveUserS3Path`.
	OneDriveUserS3Path interface{} `json:"oneDriveUserS3Path"`
}

type CfnDataSource_S3DataSourceConfigurationProperty struct {
	// `CfnDataSource.S3DataSourceConfigurationProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnDataSource.S3DataSourceConfigurationProperty.AccessControlListConfiguration`.
	AccessControlListConfiguration interface{} `json:"accessControlListConfiguration"`
	// `CfnDataSource.S3DataSourceConfigurationProperty.DocumentsMetadataConfiguration`.
	DocumentsMetadataConfiguration interface{} `json:"documentsMetadataConfiguration"`
	// `CfnDataSource.S3DataSourceConfigurationProperty.ExclusionPatterns`.
	ExclusionPatterns interface{} `json:"exclusionPatterns"`
	// `CfnDataSource.S3DataSourceConfigurationProperty.InclusionPatterns`.
	InclusionPatterns interface{} `json:"inclusionPatterns"`
	// `CfnDataSource.S3DataSourceConfigurationProperty.InclusionPrefixes`.
	InclusionPrefixes interface{} `json:"inclusionPrefixes"`
}

type CfnDataSource_S3PathProperty struct {
	// `CfnDataSource.S3PathProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnDataSource.S3PathProperty.Key`.
	Key *string `json:"key"`
}

type CfnDataSource_SalesforceChatterFeedConfigurationProperty struct {
	// `CfnDataSource.SalesforceChatterFeedConfigurationProperty.DocumentDataFieldName`.
	DocumentDataFieldName *string `json:"documentDataFieldName"`
	// `CfnDataSource.SalesforceChatterFeedConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.SalesforceChatterFeedConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
	// `CfnDataSource.SalesforceChatterFeedConfigurationProperty.IncludeFilterTypes`.
	IncludeFilterTypes interface{} `json:"includeFilterTypes"`
}

type CfnDataSource_SalesforceChatterFeedIncludeFilterTypesProperty struct {
	// `CfnDataSource.SalesforceChatterFeedIncludeFilterTypesProperty.SalesforceChatterFeedIncludeFilterTypes`.
	SalesforceChatterFeedIncludeFilterTypes *[]*string `json:"salesforceChatterFeedIncludeFilterTypes"`
}

type CfnDataSource_SalesforceConfigurationProperty struct {
	// `CfnDataSource.SalesforceConfigurationProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDataSource.SalesforceConfigurationProperty.ServerUrl`.
	ServerUrl *string `json:"serverUrl"`
	// `CfnDataSource.SalesforceConfigurationProperty.ChatterFeedConfiguration`.
	ChatterFeedConfiguration interface{} `json:"chatterFeedConfiguration"`
	// `CfnDataSource.SalesforceConfigurationProperty.CrawlAttachments`.
	CrawlAttachments interface{} `json:"crawlAttachments"`
	// `CfnDataSource.SalesforceConfigurationProperty.ExcludeAttachmentFilePatterns`.
	ExcludeAttachmentFilePatterns interface{} `json:"excludeAttachmentFilePatterns"`
	// `CfnDataSource.SalesforceConfigurationProperty.IncludeAttachmentFilePatterns`.
	IncludeAttachmentFilePatterns interface{} `json:"includeAttachmentFilePatterns"`
	// `CfnDataSource.SalesforceConfigurationProperty.KnowledgeArticleConfiguration`.
	KnowledgeArticleConfiguration interface{} `json:"knowledgeArticleConfiguration"`
	// `CfnDataSource.SalesforceConfigurationProperty.StandardObjectAttachmentConfiguration`.
	StandardObjectAttachmentConfiguration interface{} `json:"standardObjectAttachmentConfiguration"`
	// `CfnDataSource.SalesforceConfigurationProperty.StandardObjectConfigurations`.
	StandardObjectConfigurations interface{} `json:"standardObjectConfigurations"`
}

type CfnDataSource_SalesforceCustomKnowledgeArticleTypeConfigurationListProperty struct {
	// `CfnDataSource.SalesforceCustomKnowledgeArticleTypeConfigurationListProperty.SalesforceCustomKnowledgeArticleTypeConfigurationList`.
	SalesforceCustomKnowledgeArticleTypeConfigurationList interface{} `json:"salesforceCustomKnowledgeArticleTypeConfigurationList"`
}

type CfnDataSource_SalesforceCustomKnowledgeArticleTypeConfigurationProperty struct {
	// `CfnDataSource.SalesforceCustomKnowledgeArticleTypeConfigurationProperty.DocumentDataFieldName`.
	DocumentDataFieldName *string `json:"documentDataFieldName"`
	// `CfnDataSource.SalesforceCustomKnowledgeArticleTypeConfigurationProperty.Name`.
	Name *string `json:"name"`
	// `CfnDataSource.SalesforceCustomKnowledgeArticleTypeConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.SalesforceCustomKnowledgeArticleTypeConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
}

type CfnDataSource_SalesforceKnowledgeArticleConfigurationProperty struct {
	// `CfnDataSource.SalesforceKnowledgeArticleConfigurationProperty.IncludedStates`.
	IncludedStates interface{} `json:"includedStates"`
	// `CfnDataSource.SalesforceKnowledgeArticleConfigurationProperty.CustomKnowledgeArticleTypeConfigurations`.
	CustomKnowledgeArticleTypeConfigurations interface{} `json:"customKnowledgeArticleTypeConfigurations"`
	// `CfnDataSource.SalesforceKnowledgeArticleConfigurationProperty.StandardKnowledgeArticleTypeConfiguration`.
	StandardKnowledgeArticleTypeConfiguration interface{} `json:"standardKnowledgeArticleTypeConfiguration"`
}

type CfnDataSource_SalesforceKnowledgeArticleStateListProperty struct {
	// `CfnDataSource.SalesforceKnowledgeArticleStateListProperty.SalesforceKnowledgeArticleStateList`.
	SalesforceKnowledgeArticleStateList *[]*string `json:"salesforceKnowledgeArticleStateList"`
}

type CfnDataSource_SalesforceStandardKnowledgeArticleTypeConfigurationProperty struct {
	// `CfnDataSource.SalesforceStandardKnowledgeArticleTypeConfigurationProperty.DocumentDataFieldName`.
	DocumentDataFieldName *string `json:"documentDataFieldName"`
	// `CfnDataSource.SalesforceStandardKnowledgeArticleTypeConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.SalesforceStandardKnowledgeArticleTypeConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
}

type CfnDataSource_SalesforceStandardObjectAttachmentConfigurationProperty struct {
	// `CfnDataSource.SalesforceStandardObjectAttachmentConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.SalesforceStandardObjectAttachmentConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
}

type CfnDataSource_SalesforceStandardObjectConfigurationListProperty struct {
	// `CfnDataSource.SalesforceStandardObjectConfigurationListProperty.SalesforceStandardObjectConfigurationList`.
	SalesforceStandardObjectConfigurationList interface{} `json:"salesforceStandardObjectConfigurationList"`
}

type CfnDataSource_SalesforceStandardObjectConfigurationProperty struct {
	// `CfnDataSource.SalesforceStandardObjectConfigurationProperty.DocumentDataFieldName`.
	DocumentDataFieldName *string `json:"documentDataFieldName"`
	// `CfnDataSource.SalesforceStandardObjectConfigurationProperty.Name`.
	Name *string `json:"name"`
	// `CfnDataSource.SalesforceStandardObjectConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.SalesforceStandardObjectConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
}

type CfnDataSource_ServiceNowConfigurationProperty struct {
	// `CfnDataSource.ServiceNowConfigurationProperty.HostUrl`.
	HostUrl *string `json:"hostUrl"`
	// `CfnDataSource.ServiceNowConfigurationProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDataSource.ServiceNowConfigurationProperty.ServiceNowBuildVersion`.
	ServiceNowBuildVersion *string `json:"serviceNowBuildVersion"`
	// `CfnDataSource.ServiceNowConfigurationProperty.KnowledgeArticleConfiguration`.
	KnowledgeArticleConfiguration interface{} `json:"knowledgeArticleConfiguration"`
	// `CfnDataSource.ServiceNowConfigurationProperty.ServiceCatalogConfiguration`.
	ServiceCatalogConfiguration interface{} `json:"serviceCatalogConfiguration"`
}

type CfnDataSource_ServiceNowKnowledgeArticleConfigurationProperty struct {
	// `CfnDataSource.ServiceNowKnowledgeArticleConfigurationProperty.DocumentDataFieldName`.
	DocumentDataFieldName *string `json:"documentDataFieldName"`
	// `CfnDataSource.ServiceNowKnowledgeArticleConfigurationProperty.CrawlAttachments`.
	CrawlAttachments interface{} `json:"crawlAttachments"`
	// `CfnDataSource.ServiceNowKnowledgeArticleConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.ServiceNowKnowledgeArticleConfigurationProperty.ExcludeAttachmentFilePatterns`.
	ExcludeAttachmentFilePatterns interface{} `json:"excludeAttachmentFilePatterns"`
	// `CfnDataSource.ServiceNowKnowledgeArticleConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
	// `CfnDataSource.ServiceNowKnowledgeArticleConfigurationProperty.IncludeAttachmentFilePatterns`.
	IncludeAttachmentFilePatterns interface{} `json:"includeAttachmentFilePatterns"`
}

type CfnDataSource_ServiceNowServiceCatalogConfigurationProperty struct {
	// `CfnDataSource.ServiceNowServiceCatalogConfigurationProperty.DocumentDataFieldName`.
	DocumentDataFieldName *string `json:"documentDataFieldName"`
	// `CfnDataSource.ServiceNowServiceCatalogConfigurationProperty.CrawlAttachments`.
	CrawlAttachments interface{} `json:"crawlAttachments"`
	// `CfnDataSource.ServiceNowServiceCatalogConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.ServiceNowServiceCatalogConfigurationProperty.ExcludeAttachmentFilePatterns`.
	ExcludeAttachmentFilePatterns interface{} `json:"excludeAttachmentFilePatterns"`
	// `CfnDataSource.ServiceNowServiceCatalogConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
	// `CfnDataSource.ServiceNowServiceCatalogConfigurationProperty.IncludeAttachmentFilePatterns`.
	IncludeAttachmentFilePatterns interface{} `json:"includeAttachmentFilePatterns"`
}

type CfnDataSource_SharePointConfigurationProperty struct {
	// `CfnDataSource.SharePointConfigurationProperty.SecretArn`.
	SecretArn *string `json:"secretArn"`
	// `CfnDataSource.SharePointConfigurationProperty.SharePointVersion`.
	SharePointVersion *string `json:"sharePointVersion"`
	// `CfnDataSource.SharePointConfigurationProperty.Urls`.
	Urls *[]*string `json:"urls"`
	// `CfnDataSource.SharePointConfigurationProperty.CrawlAttachments`.
	CrawlAttachments interface{} `json:"crawlAttachments"`
	// `CfnDataSource.SharePointConfigurationProperty.DisableLocalGroups`.
	DisableLocalGroups interface{} `json:"disableLocalGroups"`
	// `CfnDataSource.SharePointConfigurationProperty.DocumentTitleFieldName`.
	DocumentTitleFieldName *string `json:"documentTitleFieldName"`
	// `CfnDataSource.SharePointConfigurationProperty.ExclusionPatterns`.
	ExclusionPatterns interface{} `json:"exclusionPatterns"`
	// `CfnDataSource.SharePointConfigurationProperty.FieldMappings`.
	FieldMappings interface{} `json:"fieldMappings"`
	// `CfnDataSource.SharePointConfigurationProperty.InclusionPatterns`.
	InclusionPatterns interface{} `json:"inclusionPatterns"`
	// `CfnDataSource.SharePointConfigurationProperty.UseChangeLog`.
	UseChangeLog interface{} `json:"useChangeLog"`
	// `CfnDataSource.SharePointConfigurationProperty.VpcConfiguration`.
	VpcConfiguration interface{} `json:"vpcConfiguration"`
}

type CfnDataSource_SqlConfigurationProperty struct {
	// `CfnDataSource.SqlConfigurationProperty.QueryIdentifiersEnclosingOption`.
	QueryIdentifiersEnclosingOption *string `json:"queryIdentifiersEnclosingOption"`
}

// Properties for defining a `AWS::Kendra::DataSource`.
type CfnDataSourceProps struct {
	// `AWS::Kendra::DataSource.IndexId`.
	IndexId *string `json:"indexId"`
	// `AWS::Kendra::DataSource.Name`.
	Name *string `json:"name"`
	// `AWS::Kendra::DataSource.Type`.
	Type *string `json:"type"`
	// `AWS::Kendra::DataSource.DataSourceConfiguration`.
	DataSourceConfiguration interface{} `json:"dataSourceConfiguration"`
	// `AWS::Kendra::DataSource.Description`.
	Description *string `json:"description"`
	// `AWS::Kendra::DataSource.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::Kendra::DataSource.Schedule`.
	Schedule *string `json:"schedule"`
	// `AWS::Kendra::DataSource.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::Kendra::Faq`.
type CfnFaq interface {
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
	FileFormat() *string
	SetFileFormat(val *string)
	IndexId() *string
	SetIndexId(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	S3Path() interface{}
	SetS3Path(val interface{})
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

// The jsii proxy struct for CfnFaq
type jsiiProxy_CfnFaq struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFaq) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) FileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) S3Path() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::Faq`.
func NewCfnFaq(scope constructs.Construct, id *string, props *CfnFaqProps) CfnFaq {
	_init_.Initialize()

	j := jsiiProxy_CfnFaq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::Faq`.
func NewCfnFaq_Override(c CfnFaq, scope constructs.Construct, id *string, props *CfnFaqProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFaq) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetFileFormat(val *string) {
	_jsii_.Set(
		j,
		"fileFormat",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetIndexId(val *string) {
	_jsii_.Set(
		j,
		"indexId",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetS3Path(val interface{}) {
	_jsii_.Set(
		j,
		"s3Path",
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
func CfnFaq_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFaq_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
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
func CfnFaq_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFaq_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFaq) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFaq) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFaq) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFaq) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFaq) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnFaq) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnFaq) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFaq) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFaq) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFaq) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnFaq) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFaq) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFaq) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFaq) ToString() *string {
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
func (c *jsiiProxy_CfnFaq) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFaq_S3PathProperty struct {
	// `CfnFaq.S3PathProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnFaq.S3PathProperty.Key`.
	Key *string `json:"key"`
}

// Properties for defining a `AWS::Kendra::Faq`.
type CfnFaqProps struct {
	// `AWS::Kendra::Faq.IndexId`.
	IndexId *string `json:"indexId"`
	// `AWS::Kendra::Faq.Name`.
	Name *string `json:"name"`
	// `AWS::Kendra::Faq.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::Kendra::Faq.S3Path`.
	S3Path interface{} `json:"s3Path"`
	// `AWS::Kendra::Faq.Description`.
	Description *string `json:"description"`
	// `AWS::Kendra::Faq.FileFormat`.
	FileFormat *string `json:"fileFormat"`
	// `AWS::Kendra::Faq.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::Kendra::Index`.
type CfnIndex interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CapacityUnits() interface{}
	SetCapacityUnits(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DocumentMetadataConfigurations() interface{}
	SetDocumentMetadataConfigurations(val interface{})
	Edition() *string
	SetEdition(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	ServerSideEncryptionConfiguration() interface{}
	SetServerSideEncryptionConfiguration(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UserContextPolicy() *string
	SetUserContextPolicy(val *string)
	UserTokenConfigurations() interface{}
	SetUserTokenConfigurations(val interface{})
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

// The jsii proxy struct for CfnIndex
type jsiiProxy_CfnIndex struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIndex) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CapacityUnits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) DocumentMetadataConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentMetadataConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) ServerSideEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserContextPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userContextPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserTokenConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userTokenConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::Index`.
func NewCfnIndex(scope constructs.Construct, id *string, props *CfnIndexProps) CfnIndex {
	_init_.Initialize()

	j := jsiiProxy_CfnIndex{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::Index`.
func NewCfnIndex_Override(c CfnIndex, scope constructs.Construct, id *string, props *CfnIndexProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIndex) SetCapacityUnits(val interface{}) {
	_jsii_.Set(
		j,
		"capacityUnits",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetDocumentMetadataConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"documentMetadataConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetEdition(val *string) {
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetServerSideEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"serverSideEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetUserContextPolicy(val *string) {
	_jsii_.Set(
		j,
		"userContextPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetUserTokenConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"userTokenConfigurations",
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
func CfnIndex_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIndex_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
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
func CfnIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIndex_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnIndex) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnIndex) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnIndex) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnIndex) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnIndex) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnIndex) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnIndex) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnIndex) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnIndex) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnIndex) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnIndex) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIndex) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnIndex) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnIndex) ToString() *string {
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
func (c *jsiiProxy_CfnIndex) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnIndex_CapacityUnitsConfigurationProperty struct {
	// `CfnIndex.CapacityUnitsConfigurationProperty.QueryCapacityUnits`.
	QueryCapacityUnits *float64 `json:"queryCapacityUnits"`
	// `CfnIndex.CapacityUnitsConfigurationProperty.StorageCapacityUnits`.
	StorageCapacityUnits *float64 `json:"storageCapacityUnits"`
}

type CfnIndex_DocumentMetadataConfigurationProperty struct {
	// `CfnIndex.DocumentMetadataConfigurationProperty.Name`.
	Name *string `json:"name"`
	// `CfnIndex.DocumentMetadataConfigurationProperty.Type`.
	Type *string `json:"type"`
	// `CfnIndex.DocumentMetadataConfigurationProperty.Relevance`.
	Relevance interface{} `json:"relevance"`
	// `CfnIndex.DocumentMetadataConfigurationProperty.Search`.
	Search interface{} `json:"search"`
}

type CfnIndex_JsonTokenTypeConfigurationProperty struct {
	// `CfnIndex.JsonTokenTypeConfigurationProperty.GroupAttributeField`.
	GroupAttributeField *string `json:"groupAttributeField"`
	// `CfnIndex.JsonTokenTypeConfigurationProperty.UserNameAttributeField`.
	UserNameAttributeField *string `json:"userNameAttributeField"`
}

type CfnIndex_JwtTokenTypeConfigurationProperty struct {
	// `CfnIndex.JwtTokenTypeConfigurationProperty.KeyLocation`.
	KeyLocation *string `json:"keyLocation"`
	// `CfnIndex.JwtTokenTypeConfigurationProperty.ClaimRegex`.
	ClaimRegex *string `json:"claimRegex"`
	// `CfnIndex.JwtTokenTypeConfigurationProperty.GroupAttributeField`.
	GroupAttributeField *string `json:"groupAttributeField"`
	// `CfnIndex.JwtTokenTypeConfigurationProperty.Issuer`.
	Issuer *string `json:"issuer"`
	// `CfnIndex.JwtTokenTypeConfigurationProperty.SecretManagerArn`.
	SecretManagerArn *string `json:"secretManagerArn"`
	// `CfnIndex.JwtTokenTypeConfigurationProperty.URL`.
	Url *string `json:"url"`
	// `CfnIndex.JwtTokenTypeConfigurationProperty.UserNameAttributeField`.
	UserNameAttributeField *string `json:"userNameAttributeField"`
}

type CfnIndex_RelevanceProperty struct {
	// `CfnIndex.RelevanceProperty.Duration`.
	Duration *string `json:"duration"`
	// `CfnIndex.RelevanceProperty.Freshness`.
	Freshness interface{} `json:"freshness"`
	// `CfnIndex.RelevanceProperty.Importance`.
	Importance *float64 `json:"importance"`
	// `CfnIndex.RelevanceProperty.RankOrder`.
	RankOrder *string `json:"rankOrder"`
	// `CfnIndex.RelevanceProperty.ValueImportanceItems`.
	ValueImportanceItems interface{} `json:"valueImportanceItems"`
}

type CfnIndex_SearchProperty struct {
	// `CfnIndex.SearchProperty.Displayable`.
	Displayable interface{} `json:"displayable"`
	// `CfnIndex.SearchProperty.Facetable`.
	Facetable interface{} `json:"facetable"`
	// `CfnIndex.SearchProperty.Searchable`.
	Searchable interface{} `json:"searchable"`
	// `CfnIndex.SearchProperty.Sortable`.
	Sortable interface{} `json:"sortable"`
}

type CfnIndex_ServerSideEncryptionConfigurationProperty struct {
	// `CfnIndex.ServerSideEncryptionConfigurationProperty.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
}

type CfnIndex_UserTokenConfigurationProperty struct {
	// `CfnIndex.UserTokenConfigurationProperty.JsonTokenTypeConfiguration`.
	JsonTokenTypeConfiguration interface{} `json:"jsonTokenTypeConfiguration"`
	// `CfnIndex.UserTokenConfigurationProperty.JwtTokenTypeConfiguration`.
	JwtTokenTypeConfiguration interface{} `json:"jwtTokenTypeConfiguration"`
}

type CfnIndex_ValueImportanceItemProperty struct {
	// `CfnIndex.ValueImportanceItemProperty.Key`.
	Key *string `json:"key"`
	// `CfnIndex.ValueImportanceItemProperty.Value`.
	Value *float64 `json:"value"`
}

type CfnIndex_ValueImportanceItemsProperty struct {
	// `CfnIndex.ValueImportanceItemsProperty.ValueImportanceItems`.
	ValueImportanceItems interface{} `json:"valueImportanceItems"`
}

// Properties for defining a `AWS::Kendra::Index`.
type CfnIndexProps struct {
	// `AWS::Kendra::Index.Edition`.
	Edition *string `json:"edition"`
	// `AWS::Kendra::Index.Name`.
	Name *string `json:"name"`
	// `AWS::Kendra::Index.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `AWS::Kendra::Index.CapacityUnits`.
	CapacityUnits interface{} `json:"capacityUnits"`
	// `AWS::Kendra::Index.Description`.
	Description *string `json:"description"`
	// `AWS::Kendra::Index.DocumentMetadataConfigurations`.
	DocumentMetadataConfigurations interface{} `json:"documentMetadataConfigurations"`
	// `AWS::Kendra::Index.ServerSideEncryptionConfiguration`.
	ServerSideEncryptionConfiguration interface{} `json:"serverSideEncryptionConfiguration"`
	// `AWS::Kendra::Index.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::Kendra::Index.UserContextPolicy`.
	UserContextPolicy *string `json:"userContextPolicy"`
	// `AWS::Kendra::Index.UserTokenConfigurations`.
	UserTokenConfigurations interface{} `json:"userTokenConfigurations"`
}


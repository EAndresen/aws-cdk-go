package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a custom resource that is materialized using specific AWS API calls.
//
// Use this to bridge any gap that might exist in the CloudFormation Coverage.
// You can specify exactly which calls are invoked for the 'CREATE', 'UPDATE' and 'DELETE' life cycle events.
// Experimental.
type AwsCustomResource interface {
	constructs.Construct
	awsiam.IGrantable
	GrantPrincipal() awsiam.IPrincipal
	Node() constructs.Node
	GetResponseField(dataPath *string) *string
	GetResponseFieldReference(dataPath *string) awscdk.Reference
	ToString() *string
}

// The jsii proxy struct for AwsCustomResource
type jsiiProxy_AwsCustomResource struct {
	internal.Type__constructsConstruct
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_AwsCustomResource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCustomResource(scope constructs.Construct, id *string, props *AwsCustomResourceProps) AwsCustomResource {
	_init_.Initialize()

	j := jsiiProxy_AwsCustomResource{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCustomResource_Override(a AwsCustomResource, scope constructs.Construct, id *string, props *AwsCustomResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AwsCustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns response data for the AWS SDK call as string.
//
// Example for S3 / listBucket : 'Buckets.0.Name'
//
// Note that you cannot use this method if `ignoreErrorCodesMatching`
// is configured for any of the SDK calls. This is because in such a case,
// the response data might not exist, and will cause a CloudFormation deploy time error.
// Experimental.
func (a *jsiiProxy_AwsCustomResource) GetResponseField(dataPath *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResponseField",
		[]interface{}{dataPath},
		&returns,
	)

	return returns
}

// Returns response data for the AWS SDK call.
//
// Example for S3 / listBucket : 'Buckets.0.Name'
//
// Use `Token.asXxx` to encode the returned `Reference` as a specific type or
// use the convenience `getDataString` for string attributes.
//
// Note that you cannot use this method if `ignoreErrorCodesMatching`
// is configured for any of the SDK calls. This is because in such a case,
// the response data might not exist, and will cause a CloudFormation deploy time error.
// Experimental.
func (a *jsiiProxy_AwsCustomResource) GetResponseFieldReference(dataPath *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		a,
		"getResponseFieldReference",
		[]interface{}{dataPath},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AwsCustomResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The IAM Policy that will be applied to the different calls.
// Experimental.
type AwsCustomResourcePolicy interface {
	Resources() *[]*string
	Statements() *[]awsiam.PolicyStatement
}

// The jsii proxy struct for AwsCustomResourcePolicy
type jsiiProxy_AwsCustomResourcePolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_AwsCustomResourcePolicy) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomResourcePolicy) Statements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"statements",
		&returns,
	)
	return returns
}


// Generate IAM Policy Statements from the configured SDK calls.
//
// Each SDK call with be translated to an IAM Policy Statement in the form of: `call.service:call.action` (e.g `s3:PutObject`).
// Experimental.
func AwsCustomResourcePolicy_FromSdkCalls(options *SdkCallsPolicyOptions) AwsCustomResourcePolicy {
	_init_.Initialize()

	var returns AwsCustomResourcePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.AwsCustomResourcePolicy",
		"fromSdkCalls",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Explicit IAM Policy Statements.
// Experimental.
func AwsCustomResourcePolicy_FromStatements(statements *[]awsiam.PolicyStatement) AwsCustomResourcePolicy {
	_init_.Initialize()

	var returns AwsCustomResourcePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.AwsCustomResourcePolicy",
		"fromStatements",
		[]interface{}{statements},
		&returns,
	)

	return returns
}

func AwsCustomResourcePolicy_ANY_RESOURCE() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"aws-cdk-lib.custom_resources.AwsCustomResourcePolicy",
		"ANY_RESOURCE",
		&returns,
	)
	return returns
}

// Properties for AwsCustomResource.
//
// Note that at least onCreate, onUpdate or onDelete must be specified.
// Experimental.
type AwsCustomResourceProps struct {
	// The policy that will be added to the execution role of the Lambda function implementing this custom resource provider.
	//
	// The custom resource also implements `iam.IGrantable`, making it possible
	// to use the `grantXxx()` methods.
	//
	// As this custom resource uses a singleton Lambda function, it's important
	// to note the that function's role will eventually accumulate the
	// permissions/grants from all resources.
	// See: Policy.fromSdkCalls
	//
	// Experimental.
	Policy AwsCustomResourcePolicy `json:"policy"`
	// A name for the Lambda function implementing this custom resource.
	// Experimental.
	FunctionName *string `json:"functionName"`
	// Whether to install the latest AWS SDK v2. Allows to use the latest API calls documented at https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html.
	//
	// The installation takes around 60 seconds.
	// Experimental.
	InstallLatestAwsSdk *bool `json:"installLatestAwsSdk"`
	// The number of days log events of the Lambda function implementing this custom resource are kept in CloudWatch Logs.
	// Experimental.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// The AWS SDK call to make when the resource is created.
	// Experimental.
	OnCreate *AwsSdkCall `json:"onCreate"`
	// The AWS SDK call to make when the resource is deleted.
	// Experimental.
	OnDelete *AwsSdkCall `json:"onDelete"`
	// The AWS SDK call to make when the resource is updated.
	// Experimental.
	OnUpdate *AwsSdkCall `json:"onUpdate"`
	// Cloudformation Resource type.
	// Experimental.
	ResourceType *string `json:"resourceType"`
	// The execution role for the Lambda function implementing this custom resource provider.
	//
	// This role will apply to all `AwsCustomResource`
	// instances in the stack. The role must be assumable by the
	// `lambda.amazonaws.com` service principal.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The timeout for the Lambda function implementing this custom resource.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
}

// An AWS SDK call.
// Experimental.
type AwsSdkCall struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Action *string `json:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Service *string `json:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	// Experimental.
	ApiVersion *string `json:"apiVersion"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	// Experimental.
	IgnoreErrorCodesMatching *string `json:"ignoreErrorCodesMatching"`
	// Restrict the data returned by the custom resource to a specific path in the API response.
	//
	// Use this to limit the data returned by the custom
	// resource if working with API calls that could potentially result in custom
	// response objects exceeding the hard limit of 4096 bytes.
	//
	// Example for ECS / updateService: 'service.deploymentConfiguration.maximumPercent'
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Parameters interface{} `json:"parameters"`
	// The physical resource id of the custom resource for this call.
	//
	// Mandatory for onCreate or onUpdate calls.
	// Experimental.
	PhysicalResourceId PhysicalResourceId `json:"physicalResourceId"`
	// The region to send service requests to.
	//
	// **Note: Cross-region operations are generally considered an anti-pattern.**
	// **Consider first deploying a stack in that region.**
	// Experimental.
	Region *string `json:"region"`
}

// Physical ID of the custom resource.
// Experimental.
type PhysicalResourceId interface {
	Id() *string
	ResponsePath() *string
}

// The jsii proxy struct for PhysicalResourceId
type jsiiProxy_PhysicalResourceId struct {
	_ byte // padding
}

func (j *jsiiProxy_PhysicalResourceId) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalResourceId) ResponsePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responsePath",
		&returns,
	)
	return returns
}


// Extract the physical resource id from the path (dot notation) to the data in the API call response.
// Experimental.
func PhysicalResourceId_FromResponse(responsePath *string) PhysicalResourceId {
	_init_.Initialize()

	var returns PhysicalResourceId

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.PhysicalResourceId",
		"fromResponse",
		[]interface{}{responsePath},
		&returns,
	)

	return returns
}

// Explicit physical resource id.
// Experimental.
func PhysicalResourceId_Of(id *string) PhysicalResourceId {
	_init_.Initialize()

	var returns PhysicalResourceId

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.PhysicalResourceId",
		"of",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Reference to the physical resource id that can be passed to the AWS operation as a parameter.
// Experimental.
type PhysicalResourceIdReference interface {
	awscdk.IResolvable
	CreationStack() *[]*string
	Resolve(_arg awscdk.IResolveContext) interface{}
	ToJSON() *string
	ToString() *string
}

// The jsii proxy struct for PhysicalResourceIdReference
type jsiiProxy_PhysicalResourceIdReference struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_PhysicalResourceIdReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


// Experimental.
func NewPhysicalResourceIdReference() PhysicalResourceIdReference {
	_init_.Initialize()

	j := jsiiProxy_PhysicalResourceIdReference{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.PhysicalResourceIdReference",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewPhysicalResourceIdReference_Override(p PhysicalResourceIdReference) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.PhysicalResourceIdReference",
		nil, // no parameters
		p,
	)
}

// Produce the Token's value at resolution time.
// Experimental.
func (p *jsiiProxy_PhysicalResourceIdReference) Resolve(_arg awscdk.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// toJSON serialization to replace `PhysicalResourceIdReference` with a magic string.
// Experimental.
func (p *jsiiProxy_PhysicalResourceIdReference) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Return a string representation of this resolvable object.
//
// Returns a reversible string representation.
// Experimental.
func (p *jsiiProxy_PhysicalResourceIdReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Defines an AWS CloudFormation custom resource provider.
// Experimental.
type Provider interface {
	constructs.Construct
	awscloudformation.ICustomResourceProvider
	IsCompleteHandler() awslambda.IFunction
	Node() constructs.Node
	OnEventHandler() awslambda.IFunction
	ServiceToken() *string
	Bind(_scope constructs.Construct) *awscloudformation.CustomResourceProviderConfig
	ToString() *string
}

// The jsii proxy struct for Provider
type jsiiProxy_Provider struct {
	internal.Type__constructsConstruct
	internal.Type__awscloudformationICustomResourceProvider
}

func (j *jsiiProxy_Provider) IsCompleteHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"isCompleteHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) OnEventHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"onEventHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


// Experimental.
func NewProvider(scope constructs.Construct, id *string, props *ProviderProps) Provider {
	_init_.Initialize()

	j := jsiiProxy_Provider{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.Provider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewProvider_Override(p Provider, scope constructs.Construct, id *string, props *ProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.Provider",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Provider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.Provider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Called by `CustomResource` which uses this provider.
// Deprecated: use `provider.serviceToken` instead
func (p *jsiiProxy_Provider) Bind(_scope constructs.Construct) *awscloudformation.CustomResourceProviderConfig {
	var returns *awscloudformation.CustomResourceProviderConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_Provider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Initialization properties for the `Provider` construct.
// Experimental.
type ProviderProps struct {
	// The AWS Lambda function to invoke for all resource lifecycle operations (CREATE/UPDATE/DELETE).
	//
	// This function is responsible to begin the requested resource operation
	// (CREATE/UPDATE/DELETE) and return any additional properties to add to the
	// event, which will later be passed to `isComplete`. The `PhysicalResourceId`
	// property must be included in the response.
	// Experimental.
	OnEventHandler awslambda.IFunction `json:"onEventHandler"`
	// The AWS Lambda function to invoke in order to determine if the operation is complete.
	//
	// This function will be called immediately after `onEvent` and then
	// periodically based on the configured query interval as long as it returns
	// `false`. If the function still returns `false` and the alloted timeout has
	// passed, the operation will fail.
	// Experimental.
	IsCompleteHandler awslambda.IFunction `json:"isCompleteHandler"`
	// The number of days framework log events are kept in CloudWatch Logs.
	//
	// When
	// updating this property, unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// Time between calls to the `isComplete` handler which determines if the resource has been stabilized.
	//
	// The first `isComplete` will be called immediately after `handler` and then
	// every `queryInterval` seconds, and until `timeout` has been reached or until
	// `isComplete` returns `true`.
	// Experimental.
	QueryInterval awscdk.Duration `json:"queryInterval"`
	// Security groups to attach to the provider functions.
	//
	// Only used if 'vpc' is supplied
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Total timeout for the entire operation.
	//
	// The maximum timeout is 2 hours (yes, it can exceed the AWS Lambda 15 minutes)
	// Experimental.
	TotalTimeout awscdk.Duration `json:"totalTimeout"`
	// The vpc to provision the lambda functions in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Which subnets from the VPC to place the lambda functions in.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// Options for the auto-generation of policies based on the configured SDK calls.
// Experimental.
type SdkCallsPolicyOptions struct {
	// The resources that the calls will have access to.
	//
	// It is best to use specific resource ARN's when possible. However, you can also use `AwsCustomResourcePolicy.ANY_RESOURCE`
	// to allow access to all resources. For example, when `onCreate` is used to create a resource which you don't
	// know the physical name of in advance.
	//
	// Note that will apply to ALL SDK calls.
	// Experimental.
	Resources *[]*string `json:"resources"`
}


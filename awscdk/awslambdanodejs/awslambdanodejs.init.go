package awslambdanodejs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_nodejs.BundlingOptions",
		reflect.TypeOf((*BundlingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_lambda_nodejs.ICommandHooks",
		reflect.TypeOf((*ICommandHooks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "afterBundling", GoMethod: "AfterBundling"},
			_jsii_.MemberMethod{JsiiMethod: "beforeBundling", GoMethod: "BeforeBundling"},
			_jsii_.MemberMethod{JsiiMethod: "beforeInstall", GoMethod: "BeforeInstall"},
		},
		func() interface{} {
			return &jsiiProxy_ICommandHooks{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_lambda_nodejs.LogLevel",
		reflect.TypeOf((*LogLevel)(nil)).Elem(),
		map[string]interface{}{
			"INFO": LogLevel_INFO,
			"WARNING": LogLevel_WARNING,
			"ERROR": LogLevel_ERROR,
			"SILENT": LogLevel_SILENT,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		reflect.TypeOf((*NodejsFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEnvironment", GoMethod: "AddEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSource", GoMethod: "AddEventSource"},
			_jsii_.MemberMethod{JsiiMethod: "addEventSourceMapping", GoMethod: "AddEventSourceMapping"},
			_jsii_.MemberMethod{JsiiMethod: "addLayers", GoMethod: "AddLayers"},
			_jsii_.MemberMethod{JsiiMethod: "addPermission", GoMethod: "AddPermission"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addVersion", GoMethod: "AddVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "canCreatePermissions", GoGetter: "CanCreatePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "configureAsyncInvoke", GoMethod: "ConfigureAsyncInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "currentVersion", GoGetter: "CurrentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "isBoundToVpc", GoGetter: "IsBoundToVpc"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersion", GoGetter: "LatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricErrors", GoMethod: "MetricErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricThrottles", GoMethod: "MetricThrottles"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsNode", GoGetter: "PermissionsNode"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NodejsFunction{}
			_jsii_.InitJsiiProxy(&j.Type__awslambdaFunction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunctionProps",
		reflect.TypeOf((*NodejsFunctionProps)(nil)).Elem(),
	)
}

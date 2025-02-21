package awscloudtrail

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudtrail.AddEventSelectorOptions",
		reflect.TypeOf((*AddEventSelectorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail",
		reflect.TypeOf((*CfnTrail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrSnsTopicArn", GoGetter: "AttrSnsTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchLogsLogGroupArn", GoGetter: "CloudWatchLogsLogGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchLogsRoleArn", GoGetter: "CloudWatchLogsRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableLogFileValidation", GoGetter: "EnableLogFileValidation"},
			_jsii_.MemberProperty{JsiiProperty: "eventSelectors", GoGetter: "EventSelectors"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "includeGlobalServiceEvents", GoGetter: "IncludeGlobalServiceEvents"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "isLogging", GoGetter: "IsLogging"},
			_jsii_.MemberProperty{JsiiProperty: "isMultiRegionTrail", GoGetter: "IsMultiRegionTrail"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketName", GoGetter: "S3BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "s3KeyPrefix", GoGetter: "S3KeyPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopicName", GoGetter: "SnsTopicName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trailName", GoGetter: "TrailName"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrail{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail.DataResourceProperty",
		reflect.TypeOf((*CfnTrail_DataResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudtrail.CfnTrail.EventSelectorProperty",
		reflect.TypeOf((*CfnTrail_EventSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudtrail.CfnTrailProps",
		reflect.TypeOf((*CfnTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudtrail.DataResourceType",
		reflect.TypeOf((*DataResourceType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA_FUNCTION": DataResourceType_LAMBDA_FUNCTION,
			"S3_OBJECT": DataResourceType_S3_OBJECT,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_cloudtrail.ReadWriteType",
		reflect.TypeOf((*ReadWriteType)(nil)).Elem(),
		map[string]interface{}{
			"READ_ONLY": ReadWriteType_READ_ONLY,
			"WRITE_ONLY": ReadWriteType_WRITE_ONLY,
			"ALL": ReadWriteType_ALL,
			"NONE": ReadWriteType_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudtrail.S3EventSelector",
		reflect.TypeOf((*S3EventSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_cloudtrail.Trail",
		reflect.TypeOf((*Trail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventSelector", GoMethod: "AddEventSelector"},
			_jsii_.MemberMethod{JsiiMethod: "addLambdaEventSelector", GoMethod: "AddLambdaEventSelector"},
			_jsii_.MemberMethod{JsiiMethod: "addS3EventSelector", GoMethod: "AddS3EventSelector"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "logAllLambdaDataEvents", GoMethod: "LogAllLambdaDataEvents"},
			_jsii_.MemberMethod{JsiiMethod: "logAllS3DataEvents", GoMethod: "LogAllS3DataEvents"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trailArn", GoGetter: "TrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "trailSnsTopicArn", GoGetter: "TrailSnsTopicArn"},
		},
		func() interface{} {
			j := jsiiProxy_Trail{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_cloudtrail.TrailProps",
		reflect.TypeOf((*TrailProps)(nil)).Elem(),
	)
}

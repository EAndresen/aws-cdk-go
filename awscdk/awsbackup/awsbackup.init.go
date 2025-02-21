package awsbackup

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.BackupPlan",
		reflect.TypeOf((*BackupPlan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRule", GoMethod: "AddRule"},
			_jsii_.MemberMethod{JsiiMethod: "addSelection", GoMethod: "AddSelection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanArn", GoGetter: "BackupPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanId", GoGetter: "BackupPlanId"},
			_jsii_.MemberProperty{JsiiProperty: "backupVault", GoGetter: "BackupVault"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
		},
		func() interface{} {
			j := jsiiProxy_BackupPlan{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBackupPlan)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.BackupPlanProps",
		reflect.TypeOf((*BackupPlanProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		reflect.TypeOf((*BackupPlanRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_BackupPlanRule{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.BackupPlanRuleProps",
		reflect.TypeOf((*BackupPlanRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.BackupResource",
		reflect.TypeOf((*BackupResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "construct", GoGetter: "Construct"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "tagCondition", GoGetter: "TagCondition"},
		},
		func() interface{} {
			return &jsiiProxy_BackupResource{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.BackupSelection",
		reflect.TypeOf((*BackupSelection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanId", GoGetter: "BackupPlanId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "selectionId", GoGetter: "SelectionId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupSelection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.BackupSelectionOptions",
		reflect.TypeOf((*BackupSelectionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.BackupSelectionProps",
		reflect.TypeOf((*BackupSelectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.BackupVault",
		reflect.TypeOf((*BackupVault)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "backupVaultArn", GoGetter: "BackupVaultArn"},
			_jsii_.MemberProperty{JsiiProperty: "backupVaultName", GoGetter: "BackupVaultName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BackupVault{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBackupVault)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_backup.BackupVaultEvents",
		reflect.TypeOf((*BackupVaultEvents)(nil)).Elem(),
		map[string]interface{}{
			"BACKUP_JOB_STARTED": BackupVaultEvents_BACKUP_JOB_STARTED,
			"BACKUP_JOB_COMPLETED": BackupVaultEvents_BACKUP_JOB_COMPLETED,
			"BACKUP_JOB_SUCCESSFUL": BackupVaultEvents_BACKUP_JOB_SUCCESSFUL,
			"BACKUP_JOB_FAILED": BackupVaultEvents_BACKUP_JOB_FAILED,
			"BACKUP_JOB_EXPIRED": BackupVaultEvents_BACKUP_JOB_EXPIRED,
			"RESTORE_JOB_STARTED": BackupVaultEvents_RESTORE_JOB_STARTED,
			"RESTORE_JOB_COMPLETED": BackupVaultEvents_RESTORE_JOB_COMPLETED,
			"RESTORE_JOB_SUCCESSFUL": BackupVaultEvents_RESTORE_JOB_SUCCESSFUL,
			"RESTORE_JOB_FAILED": BackupVaultEvents_RESTORE_JOB_FAILED,
			"COPY_JOB_STARTED": BackupVaultEvents_COPY_JOB_STARTED,
			"COPY_JOB_SUCCESSFUL": BackupVaultEvents_COPY_JOB_SUCCESSFUL,
			"COPY_JOB_FAILED": BackupVaultEvents_COPY_JOB_FAILED,
			"RECOVERY_POINT_MODIFIED": BackupVaultEvents_RECOVERY_POINT_MODIFIED,
			"BACKUP_PLAN_CREATED": BackupVaultEvents_BACKUP_PLAN_CREATED,
			"BACKUP_PLAN_MODIFIED": BackupVaultEvents_BACKUP_PLAN_MODIFIED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.BackupVaultProps",
		reflect.TypeOf((*BackupVaultProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.CfnBackupPlan",
		reflect.TypeOf((*CfnBackupPlan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrBackupPlanArn", GoGetter: "AttrBackupPlanArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrBackupPlanId", GoGetter: "AttrBackupPlanId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersionId", GoGetter: "AttrVersionId"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlan", GoGetter: "BackupPlan"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanTags", GoGetter: "BackupPlanTags"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupPlan{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupPlan.AdvancedBackupSettingResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlan_AdvancedBackupSettingResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupPlan.BackupPlanResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlan_BackupPlanResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupPlan.BackupRuleResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlan_BackupRuleResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupPlan.CopyActionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlan_CopyActionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupPlan.LifecycleResourceTypeProperty",
		reflect.TypeOf((*CfnBackupPlan_LifecycleResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupPlanProps",
		reflect.TypeOf((*CfnBackupPlanProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.CfnBackupSelection",
		reflect.TypeOf((*CfnBackupSelection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrBackupPlanId", GoGetter: "AttrBackupPlanId"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrSelectionId", GoGetter: "AttrSelectionId"},
			_jsii_.MemberProperty{JsiiProperty: "backupPlanId", GoGetter: "BackupPlanId"},
			_jsii_.MemberProperty{JsiiProperty: "backupSelection", GoGetter: "BackupSelection"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupSelection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupSelection.BackupSelectionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupSelection_BackupSelectionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupSelection.ConditionResourceTypeProperty",
		reflect.TypeOf((*CfnBackupSelection_ConditionResourceTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupSelectionProps",
		reflect.TypeOf((*CfnBackupSelectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_backup.CfnBackupVault",
		reflect.TypeOf((*CfnBackupVault)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessPolicy", GoGetter: "AccessPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrBackupVaultArn", GoGetter: "AttrBackupVaultArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrBackupVaultName", GoGetter: "AttrBackupVaultName"},
			_jsii_.MemberProperty{JsiiProperty: "backupVaultName", GoGetter: "BackupVaultName"},
			_jsii_.MemberProperty{JsiiProperty: "backupVaultTags", GoGetter: "BackupVaultTags"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKeyArn", GoGetter: "EncryptionKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifications", GoGetter: "Notifications"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBackupVault{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupVault.NotificationObjectTypeProperty",
		reflect.TypeOf((*CfnBackupVault_NotificationObjectTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.CfnBackupVaultProps",
		reflect.TypeOf((*CfnBackupVaultProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_backup.IBackupPlan",
		reflect.TypeOf((*IBackupPlan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backupPlanId", GoGetter: "BackupPlanId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBackupPlan{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_backup.IBackupVault",
		reflect.TypeOf((*IBackupVault)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backupVaultName", GoGetter: "BackupVaultName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBackupVault{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_backup.TagCondition",
		reflect.TypeOf((*TagCondition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_backup.TagOperation",
		reflect.TypeOf((*TagOperation)(nil)).Elem(),
		map[string]interface{}{
			"STRING_EQUALS": TagOperation_STRING_EQUALS,
			"DUMMY": TagOperation_DUMMY,
		},
	)
}

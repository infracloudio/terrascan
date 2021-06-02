/*
    Copyright (C) 2021 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package store

// CFT equivalent TF resource types
const (
	AwsDocDBCluster                  = "aws_docdb_cluster"
	AwsAPIGatewayRestAPI             = "aws_api_gateway_rest_api"
	AwsLogGroup                      = "aws_cloudwatch_log_group"
	AwsAPIGatewayStage               = "aws_api_gateway_stage"
	AwsAPIGatewayStageMethodSettings = "aws_api_gateway_method_settings"
	AwsAPIGatewayV2Stage             = "aws_apigatewayv2_stage"
	AwsEcsService                    = "aws_ecs_service"
	AwsDynamoDBTable                 = "aws_dynamodb_table"
	AwsDaxCluster                    = "aws_dax_cluster"
	AwsDBInstance                    = "aws_db_instance"
	AwsIamRole                       = "aws_iam_role"
	AwsIamRolePolicy                 = "aws_iam_role_policy"
	AwsIamGroupPolicy                = "aws_iam_group_policy"
	AwsIamPolicy                     = "aws_iam_policy"
	AwsIamAccessKey                  = "aws_iam_access_key"
	AwsIamUser                       = "aws_iam_user"
	AwsIamUserLoginProfile           = "aws_iam_user_login_profile"
	AwsIamUserPolicy                 = "aws_iam_user_policy"
	AwsDBSecurityGroup               = "aws_db_security_group"
	AwsEbsVolume                     = "aws_ebs_volume"
	AwsEfsFileSystem                 = "aws_efs_file_system"
	AwsElastiCacheCluster            = "aws_elasticache_cluster"
	AwsElastiCacheReplicationGroup   = "aws_elasticache_replication_group"
	AwsGuardDutyDetector             = "aws_guardduty_detector"
	AwsMqBroker                      = "aws_mq_broker"
	AwsRedshiftCluster               = "aws_redshift_cluster"
	AwsRdsCluster                    = "aws_rds_cluster"
	AwsRoute53Record                 = "aws_route53_record"
	AwsSecurityGroup                 = "aws_security_group"
	AwsWorkspacesWorkspace           = "aws_workspaces_workspace"
	AwsNeptuneCluster                = "aws_neptune_cluster"
	AwsSecretsManagerSecret          = "aws_secretsmanager_secret"
	AwsEcrRepository                 = "aws_ecr_repository"
	AwsKmsKey                        = "aws_kms_key"
	AwsKinesisStream                 = "aws_kinesis_stream"
	AwsKinesisFirehoseDeliveryStream = "aws_kinesis_firehose_delivery_stream"
	AwsCloudFormationStack           = "aws_cloudformation_stack"
	AwsCloudFrontDistribution        = "aws_cloudfront_distribution"
	AwsCloudTrail                    = "aws_cloudtrail"
	AwsConfigConfigRule              = "aws_config_config_rule"
	AwsConfigConfigurationAggregator = "aws_config_configuration_aggregator"
	AwsLbListener                    = "aws_lb_listener"
	AwsLbTargetGroup                 = "aws_lb_target_group"
	AwsElb                           = "aws_elb"
	AwsElasticsearchDomain           = "aws_elasticsearch_domain"
	AwsElasticsearchDomainPolicy     = "aws_elasticsearch_domain_policy"
	AwsEfsFileSystemPolicy           = "aws_efs_file_system_policy"
	AwsSecretsManagerResourcePolicy  = "aws_secretsmanager_secret_policy"
	AwsEcsTaskDefinition             = "aws_ecs_task_definition"
)

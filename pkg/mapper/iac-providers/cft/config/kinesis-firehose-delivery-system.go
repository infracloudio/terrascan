package config

import (
	"github.com/awslabs/goformation/v4/cloudformation/kinesisfirehose"
)

// KinesisFirehoseDeliveryStreamConfig holds the config for aws_kinesis_firehose_delivery_stream
type KinesisFirehoseDeliveryStreamConfig struct {
	ServerSideEncryption interface{} `json:"server_side_encryption"`
	Config
}

type KinesisFirehoseDeliveryStreamSseConfig struct {
	KeyType string `json:"key_type,omitempty"`
	KeyARN  string `json:"key_arn,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

// GetKinesisFirehoseDeliveryStreamConfig returns aws_kinesis_firehose_delivery_stream
func GetKinesisFirehoseDeliveryStreamConfig(k *kinesisfirehose.DeliveryStream) []AWSResourceConfig {
	cf := KinesisFirehoseDeliveryStreamConfig{
		Config: Config{
			Name: k.DeliveryStreamName,
			Tags: k.Tags,
		},
	}
	sseConfig := KinesisFirehoseDeliveryStreamSseConfig{}
	if k.DeliveryStreamEncryptionConfigurationInput != nil {
		sseConfig.Enabled = true
		sseConfig.KeyType = k.DeliveryStreamEncryptionConfigurationInput.KeyType
		sseConfig.KeyARN = k.DeliveryStreamEncryptionConfigurationInput.KeyARN
	}
	cf.ServerSideEncryption = []KinesisFirehoseDeliveryStreamSseConfig{sseConfig}
	return []AWSResourceConfig{{Resource: cf}}
}

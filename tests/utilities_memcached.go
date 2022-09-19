package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testMemcached(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	name := terraform.Output(t, terraformOptions, "name")
	arn := terraform.Output(t, terraformOptions, "arn")
	SGIDs := terraform.Output(t, terraformOptions, "sg_ids")
	clusterAddress := terraform.Output(t, terraformOptions, "cluster_address")
	configurationEndpoint := terraform.Output(t, terraformOptions, "configuration_endpoint")

	accountID := getAWSAccountID(t)
	region := getAWSRegion(t)

	expectedName := fmt.Sprintf("ex-tf-memcached-%s", variant)
	expectedARN := fmt.Sprintf("arn:aws:elasticache:%s:%s:cluster:%s", region, accountID, expectedName)
	expectedPartialClusterAddress := expectedName
	expectedPartialConfigurationEndpoint := expectedName

	assert.Equal(t, expectedName, name)
	assert.Equal(t, expectedARN, arn)
	assert.NotNil(t, SGIDs)
	assert.Contains(t, clusterAddress, expectedPartialClusterAddress)
	assert.Contains(t, configurationEndpoint, expectedPartialConfigurationEndpoint)
}

// See https://terratest.gruntwork.io/docs/ for full Terratest Documentation.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	//"github.com/stretchr/testify/assert"
)

// Module tests, using default or provided values.
func TestTerraformModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located--relative to tests.
		TerraformDir: "..",

		// By default, tests will read input values from .auto.tfvars
		VarFiles: []string{".auto.tfvars"},

		// Variables can also be declared within this file directly.
		// Vars: map[string]interface{}{
		// 	"someVariable": someVariable,
		// },

		// Environment Variables can also be passed in to tests.
		// EnvVars: map[string]string{
		// 	"AWS_DEFAULT_REGION": awsRegion,
		// },

		// Retry this many times for non-fatal errors.
		MaxRetries: 3,
		// Wait this long between retrying again (in nanoseconds)
		TimeBetweenRetries: 5000000000, // 5 seconds
	}

	// Destroys our Terraform Module after it has been tested.
	defer terraform.Destroy(t, terraformOptions)

	// Performs a `terraform init` and `terraform apply` with to the module.
	terraform.InitAndApply(t, terraformOptions)

	// Test Assertions

	// examples:
	// actualValue := terraform.Output(t, terraformOptions, "some_module_output")
	// expectedValue := "some_module_output"
	//
	// assert.Equal(t, actualValue, expectedValue)
	// assert.Contains(t, someString, someSubString)
	// assert.NotContains(t, someString, someSubString)
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
)

type mocks int

func (mocks) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	return args.Name + "_id", args.Inputs, nil
}

func (mocks) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

func TestInfrastructure(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		_, err := createInfra(ctx)
		assert.NoError(t, err)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks(0)))
	assert.NoError(t, err)
}

func TestInfrastructureIntegration(t *testing.T) {
	cwd, _ := os.Getwd()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:   true,
		Verbose: true,
		Dir:     path.Join(cwd, "..", "s3"),
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			url := fmt.Sprintf("https://%s.s3.us-east-1.amazonaws.com/", stack.Outputs["bucketName"])
			resp, err := http.Get(url)
			assert.NoError(t, err)
			assert.Equal(t, 403, resp.StatusCode)

			for _, resource := range stack.Deployment.Resources {
				if resource.Type == "aws:s3/bucket:Bucket" {
					assert.NotNil(t, resource.Outputs["serverSideEncryptionConfiguration"])
				}
			}
		},
		Config: map[string]string{
			"aws:region": "us-east-1",
		},
	})
}

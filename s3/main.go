package main

import (
	s3 "github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createInfra(ctx *pulumi.Context) (*s3.Bucket, error) {
	// Create an S3 bucket
	bucket, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
		ServerSideEncryptionConfiguration: s3.BucketServerSideEncryptionConfigurationArgs{
			Rule: s3.BucketServerSideEncryptionConfigurationRuleArgs{
				ApplyServerSideEncryptionByDefault: s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
					SseAlgorithm: pulumi.String("aws:kms"),
				},
			},
		},
		Versioning: s3.BucketVersioningArgs{
			Enabled: pulumi.BoolPtr(true),
		},
	})

	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := createInfra(ctx)

		if err != nil {
			return err
		}

		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}

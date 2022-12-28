import { AwsGuard } from "@pulumi/awsguard";

new AwsGuard({ 
    all: "mandatory",
    s3BucketLoggingEnabled: "advisory"
});

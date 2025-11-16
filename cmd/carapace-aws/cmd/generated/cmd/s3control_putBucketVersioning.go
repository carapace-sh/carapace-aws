package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putBucketVersioningCmd = &cobra.Command{
	Use:   "put-bucket-versioning",
	Short: "This operation sets the versioning state for S3 on Outposts buckets only.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putBucketVersioningCmd).Standalone()

	s3control_putBucketVersioningCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 on Outposts bucket.")
	s3control_putBucketVersioningCmd.Flags().String("bucket", "", "The S3 on Outposts bucket to set the versioning state for.")
	s3control_putBucketVersioningCmd.Flags().String("mfa", "", "The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.")
	s3control_putBucketVersioningCmd.Flags().String("versioning-configuration", "", "The root-level tag for the `VersioningConfiguration` parameters.")
	s3control_putBucketVersioningCmd.MarkFlagRequired("account-id")
	s3control_putBucketVersioningCmd.MarkFlagRequired("bucket")
	s3control_putBucketVersioningCmd.MarkFlagRequired("versioning-configuration")
	s3controlCmd.AddCommand(s3control_putBucketVersioningCmd)
}

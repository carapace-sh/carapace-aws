package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putObjectLockConfigurationCmd = &cobra.Command{
	Use:   "put-object-lock-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putObjectLockConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putObjectLockConfigurationCmd).Standalone()

		s3_putObjectLockConfigurationCmd.Flags().String("bucket", "", "The bucket whose Object Lock configuration you want to create or replace.")
		s3_putObjectLockConfigurationCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
		s3_putObjectLockConfigurationCmd.Flags().String("content-md5", "", "The MD5 hash for the request body.")
		s3_putObjectLockConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putObjectLockConfigurationCmd.Flags().String("object-lock-configuration", "", "The Object Lock configuration that you want to apply to the specified bucket.")
		s3_putObjectLockConfigurationCmd.Flags().String("request-payer", "", "")
		s3_putObjectLockConfigurationCmd.Flags().String("token", "", "A token to allow Object Lock to be enabled for an existing bucket.")
		s3_putObjectLockConfigurationCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_putObjectLockConfigurationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectLockConfigurationCmd = &cobra.Command{
	Use:   "get-object-lock-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectLockConfigurationCmd).Standalone()

	s3_getObjectLockConfigurationCmd.Flags().String("bucket", "", "The bucket whose Object Lock configuration you want to retrieve.")
	s3_getObjectLockConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getObjectLockConfigurationCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getObjectLockConfigurationCmd)
}

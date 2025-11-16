package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectAclCmd = &cobra.Command{
	Use:   "get-object-acl",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getObjectAclCmd).Standalone()

		s3_getObjectAclCmd.Flags().String("bucket", "", "The bucket name that contains the object for which to get the ACL information.")
		s3_getObjectAclCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getObjectAclCmd.Flags().String("key", "", "The key of the object for which to get the ACL information.")
		s3_getObjectAclCmd.Flags().String("request-payer", "", "")
		s3_getObjectAclCmd.Flags().String("version-id", "", "Version ID used to reference a specific version of the object.")
		s3_getObjectAclCmd.MarkFlagRequired("bucket")
		s3_getObjectAclCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_getObjectAclCmd)
}

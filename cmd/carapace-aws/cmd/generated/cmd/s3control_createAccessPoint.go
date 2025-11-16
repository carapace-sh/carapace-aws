package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createAccessPointCmd = &cobra.Command{
	Use:   "create-access-point",
	Short: "Creates an access point and associates it to a specified bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createAccessPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_createAccessPointCmd).Standalone()

		s3control_createAccessPointCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the account that owns the specified access point.")
		s3control_createAccessPointCmd.Flags().String("bucket", "", "The name of the bucket that you want to associate this access point with.")
		s3control_createAccessPointCmd.Flags().String("bucket-account-id", "", "The Amazon Web Services account ID associated with the S3 bucket associated with this access point.")
		s3control_createAccessPointCmd.Flags().String("name", "", "The name you want to assign to this access point.")
		s3control_createAccessPointCmd.Flags().String("public-access-block-configuration", "", "The `PublicAccessBlock` configuration that you want to apply to the access point.")
		s3control_createAccessPointCmd.Flags().String("scope", "", "For directory buckets, you can filter access control to specific prefixes, API operations, or a combination of both.")
		s3control_createAccessPointCmd.Flags().String("tags", "", "An array of tags that you can apply to an access point.")
		s3control_createAccessPointCmd.Flags().String("vpc-configuration", "", "If you include this field, Amazon S3 restricts access to this access point to requests from the specified virtual private cloud (VPC).")
		s3control_createAccessPointCmd.MarkFlagRequired("account-id")
		s3control_createAccessPointCmd.MarkFlagRequired("bucket")
		s3control_createAccessPointCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_createAccessPointCmd)
}

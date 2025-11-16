package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessGrantsInstanceResourcePolicyCmd = &cobra.Command{
	Use:   "get-access-grants-instance-resource-policy",
	Short: "Returns the resource policy of the S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessGrantsInstanceResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessGrantsInstanceResourcePolicyCmd).Standalone()

		s3control_getAccessGrantsInstanceResourcePolicyCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_getAccessGrantsInstanceResourcePolicyCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_getAccessGrantsInstanceResourcePolicyCmd)
}

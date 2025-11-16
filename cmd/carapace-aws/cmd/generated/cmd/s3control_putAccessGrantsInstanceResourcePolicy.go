package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putAccessGrantsInstanceResourcePolicyCmd = &cobra.Command{
	Use:   "put-access-grants-instance-resource-policy",
	Short: "Updates the resource policy of the S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putAccessGrantsInstanceResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putAccessGrantsInstanceResourcePolicyCmd).Standalone()

		s3control_putAccessGrantsInstanceResourcePolicyCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_putAccessGrantsInstanceResourcePolicyCmd.Flags().String("organization", "", "The Organization of the resource policy of the S3 Access Grants instance.")
		s3control_putAccessGrantsInstanceResourcePolicyCmd.Flags().String("policy", "", "The resource policy of the S3 Access Grants instance that you are updating.")
		s3control_putAccessGrantsInstanceResourcePolicyCmd.MarkFlagRequired("account-id")
		s3control_putAccessGrantsInstanceResourcePolicyCmd.MarkFlagRequired("policy")
	})
	s3controlCmd.AddCommand(s3control_putAccessGrantsInstanceResourcePolicyCmd)
}

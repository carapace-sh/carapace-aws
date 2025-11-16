package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_dissociateAccessGrantsIdentityCenterCmd = &cobra.Command{
	Use:   "dissociate-access-grants-identity-center",
	Short: "Dissociates the Amazon Web Services IAM Identity Center instance from the S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_dissociateAccessGrantsIdentityCenterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_dissociateAccessGrantsIdentityCenterCmd).Standalone()

		s3control_dissociateAccessGrantsIdentityCenterCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_dissociateAccessGrantsIdentityCenterCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_dissociateAccessGrantsIdentityCenterCmd)
}

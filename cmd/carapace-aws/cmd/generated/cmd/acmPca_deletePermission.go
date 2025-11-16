package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_deletePermissionCmd = &cobra.Command{
	Use:   "delete-permission",
	Short: "Revokes permissions on a private CA granted to the Certificate Manager (ACM) service principal (acm.amazonaws.com).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_deletePermissionCmd).Standalone()

	acmPca_deletePermissionCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Number (ARN) of the private CA that issued the permissions.")
	acmPca_deletePermissionCmd.Flags().String("principal", "", "The Amazon Web Services service or identity that will have its CA permissions revoked.")
	acmPca_deletePermissionCmd.Flags().String("source-account", "", "The Amazon Web Services account that calls this action.")
	acmPca_deletePermissionCmd.MarkFlagRequired("certificate-authority-arn")
	acmPca_deletePermissionCmd.MarkFlagRequired("principal")
	acmPcaCmd.AddCommand(acmPca_deletePermissionCmd)
}

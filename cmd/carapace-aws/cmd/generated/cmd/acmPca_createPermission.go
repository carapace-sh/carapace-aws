package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_createPermissionCmd = &cobra.Command{
	Use:   "create-permission",
	Short: "Grants one or more permissions on a private CA to the Certificate Manager (ACM) service principal (`acm.amazonaws.com`).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_createPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_createPermissionCmd).Standalone()

		acmPca_createPermissionCmd.Flags().String("actions", "", "The actions that the specified Amazon Web Services service principal can use.")
		acmPca_createPermissionCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) of the CA that grants the permissions.")
		acmPca_createPermissionCmd.Flags().String("principal", "", "The Amazon Web Services service or identity that receives the permission.")
		acmPca_createPermissionCmd.Flags().String("source-account", "", "The ID of the calling account.")
		acmPca_createPermissionCmd.MarkFlagRequired("actions")
		acmPca_createPermissionCmd.MarkFlagRequired("certificate-authority-arn")
		acmPca_createPermissionCmd.MarkFlagRequired("principal")
	})
	acmPcaCmd.AddCommand(acmPca_createPermissionCmd)
}

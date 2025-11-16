package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_listPermissionsCmd = &cobra.Command{
	Use:   "list-permissions",
	Short: "List all permissions on a private CA, if any, granted to the Certificate Manager (ACM) service principal (acm.amazonaws.com).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_listPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_listPermissionsCmd).Standalone()

		acmPca_listPermissionsCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Number (ARN) of the private CA to inspect.")
		acmPca_listPermissionsCmd.Flags().String("max-results", "", "When paginating results, use this parameter to specify the maximum number of items to return in the response.")
		acmPca_listPermissionsCmd.Flags().String("next-token", "", "When paginating results, use this parameter in a subsequent request after you receive a response with truncated results.")
		acmPca_listPermissionsCmd.MarkFlagRequired("certificate-authority-arn")
	})
	acmPcaCmd.AddCommand(acmPca_listPermissionsCmd)
}

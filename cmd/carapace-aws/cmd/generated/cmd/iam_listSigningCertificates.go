package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listSigningCertificatesCmd = &cobra.Command{
	Use:   "list-signing-certificates",
	Short: "Returns information about the signing certificates associated with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listSigningCertificatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listSigningCertificatesCmd).Standalone()

		iam_listSigningCertificatesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listSigningCertificatesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listSigningCertificatesCmd.Flags().String("user-name", "", "The name of the IAM user whose signing certificates you want to examine.")
	})
	iamCmd.AddCommand(iam_listSigningCertificatesCmd)
}

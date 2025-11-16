package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateSigningCertificateCmd = &cobra.Command{
	Use:   "update-signing-certificate",
	Short: "Changes the status of the specified user signing certificate from active to disabled, or vice versa.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateSigningCertificateCmd).Standalone()

	iam_updateSigningCertificateCmd.Flags().String("certificate-id", "", "The ID of the signing certificate you want to update.")
	iam_updateSigningCertificateCmd.Flags().String("status", "", "The status you want to assign to the certificate.")
	iam_updateSigningCertificateCmd.Flags().String("user-name", "", "The name of the IAM user the signing certificate belongs to.")
	iam_updateSigningCertificateCmd.MarkFlagRequired("certificate-id")
	iam_updateSigningCertificateCmd.MarkFlagRequired("status")
	iamCmd.AddCommand(iam_updateSigningCertificateCmd)
}

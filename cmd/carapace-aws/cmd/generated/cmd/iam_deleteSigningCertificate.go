package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteSigningCertificateCmd = &cobra.Command{
	Use:   "delete-signing-certificate",
	Short: "Deletes a signing certificate associated with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteSigningCertificateCmd).Standalone()

	iam_deleteSigningCertificateCmd.Flags().String("certificate-id", "", "The ID of the signing certificate to delete.")
	iam_deleteSigningCertificateCmd.Flags().String("user-name", "", "The name of the user the signing certificate belongs to.")
	iam_deleteSigningCertificateCmd.MarkFlagRequired("certificate-id")
	iamCmd.AddCommand(iam_deleteSigningCertificateCmd)
}

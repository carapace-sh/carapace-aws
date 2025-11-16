package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateServerCertificateCmd = &cobra.Command{
	Use:   "update-server-certificate",
	Short: "Updates the name and/or the path of the specified server certificate stored in IAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateServerCertificateCmd).Standalone()

	iam_updateServerCertificateCmd.Flags().String("new-path", "", "The new path for the server certificate.")
	iam_updateServerCertificateCmd.Flags().String("new-server-certificate-name", "", "The new name for the server certificate.")
	iam_updateServerCertificateCmd.Flags().String("server-certificate-name", "", "The name of the server certificate that you want to update.")
	iam_updateServerCertificateCmd.MarkFlagRequired("server-certificate-name")
	iamCmd.AddCommand(iam_updateServerCertificateCmd)
}

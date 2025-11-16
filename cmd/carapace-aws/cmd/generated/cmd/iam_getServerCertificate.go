package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getServerCertificateCmd = &cobra.Command{
	Use:   "get-server-certificate",
	Short: "Retrieves information about the specified server certificate stored in IAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getServerCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getServerCertificateCmd).Standalone()

		iam_getServerCertificateCmd.Flags().String("server-certificate-name", "", "The name of the server certificate you want to retrieve information about.")
		iam_getServerCertificateCmd.MarkFlagRequired("server-certificate-name")
	})
	iamCmd.AddCommand(iam_getServerCertificateCmd)
}

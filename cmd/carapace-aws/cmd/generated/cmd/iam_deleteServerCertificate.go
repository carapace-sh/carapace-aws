package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteServerCertificateCmd = &cobra.Command{
	Use:   "delete-server-certificate",
	Short: "Deletes the specified server certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteServerCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deleteServerCertificateCmd).Standalone()

		iam_deleteServerCertificateCmd.Flags().String("server-certificate-name", "", "The name of the server certificate you want to delete.")
		iam_deleteServerCertificateCmd.MarkFlagRequired("server-certificate-name")
	})
	iamCmd.AddCommand(iam_deleteServerCertificateCmd)
}

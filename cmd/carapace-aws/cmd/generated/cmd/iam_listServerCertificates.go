package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listServerCertificatesCmd = &cobra.Command{
	Use:   "list-server-certificates",
	Short: "Lists the server certificates stored in IAM that have the specified path prefix.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listServerCertificatesCmd).Standalone()

	iam_listServerCertificatesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listServerCertificatesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listServerCertificatesCmd.Flags().String("path-prefix", "", "The path prefix for filtering the results.")
	iamCmd.AddCommand(iam_listServerCertificatesCmd)
}

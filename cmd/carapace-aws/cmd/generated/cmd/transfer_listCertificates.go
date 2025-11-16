package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listCertificatesCmd = &cobra.Command{
	Use:   "list-certificates",
	Short: "Returns a list of the current certificates that have been imported into Transfer Family.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listCertificatesCmd).Standalone()

	transfer_listCertificatesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	transfer_listCertificatesCmd.Flags().String("next-token", "", "When you can get additional results from the `ListCertificates` call, a `NextToken` parameter is returned in the output.")
	transferCmd.AddCommand(transfer_listCertificatesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listAgreementsCmd = &cobra.Command{
	Use:   "list-agreements",
	Short: "Returns a list of the agreements for the server that's identified by the `ServerId` that you supply.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listAgreementsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_listAgreementsCmd).Standalone()

		transfer_listAgreementsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		transfer_listAgreementsCmd.Flags().String("next-token", "", "When you can get additional results from the `ListAgreements` call, a `NextToken` parameter is returned in the output.")
		transfer_listAgreementsCmd.Flags().String("server-id", "", "The identifier of the server for which you want a list of agreements.")
		transfer_listAgreementsCmd.MarkFlagRequired("server-id")
	})
	transferCmd.AddCommand(transfer_listAgreementsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listPartnerAccountsCmd = &cobra.Command{
	Use:   "list-partner-accounts",
	Short: "Lists the partner accounts associated with your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listPartnerAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listPartnerAccountsCmd).Standalone()

		iotwireless_listPartnerAccountsCmd.Flags().String("max-results", "", "The maximum number of results to return in this operation.")
		iotwireless_listPartnerAccountsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	})
	iotwirelessCmd.AddCommand(iotwireless_listPartnerAccountsCmd)
}

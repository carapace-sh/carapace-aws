package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listUsageTotalsCmd = &cobra.Command{
	Use:   "list-usage-totals",
	Short: "Lists the Amazon Inspector usage totals over the last 30 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listUsageTotalsCmd).Standalone()

	inspector2_listUsageTotalsCmd.Flags().String("account-ids", "", "The Amazon Web Services account IDs to retrieve usage totals for.")
	inspector2_listUsageTotalsCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
	inspector2_listUsageTotalsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	inspector2Cmd.AddCommand(inspector2_listUsageTotalsCmd)
}

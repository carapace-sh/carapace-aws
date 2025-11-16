package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getUpgradeHistoryCmd = &cobra.Command{
	Use:   "get-upgrade-history",
	Short: "Retrieves the complete history of the last 10 upgrades performed on an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getUpgradeHistoryCmd).Standalone()

	opensearch_getUpgradeHistoryCmd.Flags().String("domain-name", "", "The name of an existing domain.")
	opensearch_getUpgradeHistoryCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_getUpgradeHistoryCmd.Flags().String("next-token", "", "If your initial `GetUpgradeHistory` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `GetUpgradeHistory` operations, which returns results in the next page.")
	opensearch_getUpgradeHistoryCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_getUpgradeHistoryCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_getUpgradeHistoryCmd = &cobra.Command{
	Use:   "get-upgrade-history",
	Short: "Retrieves the complete history of the last 10 upgrades that were performed on the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_getUpgradeHistoryCmd).Standalone()

	es_getUpgradeHistoryCmd.Flags().String("domain-name", "", "")
	es_getUpgradeHistoryCmd.Flags().String("max-results", "", "")
	es_getUpgradeHistoryCmd.Flags().String("next-token", "", "")
	es_getUpgradeHistoryCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_getUpgradeHistoryCmd)
}

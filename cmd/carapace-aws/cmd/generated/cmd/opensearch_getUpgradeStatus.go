package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getUpgradeStatusCmd = &cobra.Command{
	Use:   "get-upgrade-status",
	Short: "Returns the most recent status of the last upgrade or upgrade eligibility check performed on an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getUpgradeStatusCmd).Standalone()

	opensearch_getUpgradeStatusCmd.Flags().String("domain-name", "", "The domain of the domain to get upgrade status information for.")
	opensearch_getUpgradeStatusCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_getUpgradeStatusCmd)
}

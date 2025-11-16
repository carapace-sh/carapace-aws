package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_getPoliciesStatsCmd = &cobra.Command{
	Use:   "get-policies-stats",
	Short: "Returns statistical information about your OpenSearch Serverless access policies, security configurations, and security policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_getPoliciesStatsCmd).Standalone()

	opensearchserverlessCmd.AddCommand(opensearchserverless_getPoliciesStatsCmd)
}

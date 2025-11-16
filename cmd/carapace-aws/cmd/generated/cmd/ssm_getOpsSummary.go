package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getOpsSummaryCmd = &cobra.Command{
	Use:   "get-ops-summary",
	Short: "View a summary of operations metadata (OpsData) based on specified filters and aggregators.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getOpsSummaryCmd).Standalone()

	ssm_getOpsSummaryCmd.Flags().String("aggregators", "", "Optional aggregators that return counts of OpsData based on one or more expressions.")
	ssm_getOpsSummaryCmd.Flags().String("filters", "", "Optional filters used to scope down the returned OpsData.")
	ssm_getOpsSummaryCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_getOpsSummaryCmd.Flags().String("next-token", "", "A token to start the list.")
	ssm_getOpsSummaryCmd.Flags().String("result-attributes", "", "The OpsData data type to return.")
	ssm_getOpsSummaryCmd.Flags().String("sync-name", "", "Specify the name of a resource data sync to get.")
	ssmCmd.AddCommand(ssm_getOpsSummaryCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getResourcesStatisticsV2Cmd = &cobra.Command{
	Use:   "get-resources-statistics-v2",
	Short: "Retrieves statistical information about Amazon Web Services resources and their associated security findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getResourcesStatisticsV2Cmd).Standalone()

	securityhub_getResourcesStatisticsV2Cmd.Flags().String("group-by-rules", "", "How resource statistics should be aggregated and organized in the response.")
	securityhub_getResourcesStatisticsV2Cmd.Flags().String("max-statistic-results", "", "The maximum number of results to be returned.")
	securityhub_getResourcesStatisticsV2Cmd.Flags().String("sort-order", "", "Sorts aggregated statistics.")
	securityhub_getResourcesStatisticsV2Cmd.MarkFlagRequired("group-by-rules")
	securityhubCmd.AddCommand(securityhub_getResourcesStatisticsV2Cmd)
}

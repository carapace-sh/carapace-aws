package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listInsightsCmd = &cobra.Command{
	Use:   "list-insights",
	Short: "Returns a list of insights in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listInsightsCmd).Standalone()

	devopsGuru_listInsightsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	devopsGuru_listInsightsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuru_listInsightsCmd.Flags().String("status-filter", "", "A filter used to filter the returned insights by their status.")
	devopsGuru_listInsightsCmd.MarkFlagRequired("status-filter")
	devopsGuruCmd.AddCommand(devopsGuru_listInsightsCmd)
}

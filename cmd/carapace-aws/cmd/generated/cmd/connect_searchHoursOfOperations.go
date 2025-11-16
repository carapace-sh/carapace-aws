package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchHoursOfOperationsCmd = &cobra.Command{
	Use:   "search-hours-of-operations",
	Short: "Searches the hours of operation in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchHoursOfOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchHoursOfOperationsCmd).Standalone()

		connect_searchHoursOfOperationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchHoursOfOperationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchHoursOfOperationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchHoursOfOperationsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return hours of operations.")
		connect_searchHoursOfOperationsCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchHoursOfOperationsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchHoursOfOperationsCmd)
}

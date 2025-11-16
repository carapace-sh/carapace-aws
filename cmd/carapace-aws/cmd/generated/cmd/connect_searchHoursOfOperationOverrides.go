package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchHoursOfOperationOverridesCmd = &cobra.Command{
	Use:   "search-hours-of-operation-overrides",
	Short: "Searches the hours of operation overrides.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchHoursOfOperationOverridesCmd).Standalone()

	connect_searchHoursOfOperationOverridesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchHoursOfOperationOverridesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchHoursOfOperationOverridesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchHoursOfOperationOverridesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return hours of operations overrides.")
	connect_searchHoursOfOperationOverridesCmd.Flags().String("search-filter", "", "")
	connect_searchHoursOfOperationOverridesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchHoursOfOperationOverridesCmd)
}

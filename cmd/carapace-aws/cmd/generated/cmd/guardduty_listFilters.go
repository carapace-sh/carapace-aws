package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listFiltersCmd = &cobra.Command{
	Use:   "list-filters",
	Short: "Returns a paginated list of the current filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listFiltersCmd).Standalone()

	guardduty_listFiltersCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with the filter.")
	guardduty_listFiltersCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items that you want in the response.")
	guardduty_listFiltersCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	guardduty_listFiltersCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_listFiltersCmd)
}

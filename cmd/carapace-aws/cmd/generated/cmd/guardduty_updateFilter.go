package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateFilterCmd = &cobra.Command{
	Use:   "update-filter",
	Short: "Updates the filter specified by the filter name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateFilterCmd).Standalone()

	guardduty_updateFilterCmd.Flags().String("action", "", "Specifies the action that is to be applied to the findings that match the filter.")
	guardduty_updateFilterCmd.Flags().String("description", "", "The description of the filter.")
	guardduty_updateFilterCmd.Flags().String("detector-id", "", "The unique ID of the detector that specifies the GuardDuty service where you want to update a filter.")
	guardduty_updateFilterCmd.Flags().String("filter-name", "", "The name of the filter.")
	guardduty_updateFilterCmd.Flags().String("finding-criteria", "", "Represents the criteria to be used in the filter for querying findings.")
	guardduty_updateFilterCmd.Flags().String("rank", "", "Specifies the position of the filter in the list of current filters.")
	guardduty_updateFilterCmd.MarkFlagRequired("detector-id")
	guardduty_updateFilterCmd.MarkFlagRequired("filter-name")
	guarddutyCmd.AddCommand(guardduty_updateFilterCmd)
}

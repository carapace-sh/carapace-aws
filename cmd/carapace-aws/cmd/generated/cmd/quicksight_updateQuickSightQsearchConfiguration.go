package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateQuickSightQsearchConfigurationCmd = &cobra.Command{
	Use:   "update-quick-sight-qsearch-configuration",
	Short: "Updates the state of a Quick Sight Q Search configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateQuickSightQsearchConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateQuickSightQsearchConfigurationCmd).Standalone()

		quicksight_updateQuickSightQsearchConfigurationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the Quick Sight Q Search configuration that you want to update.")
		quicksight_updateQuickSightQsearchConfigurationCmd.Flags().String("qsearch-status", "", "The status of the Quick Sight Q Search configuration that the user wants to update.")
		quicksight_updateQuickSightQsearchConfigurationCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateQuickSightQsearchConfigurationCmd.MarkFlagRequired("qsearch-status")
	})
	quicksightCmd.AddCommand(quicksight_updateQuickSightQsearchConfigurationCmd)
}

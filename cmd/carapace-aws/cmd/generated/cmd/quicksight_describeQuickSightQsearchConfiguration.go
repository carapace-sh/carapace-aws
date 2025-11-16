package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeQuickSightQsearchConfigurationCmd = &cobra.Command{
	Use:   "describe-quick-sight-qsearch-configuration",
	Short: "Describes the state of a Quick Sight Q Search configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeQuickSightQsearchConfigurationCmd).Standalone()

	quicksight_describeQuickSightQsearchConfigurationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the Quick Sight Q Search configuration that the user wants described.")
	quicksight_describeQuickSightQsearchConfigurationCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeQuickSightQsearchConfigurationCmd)
}

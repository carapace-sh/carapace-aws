package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateSpicecapacityConfigurationCmd = &cobra.Command{
	Use:   "update-spicecapacity-configuration",
	Short: "Updates the SPICE capacity configuration for a Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateSpicecapacityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateSpicecapacityConfigurationCmd).Standalone()

		quicksight_updateSpicecapacityConfigurationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the SPICE configuration that you want to update.")
		quicksight_updateSpicecapacityConfigurationCmd.Flags().String("purchase-mode", "", "Determines how SPICE capacity can be purchased.")
		quicksight_updateSpicecapacityConfigurationCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateSpicecapacityConfigurationCmd.MarkFlagRequired("purchase-mode")
	})
	quicksightCmd.AddCommand(quicksight_updateSpicecapacityConfigurationCmd)
}

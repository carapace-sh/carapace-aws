package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateBrandCmd = &cobra.Command{
	Use:   "update-brand",
	Short: "Updates a brand.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateBrandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateBrandCmd).Standalone()

		quicksight_updateBrandCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand.")
		quicksight_updateBrandCmd.Flags().String("brand-definition", "", "The definition of the brand.")
		quicksight_updateBrandCmd.Flags().String("brand-id", "", "The ID of the Quick Suite brand.")
		quicksight_updateBrandCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateBrandCmd.MarkFlagRequired("brand-id")
	})
	quicksightCmd.AddCommand(quicksight_updateBrandCmd)
}

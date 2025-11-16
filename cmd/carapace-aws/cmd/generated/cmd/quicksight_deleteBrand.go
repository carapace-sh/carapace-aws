package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteBrandCmd = &cobra.Command{
	Use:   "delete-brand",
	Short: "This API permanently deletes the specified Quick Sight brand.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteBrandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteBrandCmd).Standalone()

		quicksight_deleteBrandCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand.")
		quicksight_deleteBrandCmd.Flags().String("brand-id", "", "The ID of the Quick Suite brand.")
		quicksight_deleteBrandCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteBrandCmd.MarkFlagRequired("brand-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteBrandCmd)
}

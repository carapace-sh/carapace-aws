package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteDataSetRefreshPropertiesCmd = &cobra.Command{
	Use:   "delete-data-set-refresh-properties",
	Short: "Deletes the dataset refresh properties of the dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteDataSetRefreshPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteDataSetRefreshPropertiesCmd).Standalone()

		quicksight_deleteDataSetRefreshPropertiesCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_deleteDataSetRefreshPropertiesCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
		quicksight_deleteDataSetRefreshPropertiesCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteDataSetRefreshPropertiesCmd.MarkFlagRequired("data-set-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteDataSetRefreshPropertiesCmd)
}

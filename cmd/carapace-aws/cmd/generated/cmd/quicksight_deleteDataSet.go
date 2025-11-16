package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteDataSetCmd = &cobra.Command{
	Use:   "delete-data-set",
	Short: "Deletes a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteDataSetCmd).Standalone()

		quicksight_deleteDataSetCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_deleteDataSetCmd.Flags().String("data-set-id", "", "The ID for the dataset that you want to delete.")
		quicksight_deleteDataSetCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteDataSetCmd.MarkFlagRequired("data-set-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteDataSetCmd)
}

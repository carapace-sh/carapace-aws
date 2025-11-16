package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Deletes the data source permanently.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteDataSourceCmd).Standalone()

		quicksight_deleteDataSourceCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_deleteDataSourceCmd.Flags().String("data-source-id", "", "The ID of the data source.")
		quicksight_deleteDataSourceCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteDataSourceCmd.MarkFlagRequired("data-source-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteDataSourceCmd)
}

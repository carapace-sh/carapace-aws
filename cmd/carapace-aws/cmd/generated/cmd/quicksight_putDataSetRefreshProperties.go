package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_putDataSetRefreshPropertiesCmd = &cobra.Command{
	Use:   "put-data-set-refresh-properties",
	Short: "Creates or updates the dataset refresh properties for the dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_putDataSetRefreshPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_putDataSetRefreshPropertiesCmd).Standalone()

		quicksight_putDataSetRefreshPropertiesCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_putDataSetRefreshPropertiesCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
		quicksight_putDataSetRefreshPropertiesCmd.Flags().String("data-set-refresh-properties", "", "The dataset refresh properties.")
		quicksight_putDataSetRefreshPropertiesCmd.MarkFlagRequired("aws-account-id")
		quicksight_putDataSetRefreshPropertiesCmd.MarkFlagRequired("data-set-id")
		quicksight_putDataSetRefreshPropertiesCmd.MarkFlagRequired("data-set-refresh-properties")
	})
	quicksightCmd.AddCommand(quicksight_putDataSetRefreshPropertiesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDataSetRefreshPropertiesCmd = &cobra.Command{
	Use:   "describe-data-set-refresh-properties",
	Short: "Describes the refresh properties of a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDataSetRefreshPropertiesCmd).Standalone()

	quicksight_describeDataSetRefreshPropertiesCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_describeDataSetRefreshPropertiesCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
	quicksight_describeDataSetRefreshPropertiesCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeDataSetRefreshPropertiesCmd.MarkFlagRequired("data-set-id")
	quicksightCmd.AddCommand(quicksight_describeDataSetRefreshPropertiesCmd)
}

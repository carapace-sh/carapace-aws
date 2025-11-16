package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDataSetCmd = &cobra.Command{
	Use:   "describe-data-set",
	Short: "Describes a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeDataSetCmd).Standalone()

		quicksight_describeDataSetCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_describeDataSetCmd.Flags().String("data-set-id", "", "The ID for the dataset that you want to describe.")
		quicksight_describeDataSetCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeDataSetCmd.MarkFlagRequired("data-set-id")
	})
	quicksightCmd.AddCommand(quicksight_describeDataSetCmd)
}

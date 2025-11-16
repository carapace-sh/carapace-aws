package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDataSourceCmd = &cobra.Command{
	Use:   "describe-data-source",
	Short: "Describes a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeDataSourceCmd).Standalone()

		quicksight_describeDataSourceCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_describeDataSourceCmd.Flags().String("data-source-id", "", "The ID of the data source.")
		quicksight_describeDataSourceCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeDataSourceCmd.MarkFlagRequired("data-source-id")
	})
	quicksightCmd.AddCommand(quicksight_describeDataSourceCmd)
}

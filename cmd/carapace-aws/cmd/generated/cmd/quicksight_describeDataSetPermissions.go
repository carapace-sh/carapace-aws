package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDataSetPermissionsCmd = &cobra.Command{
	Use:   "describe-data-set-permissions",
	Short: "Describes the permissions on a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDataSetPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeDataSetPermissionsCmd).Standalone()

		quicksight_describeDataSetPermissionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_describeDataSetPermissionsCmd.Flags().String("data-set-id", "", "The ID for the dataset that you want to describe.")
		quicksight_describeDataSetPermissionsCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeDataSetPermissionsCmd.MarkFlagRequired("data-set-id")
	})
	quicksightCmd.AddCommand(quicksight_describeDataSetPermissionsCmd)
}

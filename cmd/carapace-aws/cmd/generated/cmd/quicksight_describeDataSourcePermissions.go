package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDataSourcePermissionsCmd = &cobra.Command{
	Use:   "describe-data-source-permissions",
	Short: "Describes the resource permissions for a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDataSourcePermissionsCmd).Standalone()

	quicksight_describeDataSourcePermissionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_describeDataSourcePermissionsCmd.Flags().String("data-source-id", "", "The ID of the data source.")
	quicksight_describeDataSourcePermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeDataSourcePermissionsCmd.MarkFlagRequired("data-source-id")
	quicksightCmd.AddCommand(quicksight_describeDataSourcePermissionsCmd)
}

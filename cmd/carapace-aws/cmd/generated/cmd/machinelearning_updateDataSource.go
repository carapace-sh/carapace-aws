package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates the `DataSourceName` of a `DataSource`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_updateDataSourceCmd).Standalone()

	machinelearning_updateDataSourceCmd.Flags().String("data-source-id", "", "The ID assigned to the `DataSource` during creation.")
	machinelearning_updateDataSourceCmd.Flags().String("data-source-name", "", "A new user-supplied name or description of the `DataSource` that will replace the current description.")
	machinelearning_updateDataSourceCmd.MarkFlagRequired("data-source-id")
	machinelearning_updateDataSourceCmd.MarkFlagRequired("data-source-name")
	machinelearningCmd.AddCommand(machinelearning_updateDataSourceCmd)
}

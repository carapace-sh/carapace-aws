package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_getDataSourceCmd = &cobra.Command{
	Use:   "get-data-source",
	Short: "Returns a `DataSource` that includes metadata and data file information, as well as the current status of the `DataSource`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_getDataSourceCmd).Standalone()

	machinelearning_getDataSourceCmd.Flags().String("data-source-id", "", "The ID assigned to the `DataSource` at creation.")
	machinelearning_getDataSourceCmd.Flags().String("verbose", "", "Specifies whether the `GetDataSource` operation should return `DataSourceSchema`.")
	machinelearning_getDataSourceCmd.MarkFlagRequired("data-source-id")
	machinelearningCmd.AddCommand(machinelearning_getDataSourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Assigns the DELETED status to a `DataSource`, rendering it unusable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_deleteDataSourceCmd).Standalone()

	machinelearning_deleteDataSourceCmd.Flags().String("data-source-id", "", "A user-supplied ID that uniquely identifies the `DataSource`.")
	machinelearning_deleteDataSourceCmd.MarkFlagRequired("data-source-id")
	machinelearningCmd.AddCommand(machinelearning_deleteDataSourceCmd)
}

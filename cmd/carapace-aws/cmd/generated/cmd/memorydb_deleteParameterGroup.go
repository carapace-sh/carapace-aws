package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_deleteParameterGroupCmd = &cobra.Command{
	Use:   "delete-parameter-group",
	Short: "Deletes the specified parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_deleteParameterGroupCmd).Standalone()

	memorydb_deleteParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to delete.")
	memorydb_deleteParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	memorydbCmd.AddCommand(memorydb_deleteParameterGroupCmd)
}

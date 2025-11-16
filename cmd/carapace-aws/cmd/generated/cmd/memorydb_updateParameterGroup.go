package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_updateParameterGroupCmd = &cobra.Command{
	Use:   "update-parameter-group",
	Short: "Updates the parameters of a parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_updateParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_updateParameterGroupCmd).Standalone()

		memorydb_updateParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to update.")
		memorydb_updateParameterGroupCmd.Flags().String("parameter-name-values", "", "An array of parameter names and values for the parameter update.")
		memorydb_updateParameterGroupCmd.MarkFlagRequired("parameter-group-name")
		memorydb_updateParameterGroupCmd.MarkFlagRequired("parameter-name-values")
	})
	memorydbCmd.AddCommand(memorydb_updateParameterGroupCmd)
}

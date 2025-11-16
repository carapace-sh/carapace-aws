package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_resetParameterGroupCmd = &cobra.Command{
	Use:   "reset-parameter-group",
	Short: "Modifies the parameters of a parameter group to the engine or system default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_resetParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_resetParameterGroupCmd).Standalone()

		memorydb_resetParameterGroupCmd.Flags().Bool("all-parameters", false, "If true, all parameters in the parameter group are reset to their default values.")
		memorydb_resetParameterGroupCmd.Flags().Bool("no-all-parameters", false, "If true, all parameters in the parameter group are reset to their default values.")
		memorydb_resetParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to reset.")
		memorydb_resetParameterGroupCmd.Flags().String("parameter-names", "", "An array of parameter names to reset to their default values.")
		memorydb_resetParameterGroupCmd.Flag("no-all-parameters").Hidden = true
		memorydb_resetParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	})
	memorydbCmd.AddCommand(memorydb_resetParameterGroupCmd)
}

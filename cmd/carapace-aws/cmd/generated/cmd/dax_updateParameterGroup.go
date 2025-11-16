package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_updateParameterGroupCmd = &cobra.Command{
	Use:   "update-parameter-group",
	Short: "Modifies the parameters of a parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_updateParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_updateParameterGroupCmd).Standalone()

		dax_updateParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group.")
		dax_updateParameterGroupCmd.Flags().String("parameter-name-values", "", "An array of name-value pairs for the parameters in the group.")
		dax_updateParameterGroupCmd.MarkFlagRequired("parameter-group-name")
		dax_updateParameterGroupCmd.MarkFlagRequired("parameter-name-values")
	})
	daxCmd.AddCommand(dax_updateParameterGroupCmd)
}

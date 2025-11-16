package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateVariableCmd = &cobra.Command{
	Use:   "update-variable",
	Short: "Updates a variable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateVariableCmd).Standalone()

	frauddetector_updateVariableCmd.Flags().String("default-value", "", "The new default value of the variable.")
	frauddetector_updateVariableCmd.Flags().String("description", "", "The new description.")
	frauddetector_updateVariableCmd.Flags().String("name", "", "The name of the variable.")
	frauddetector_updateVariableCmd.Flags().String("variable-type", "", "The variable type.")
	frauddetector_updateVariableCmd.MarkFlagRequired("name")
	frauddetectorCmd.AddCommand(frauddetector_updateVariableCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createVariableCmd = &cobra.Command{
	Use:   "create-variable",
	Short: "Creates a variable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createVariableCmd).Standalone()

	frauddetector_createVariableCmd.Flags().String("data-source", "", "The source of the data.")
	frauddetector_createVariableCmd.Flags().String("data-type", "", "The data type of the variable.")
	frauddetector_createVariableCmd.Flags().String("default-value", "", "The default value for the variable when no value is received.")
	frauddetector_createVariableCmd.Flags().String("description", "", "The description.")
	frauddetector_createVariableCmd.Flags().String("name", "", "The name of the variable.")
	frauddetector_createVariableCmd.Flags().String("tags", "", "A collection of key and value pairs.")
	frauddetector_createVariableCmd.Flags().String("variable-type", "", "The variable type.")
	frauddetector_createVariableCmd.MarkFlagRequired("data-source")
	frauddetector_createVariableCmd.MarkFlagRequired("data-type")
	frauddetector_createVariableCmd.MarkFlagRequired("default-value")
	frauddetector_createVariableCmd.MarkFlagRequired("name")
	frauddetectorCmd.AddCommand(frauddetector_createVariableCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteVariableCmd = &cobra.Command{
	Use:   "delete-variable",
	Short: "Deletes a variable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteVariableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteVariableCmd).Standalone()

		frauddetector_deleteVariableCmd.Flags().String("name", "", "The name of the variable to delete.")
		frauddetector_deleteVariableCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteVariableCmd)
}

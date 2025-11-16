package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getVariablesCmd = &cobra.Command{
	Use:   "get-variables",
	Short: "Gets all of the variables or the specific variable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getVariablesCmd).Standalone()

	frauddetector_getVariablesCmd.Flags().String("max-results", "", "The max size per page determined for the get variable request.")
	frauddetector_getVariablesCmd.Flags().String("name", "", "The name of the variable.")
	frauddetector_getVariablesCmd.Flags().String("next-token", "", "The next page token of the get variable request.")
	frauddetectorCmd.AddCommand(frauddetector_getVariablesCmd)
}

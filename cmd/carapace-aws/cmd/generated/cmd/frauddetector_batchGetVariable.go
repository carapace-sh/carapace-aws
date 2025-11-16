package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_batchGetVariableCmd = &cobra.Command{
	Use:   "batch-get-variable",
	Short: "Gets a batch of variables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_batchGetVariableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_batchGetVariableCmd).Standalone()

		frauddetector_batchGetVariableCmd.Flags().String("names", "", "The list of variable names to get.")
		frauddetector_batchGetVariableCmd.MarkFlagRequired("names")
	})
	frauddetectorCmd.AddCommand(frauddetector_batchGetVariableCmd)
}

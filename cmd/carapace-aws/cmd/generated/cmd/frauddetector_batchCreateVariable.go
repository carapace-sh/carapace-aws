package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_batchCreateVariableCmd = &cobra.Command{
	Use:   "batch-create-variable",
	Short: "Creates a batch of variables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_batchCreateVariableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_batchCreateVariableCmd).Standalone()

		frauddetector_batchCreateVariableCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_batchCreateVariableCmd.Flags().String("variable-entries", "", "The list of variables for the batch create variable request.")
		frauddetector_batchCreateVariableCmd.MarkFlagRequired("variable-entries")
	})
	frauddetectorCmd.AddCommand(frauddetector_batchCreateVariableCmd)
}

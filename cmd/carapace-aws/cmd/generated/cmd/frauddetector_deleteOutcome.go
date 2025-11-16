package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteOutcomeCmd = &cobra.Command{
	Use:   "delete-outcome",
	Short: "Deletes an outcome.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteOutcomeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteOutcomeCmd).Standalone()

		frauddetector_deleteOutcomeCmd.Flags().String("name", "", "The name of the outcome to delete.")
		frauddetector_deleteOutcomeCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteOutcomeCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_putOutcomeCmd = &cobra.Command{
	Use:   "put-outcome",
	Short: "Creates or updates an outcome.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_putOutcomeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_putOutcomeCmd).Standalone()

		frauddetector_putOutcomeCmd.Flags().String("description", "", "The outcome description.")
		frauddetector_putOutcomeCmd.Flags().String("name", "", "The name of the outcome.")
		frauddetector_putOutcomeCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_putOutcomeCmd.MarkFlagRequired("name")
	})
	frauddetectorCmd.AddCommand(frauddetector_putOutcomeCmd)
}

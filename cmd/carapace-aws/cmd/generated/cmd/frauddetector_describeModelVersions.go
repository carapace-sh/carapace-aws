package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_describeModelVersionsCmd = &cobra.Command{
	Use:   "describe-model-versions",
	Short: "Gets all of the model versions for the specified model type or for the specified model type and model ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_describeModelVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_describeModelVersionsCmd).Standalone()

		frauddetector_describeModelVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		frauddetector_describeModelVersionsCmd.Flags().String("model-id", "", "The model ID.")
		frauddetector_describeModelVersionsCmd.Flags().String("model-type", "", "The model type.")
		frauddetector_describeModelVersionsCmd.Flags().String("model-version-number", "", "The model version number.")
		frauddetector_describeModelVersionsCmd.Flags().String("next-token", "", "The next token from the previous results.")
	})
	frauddetectorCmd.AddCommand(frauddetector_describeModelVersionsCmd)
}

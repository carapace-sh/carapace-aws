package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getModelsCmd = &cobra.Command{
	Use:   "get-models",
	Short: "Gets one or more models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getModelsCmd).Standalone()

	frauddetector_getModelsCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
	frauddetector_getModelsCmd.Flags().String("model-id", "", "The model ID.")
	frauddetector_getModelsCmd.Flags().String("model-type", "", "The model type.")
	frauddetector_getModelsCmd.Flags().String("next-token", "", "The next token for the subsequent request.")
	frauddetectorCmd.AddCommand(frauddetector_getModelsCmd)
}

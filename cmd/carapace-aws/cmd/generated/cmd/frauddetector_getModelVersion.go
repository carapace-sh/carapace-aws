package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getModelVersionCmd = &cobra.Command{
	Use:   "get-model-version",
	Short: "Gets the details of the specified model version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getModelVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getModelVersionCmd).Standalone()

		frauddetector_getModelVersionCmd.Flags().String("model-id", "", "The model ID.")
		frauddetector_getModelVersionCmd.Flags().String("model-type", "", "The model type.")
		frauddetector_getModelVersionCmd.Flags().String("model-version-number", "", "The model version number.")
		frauddetector_getModelVersionCmd.MarkFlagRequired("model-id")
		frauddetector_getModelVersionCmd.MarkFlagRequired("model-type")
		frauddetector_getModelVersionCmd.MarkFlagRequired("model-version-number")
	})
	frauddetectorCmd.AddCommand(frauddetector_getModelVersionCmd)
}

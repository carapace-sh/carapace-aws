package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateModelCmd = &cobra.Command{
	Use:   "update-model",
	Short: "Updates model description.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_updateModelCmd).Standalone()

		frauddetector_updateModelCmd.Flags().String("description", "", "The new model description.")
		frauddetector_updateModelCmd.Flags().String("model-id", "", "The model ID.")
		frauddetector_updateModelCmd.Flags().String("model-type", "", "The model type.")
		frauddetector_updateModelCmd.MarkFlagRequired("model-id")
		frauddetector_updateModelCmd.MarkFlagRequired("model-type")
	})
	frauddetectorCmd.AddCommand(frauddetector_updateModelCmd)
}

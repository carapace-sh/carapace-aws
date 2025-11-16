package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateModelVersionStatusCmd = &cobra.Command{
	Use:   "update-model-version-status",
	Short: "Updates the status of a model version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateModelVersionStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_updateModelVersionStatusCmd).Standalone()

		frauddetector_updateModelVersionStatusCmd.Flags().String("model-id", "", "The model ID of the model version to update.")
		frauddetector_updateModelVersionStatusCmd.Flags().String("model-type", "", "The model type.")
		frauddetector_updateModelVersionStatusCmd.Flags().String("model-version-number", "", "The model version number.")
		frauddetector_updateModelVersionStatusCmd.Flags().String("status", "", "The model version status.")
		frauddetector_updateModelVersionStatusCmd.MarkFlagRequired("model-id")
		frauddetector_updateModelVersionStatusCmd.MarkFlagRequired("model-type")
		frauddetector_updateModelVersionStatusCmd.MarkFlagRequired("model-version-number")
		frauddetector_updateModelVersionStatusCmd.MarkFlagRequired("status")
	})
	frauddetectorCmd.AddCommand(frauddetector_updateModelVersionStatusCmd)
}

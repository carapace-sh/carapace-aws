package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteModelVersionCmd = &cobra.Command{
	Use:   "delete-model-version",
	Short: "Deletes a model version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteModelVersionCmd).Standalone()

	frauddetector_deleteModelVersionCmd.Flags().String("model-id", "", "The model ID of the model version to delete.")
	frauddetector_deleteModelVersionCmd.Flags().String("model-type", "", "The model type of the model version to delete.")
	frauddetector_deleteModelVersionCmd.Flags().String("model-version-number", "", "The model version number of the model version to delete.")
	frauddetector_deleteModelVersionCmd.MarkFlagRequired("model-id")
	frauddetector_deleteModelVersionCmd.MarkFlagRequired("model-type")
	frauddetector_deleteModelVersionCmd.MarkFlagRequired("model-version-number")
	frauddetectorCmd.AddCommand(frauddetector_deleteModelVersionCmd)
}

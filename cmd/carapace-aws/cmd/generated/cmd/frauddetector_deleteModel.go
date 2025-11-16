package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteModelCmd = &cobra.Command{
	Use:   "delete-model",
	Short: "Deletes a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteModelCmd).Standalone()

	frauddetector_deleteModelCmd.Flags().String("model-id", "", "The model ID of the model to delete.")
	frauddetector_deleteModelCmd.Flags().String("model-type", "", "The model type of the model to delete.")
	frauddetector_deleteModelCmd.MarkFlagRequired("model-id")
	frauddetector_deleteModelCmd.MarkFlagRequired("model-type")
	frauddetectorCmd.AddCommand(frauddetector_deleteModelCmd)
}

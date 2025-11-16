package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createModelCmd = &cobra.Command{
	Use:   "create-model",
	Short: "Creates a model using the specified model type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createModelCmd).Standalone()

	frauddetector_createModelCmd.Flags().String("description", "", "The model description.")
	frauddetector_createModelCmd.Flags().String("event-type-name", "", "The name of the event type.")
	frauddetector_createModelCmd.Flags().String("model-id", "", "The model ID.")
	frauddetector_createModelCmd.Flags().String("model-type", "", "The model type.")
	frauddetector_createModelCmd.Flags().String("tags", "", "A collection of key and value pairs.")
	frauddetector_createModelCmd.MarkFlagRequired("event-type-name")
	frauddetector_createModelCmd.MarkFlagRequired("model-id")
	frauddetector_createModelCmd.MarkFlagRequired("model-type")
	frauddetectorCmd.AddCommand(frauddetector_createModelCmd)
}

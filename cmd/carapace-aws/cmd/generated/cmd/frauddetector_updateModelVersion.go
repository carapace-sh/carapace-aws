package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateModelVersionCmd = &cobra.Command{
	Use:   "update-model-version",
	Short: "Updates a model version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateModelVersionCmd).Standalone()

	frauddetector_updateModelVersionCmd.Flags().String("external-events-detail", "", "The details of the external events data used for training the model version.")
	frauddetector_updateModelVersionCmd.Flags().String("ingested-events-detail", "", "The details of the ingested event used for training the model version.")
	frauddetector_updateModelVersionCmd.Flags().String("major-version-number", "", "The major version number.")
	frauddetector_updateModelVersionCmd.Flags().String("model-id", "", "The model ID.")
	frauddetector_updateModelVersionCmd.Flags().String("model-type", "", "The model type.")
	frauddetector_updateModelVersionCmd.Flags().String("tags", "", "A collection of key and value pairs.")
	frauddetector_updateModelVersionCmd.MarkFlagRequired("major-version-number")
	frauddetector_updateModelVersionCmd.MarkFlagRequired("model-id")
	frauddetector_updateModelVersionCmd.MarkFlagRequired("model-type")
	frauddetectorCmd.AddCommand(frauddetector_updateModelVersionCmd)
}

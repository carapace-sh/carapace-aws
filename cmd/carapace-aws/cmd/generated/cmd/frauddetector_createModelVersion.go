package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createModelVersionCmd = &cobra.Command{
	Use:   "create-model-version",
	Short: "Creates a version of the model using the specified model type and model id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createModelVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_createModelVersionCmd).Standalone()

		frauddetector_createModelVersionCmd.Flags().String("external-events-detail", "", "Details of the external events data used for model version training.")
		frauddetector_createModelVersionCmd.Flags().String("ingested-events-detail", "", "Details of the ingested events data used for model version training.")
		frauddetector_createModelVersionCmd.Flags().String("model-id", "", "The model ID.")
		frauddetector_createModelVersionCmd.Flags().String("model-type", "", "The model type.")
		frauddetector_createModelVersionCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_createModelVersionCmd.Flags().String("training-data-schema", "", "The training data schema.")
		frauddetector_createModelVersionCmd.Flags().String("training-data-source", "", "The training data source location in Amazon S3.")
		frauddetector_createModelVersionCmd.MarkFlagRequired("model-id")
		frauddetector_createModelVersionCmd.MarkFlagRequired("model-type")
		frauddetector_createModelVersionCmd.MarkFlagRequired("training-data-schema")
		frauddetector_createModelVersionCmd.MarkFlagRequired("training-data-source")
	})
	frauddetectorCmd.AddCommand(frauddetector_createModelVersionCmd)
}

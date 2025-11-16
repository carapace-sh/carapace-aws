package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getExternalModelsCmd = &cobra.Command{
	Use:   "get-external-models",
	Short: "Gets the details for one or more Amazon SageMaker models that have been imported into the service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getExternalModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getExternalModelsCmd).Standalone()

		frauddetector_getExternalModelsCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
		frauddetector_getExternalModelsCmd.Flags().String("model-endpoint", "", "The Amazon SageMaker model endpoint.")
		frauddetector_getExternalModelsCmd.Flags().String("next-token", "", "The next page token for the request.")
	})
	frauddetectorCmd.AddCommand(frauddetector_getExternalModelsCmd)
}

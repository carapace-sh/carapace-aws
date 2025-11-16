package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listModelsCmd = &cobra.Command{
	Use:   "list-models",
	Short: "Generates a list of all models in the account, including model name and ARN, dataset, and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_listModelsCmd).Standalone()

		lookoutequipment_listModelsCmd.Flags().String("dataset-name-begins-with", "", "The beginning of the name of the dataset of the machine learning models to be listed.")
		lookoutequipment_listModelsCmd.Flags().String("max-results", "", "Specifies the maximum number of machine learning models to list.")
		lookoutequipment_listModelsCmd.Flags().String("model-name-begins-with", "", "The beginning of the name of the machine learning models being listed.")
		lookoutequipment_listModelsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of machine learning models.")
		lookoutequipment_listModelsCmd.Flags().String("status", "", "The status of the machine learning model.")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_listModelsCmd)
}

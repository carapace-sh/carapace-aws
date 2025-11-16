package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_importModelVersionCmd = &cobra.Command{
	Use:   "import-model-version",
	Short: "Imports a model that has been trained successfully.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_importModelVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_importModelVersionCmd).Standalone()

		lookoutequipment_importModelVersionCmd.Flags().String("client-token", "", "A unique identifier for the request.")
		lookoutequipment_importModelVersionCmd.Flags().String("dataset-name", "", "The name of the dataset for the machine learning model being imported.")
		lookoutequipment_importModelVersionCmd.Flags().String("inference-data-import-strategy", "", "Indicates how to import the accumulated inference data when a model version is imported.")
		lookoutequipment_importModelVersionCmd.Flags().String("labels-input-configuration", "", "")
		lookoutequipment_importModelVersionCmd.Flags().String("model-name", "", "The name for the machine learning model to be created.")
		lookoutequipment_importModelVersionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of a role with permission to access the data source being used to create the machine learning model.")
		lookoutequipment_importModelVersionCmd.Flags().String("server-side-kms-key-id", "", "Provides the identifier of the KMS key key used to encrypt model data by Amazon Lookout for Equipment.")
		lookoutequipment_importModelVersionCmd.Flags().String("source-model-version-arn", "", "The Amazon Resource Name (ARN) of the model version to import.")
		lookoutequipment_importModelVersionCmd.Flags().String("tags", "", "The tags associated with the machine learning model to be created.")
		lookoutequipment_importModelVersionCmd.MarkFlagRequired("client-token")
		lookoutequipment_importModelVersionCmd.MarkFlagRequired("dataset-name")
		lookoutequipment_importModelVersionCmd.MarkFlagRequired("source-model-version-arn")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_importModelVersionCmd)
}

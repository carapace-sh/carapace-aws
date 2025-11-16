package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_importDatasetCmd = &cobra.Command{
	Use:   "import-dataset",
	Short: "Imports a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_importDatasetCmd).Standalone()

	lookoutequipment_importDatasetCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	lookoutequipment_importDatasetCmd.Flags().String("dataset-name", "", "The name of the machine learning dataset to be created.")
	lookoutequipment_importDatasetCmd.Flags().String("server-side-kms-key-id", "", "Provides the identifier of the KMS key key used to encrypt model data by Amazon Lookout for Equipment.")
	lookoutequipment_importDatasetCmd.Flags().String("source-dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset to import.")
	lookoutequipment_importDatasetCmd.Flags().String("tags", "", "Any tags associated with the dataset to be created.")
	lookoutequipment_importDatasetCmd.MarkFlagRequired("client-token")
	lookoutequipment_importDatasetCmd.MarkFlagRequired("source-dataset-arn")
	lookoutequipmentCmd.AddCommand(lookoutequipment_importDatasetCmd)
}

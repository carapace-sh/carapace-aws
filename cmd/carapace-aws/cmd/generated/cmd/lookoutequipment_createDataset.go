package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Creates a container for a collection of data being ingested for analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_createDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_createDatasetCmd).Standalone()

		lookoutequipment_createDatasetCmd.Flags().String("client-token", "", "A unique identifier for the request.")
		lookoutequipment_createDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset being created.")
		lookoutequipment_createDatasetCmd.Flags().String("dataset-schema", "", "A JSON description of the data that is in each time series dataset, including names, column names, and data types.")
		lookoutequipment_createDatasetCmd.Flags().String("server-side-kms-key-id", "", "Provides the identifier of the KMS key used to encrypt dataset data by Amazon Lookout for Equipment.")
		lookoutequipment_createDatasetCmd.Flags().String("tags", "", "Any tags associated with the ingested data described in the dataset.")
		lookoutequipment_createDatasetCmd.MarkFlagRequired("client-token")
		lookoutequipment_createDatasetCmd.MarkFlagRequired("dataset-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_createDatasetCmd)
}

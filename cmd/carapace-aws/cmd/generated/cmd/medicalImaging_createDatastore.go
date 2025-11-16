package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImaging_createDatastoreCmd = &cobra.Command{
	Use:   "create-datastore",
	Short: "Create a data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImaging_createDatastoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medicalImaging_createDatastoreCmd).Standalone()

		medicalImaging_createDatastoreCmd.Flags().String("client-token", "", "A unique identifier for API idempotency.")
		medicalImaging_createDatastoreCmd.Flags().String("datastore-name", "", "The data store name.")
		medicalImaging_createDatastoreCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) assigned to the Key Management Service (KMS) key for accessing encrypted data.")
		medicalImaging_createDatastoreCmd.Flags().String("lambda-authorizer-arn", "", "The ARN of the authorizer's Lambda function.")
		medicalImaging_createDatastoreCmd.Flags().String("lossless-storage-format", "", "The lossless storage format for the datastore.")
		medicalImaging_createDatastoreCmd.Flags().String("tags", "", "The tags provided when creating a data store.")
		medicalImaging_createDatastoreCmd.MarkFlagRequired("client-token")
	})
	medicalImagingCmd.AddCommand(medicalImaging_createDatastoreCmd)
}

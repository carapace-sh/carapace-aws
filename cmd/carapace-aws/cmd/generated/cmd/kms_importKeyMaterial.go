package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_importKeyMaterialCmd = &cobra.Command{
	Use:   "import-key-material",
	Short: "Imports or reimports key material into an existing KMS key that was created without key material.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_importKeyMaterialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_importKeyMaterialCmd).Standalone()

		kms_importKeyMaterialCmd.Flags().String("encrypted-key-material", "", "The encrypted key material to import.")
		kms_importKeyMaterialCmd.Flags().String("expiration-model", "", "Specifies whether the key material expires.")
		kms_importKeyMaterialCmd.Flags().String("import-token", "", "The import token that you received in the response to a previous [GetParametersForImport]() request.")
		kms_importKeyMaterialCmd.Flags().String("import-type", "", "Indicates whether the key material being imported is previously associated with this KMS key or not.")
		kms_importKeyMaterialCmd.Flags().String("key-id", "", "The identifier of the KMS key that will be associated with the imported key material.")
		kms_importKeyMaterialCmd.Flags().String("key-material-description", "", "Description for the key material being imported.")
		kms_importKeyMaterialCmd.Flags().String("key-material-id", "", "Identifies the key material being imported.")
		kms_importKeyMaterialCmd.Flags().String("valid-to", "", "The date and time when the imported key material expires.")
		kms_importKeyMaterialCmd.MarkFlagRequired("encrypted-key-material")
		kms_importKeyMaterialCmd.MarkFlagRequired("import-token")
		kms_importKeyMaterialCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_importKeyMaterialCmd)
}

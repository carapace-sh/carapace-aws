package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_deleteImportedKeyMaterialCmd = &cobra.Command{
	Use:   "delete-imported-key-material",
	Short: "Deletes key material that was previously imported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_deleteImportedKeyMaterialCmd).Standalone()

	kms_deleteImportedKeyMaterialCmd.Flags().String("key-id", "", "Identifies the KMS key from which you are deleting imported key material.")
	kms_deleteImportedKeyMaterialCmd.Flags().String("key-material-id", "", "Identifies the imported key material you are deleting.")
	kms_deleteImportedKeyMaterialCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_deleteImportedKeyMaterialCmd)
}

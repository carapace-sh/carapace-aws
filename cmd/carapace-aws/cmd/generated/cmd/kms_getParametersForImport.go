package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_getParametersForImportCmd = &cobra.Command{
	Use:   "get-parameters-for-import",
	Short: "Returns the public key and an import token you need to import or reimport key material for a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_getParametersForImportCmd).Standalone()

	kms_getParametersForImportCmd.Flags().String("key-id", "", "The identifier of the KMS key that will be associated with the imported key material.")
	kms_getParametersForImportCmd.Flags().String("wrapping-algorithm", "", "The algorithm you will use with the RSA public key (`PublicKey`) in the response to protect your key material during import.")
	kms_getParametersForImportCmd.Flags().String("wrapping-key-spec", "", "The type of RSA public key to return in the response.")
	kms_getParametersForImportCmd.MarkFlagRequired("key-id")
	kms_getParametersForImportCmd.MarkFlagRequired("wrapping-algorithm")
	kms_getParametersForImportCmd.MarkFlagRequired("wrapping-key-spec")
	kmsCmd.AddCommand(kms_getParametersForImportCmd)
}

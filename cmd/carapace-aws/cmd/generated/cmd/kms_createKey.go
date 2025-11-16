package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_createKeyCmd = &cobra.Command{
	Use:   "create-key",
	Short: "Creates a unique customer managed [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms-keys) in your Amazon Web Services account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_createKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_createKeyCmd).Standalone()

		kms_createKeyCmd.Flags().String("bypass-policy-lockout-safety-check", "", "Skips (\"bypasses\") the key policy lockout safety check.")
		kms_createKeyCmd.Flags().String("custom-key-store-id", "", "Creates the KMS key in the specified [custom key store](https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html).")
		kms_createKeyCmd.Flags().String("customer-master-key-spec", "", "Instead, use the `KeySpec` parameter.")
		kms_createKeyCmd.Flags().String("description", "", "A description of the KMS key.")
		kms_createKeyCmd.Flags().String("key-spec", "", "Specifies the type of KMS key to create.")
		kms_createKeyCmd.Flags().String("key-usage", "", "Determines the [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations) for which you can use the KMS key.")
		kms_createKeyCmd.Flags().String("multi-region", "", "Creates a multi-Region primary key that you can replicate into other Amazon Web Services Regions.")
		kms_createKeyCmd.Flags().String("origin", "", "The source of the key material for the KMS key.")
		kms_createKeyCmd.Flags().String("policy", "", "The key policy to attach to the KMS key.")
		kms_createKeyCmd.Flags().String("tags", "", "Assigns one or more tags to the KMS key.")
		kms_createKeyCmd.Flags().String("xks-key-id", "", "Identifies the [external key](https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html#concept-external-key) that serves as key material for the KMS key in an [external key store](https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html).")
	})
	kmsCmd.AddCommand(kms_createKeyCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_deleteCustomKeyStoreCmd = &cobra.Command{
	Use:   "delete-custom-key-store",
	Short: "Deletes a [custom key store](https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_deleteCustomKeyStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_deleteCustomKeyStoreCmd).Standalone()

		kms_deleteCustomKeyStoreCmd.Flags().String("custom-key-store-id", "", "Enter the ID of the custom key store you want to delete.")
		kms_deleteCustomKeyStoreCmd.MarkFlagRequired("custom-key-store-id")
	})
	kmsCmd.AddCommand(kms_deleteCustomKeyStoreCmd)
}

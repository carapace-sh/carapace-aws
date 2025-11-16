package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_disconnectCustomKeyStoreCmd = &cobra.Command{
	Use:   "disconnect-custom-key-store",
	Short: "Disconnects the [custom key store](https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html) from its backing key store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_disconnectCustomKeyStoreCmd).Standalone()

	kms_disconnectCustomKeyStoreCmd.Flags().String("custom-key-store-id", "", "Enter the ID of the custom key store you want to disconnect.")
	kms_disconnectCustomKeyStoreCmd.MarkFlagRequired("custom-key-store-id")
	kmsCmd.AddCommand(kms_disconnectCustomKeyStoreCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_connectCustomKeyStoreCmd = &cobra.Command{
	Use:   "connect-custom-key-store",
	Short: "Connects or reconnects a [custom key store](https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html) to its backing key store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_connectCustomKeyStoreCmd).Standalone()

	kms_connectCustomKeyStoreCmd.Flags().String("custom-key-store-id", "", "Enter the key store ID of the custom key store that you want to connect.")
	kms_connectCustomKeyStoreCmd.MarkFlagRequired("custom-key-store-id")
	kmsCmd.AddCommand(kms_connectCustomKeyStoreCmd)
}

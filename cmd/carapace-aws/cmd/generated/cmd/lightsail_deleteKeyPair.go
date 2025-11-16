package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteKeyPairCmd = &cobra.Command{
	Use:   "delete-key-pair",
	Short: "Deletes the specified key pair by removing the public key from Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteKeyPairCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteKeyPairCmd).Standalone()

		lightsail_deleteKeyPairCmd.Flags().String("expected-fingerprint", "", "The RSA fingerprint of the Lightsail default key pair to delete.")
		lightsail_deleteKeyPairCmd.Flags().String("key-pair-name", "", "The name of the key pair to delete.")
		lightsail_deleteKeyPairCmd.MarkFlagRequired("key-pair-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteKeyPairCmd)
}

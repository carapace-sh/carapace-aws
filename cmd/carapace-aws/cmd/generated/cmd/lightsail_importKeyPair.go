package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_importKeyPairCmd = &cobra.Command{
	Use:   "import-key-pair",
	Short: "Imports a public SSH key from a specific key pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_importKeyPairCmd).Standalone()

	lightsail_importKeyPairCmd.Flags().String("key-pair-name", "", "The name of the key pair for which you want to import the public key.")
	lightsail_importKeyPairCmd.Flags().String("public-key-base64", "", "A base64-encoded public key of the `ssh-rsa` type.")
	lightsail_importKeyPairCmd.MarkFlagRequired("key-pair-name")
	lightsail_importKeyPairCmd.MarkFlagRequired("public-key-base64")
	lightsailCmd.AddCommand(lightsail_importKeyPairCmd)
}

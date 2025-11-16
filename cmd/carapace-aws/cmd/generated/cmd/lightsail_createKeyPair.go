package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createKeyPairCmd = &cobra.Command{
	Use:   "create-key-pair",
	Short: "Creates a custom SSH key pair that you can use with an Amazon Lightsail instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createKeyPairCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createKeyPairCmd).Standalone()

		lightsail_createKeyPairCmd.Flags().String("key-pair-name", "", "The name for your new key pair.")
		lightsail_createKeyPairCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createKeyPairCmd.MarkFlagRequired("key-pair-name")
	})
	lightsailCmd.AddCommand(lightsail_createKeyPairCmd)
}

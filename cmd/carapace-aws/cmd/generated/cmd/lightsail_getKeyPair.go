package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getKeyPairCmd = &cobra.Command{
	Use:   "get-key-pair",
	Short: "Returns information about a specific key pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getKeyPairCmd).Standalone()

	lightsail_getKeyPairCmd.Flags().String("key-pair-name", "", "The name of the key pair for which you are requesting information.")
	lightsail_getKeyPairCmd.MarkFlagRequired("key-pair-name")
	lightsailCmd.AddCommand(lightsail_getKeyPairCmd)
}

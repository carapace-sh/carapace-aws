package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_downloadDefaultKeyPairCmd = &cobra.Command{
	Use:   "download-default-key-pair",
	Short: "Downloads the regional Amazon Lightsail default key pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_downloadDefaultKeyPairCmd).Standalone()

	lightsailCmd.AddCommand(lightsail_downloadDefaultKeyPairCmd)
}

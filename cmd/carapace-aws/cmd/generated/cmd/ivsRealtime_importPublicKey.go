package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_importPublicKeyCmd = &cobra.Command{
	Use:   "import-public-key",
	Short: "Import a public key to be used for signing stage participant tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_importPublicKeyCmd).Standalone()

	ivsRealtime_importPublicKeyCmd.Flags().String("name", "", "Name of the public key to be imported.")
	ivsRealtime_importPublicKeyCmd.Flags().String("public-key-material", "", "The content of the public key to be imported.")
	ivsRealtime_importPublicKeyCmd.Flags().String("tags", "", "Tags attached to the resource.")
	ivsRealtime_importPublicKeyCmd.MarkFlagRequired("public-key-material")
	ivsRealtimeCmd.AddCommand(ivsRealtime_importPublicKeyCmd)
}

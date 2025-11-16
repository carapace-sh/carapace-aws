package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_importDecoderManifestCmd = &cobra.Command{
	Use:   "import-decoder-manifest",
	Short: "Creates a decoder manifest using your existing CAN DBC file from your local device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_importDecoderManifestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_importDecoderManifestCmd).Standalone()

		iotfleetwise_importDecoderManifestCmd.Flags().String("name", "", "The name of the decoder manifest to import.")
		iotfleetwise_importDecoderManifestCmd.Flags().String("network-file-definitions", "", "The file to load into an Amazon Web Services account.")
		iotfleetwise_importDecoderManifestCmd.MarkFlagRequired("name")
		iotfleetwise_importDecoderManifestCmd.MarkFlagRequired("network-file-definitions")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_importDecoderManifestCmd)
}

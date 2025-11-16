package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_updateDecoderManifestCmd = &cobra.Command{
	Use:   "update-decoder-manifest",
	Short: "Updates a decoder manifest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_updateDecoderManifestCmd).Standalone()

	iotfleetwise_updateDecoderManifestCmd.Flags().String("default-for-unmapped-signals", "", "Use default decoders for all unmapped signals in the model.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("description", "", "A brief description of the decoder manifest to update.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("name", "", "The name of the decoder manifest to update.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("network-interfaces-to-add", "", "A list of information about the network interfaces to add to the decoder manifest.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("network-interfaces-to-remove", "", "A list of network interfaces to remove from the decoder manifest.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("network-interfaces-to-update", "", "A list of information about the network interfaces to update in the decoder manifest.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("signal-decoders-to-add", "", "A list of information about decoding additional signals to add to the decoder manifest.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("signal-decoders-to-remove", "", "A list of signal decoders to remove from the decoder manifest.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("signal-decoders-to-update", "", "A list of updated information about decoding signals to update in the decoder manifest.")
	iotfleetwise_updateDecoderManifestCmd.Flags().String("status", "", "The state of the decoder manifest.")
	iotfleetwise_updateDecoderManifestCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_updateDecoderManifestCmd)
}

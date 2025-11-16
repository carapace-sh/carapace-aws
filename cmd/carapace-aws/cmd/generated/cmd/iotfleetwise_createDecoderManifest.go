package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_createDecoderManifestCmd = &cobra.Command{
	Use:   "create-decoder-manifest",
	Short: "Creates the decoder manifest associated with a model manifest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_createDecoderManifestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_createDecoderManifestCmd).Standalone()

		iotfleetwise_createDecoderManifestCmd.Flags().String("default-for-unmapped-signals", "", "Use default decoders for all unmapped signals in the model.")
		iotfleetwise_createDecoderManifestCmd.Flags().String("description", "", "A brief description of the decoder manifest.")
		iotfleetwise_createDecoderManifestCmd.Flags().String("model-manifest-arn", "", "The Amazon Resource Name (ARN) of the vehicle model (model manifest).")
		iotfleetwise_createDecoderManifestCmd.Flags().String("name", "", "The unique name of the decoder manifest to create.")
		iotfleetwise_createDecoderManifestCmd.Flags().String("network-interfaces", "", "A list of information about available network interfaces.")
		iotfleetwise_createDecoderManifestCmd.Flags().String("signal-decoders", "", "A list of information about signal decoders.")
		iotfleetwise_createDecoderManifestCmd.Flags().String("tags", "", "Metadata that can be used to manage the decoder manifest.")
		iotfleetwise_createDecoderManifestCmd.MarkFlagRequired("model-manifest-arn")
		iotfleetwise_createDecoderManifestCmd.MarkFlagRequired("name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_createDecoderManifestCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_deleteDecoderManifestCmd = &cobra.Command{
	Use:   "delete-decoder-manifest",
	Short: "Deletes a decoder manifest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_deleteDecoderManifestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_deleteDecoderManifestCmd).Standalone()

		iotfleetwise_deleteDecoderManifestCmd.Flags().String("name", "", "The name of the decoder manifest to delete.")
		iotfleetwise_deleteDecoderManifestCmd.MarkFlagRequired("name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_deleteDecoderManifestCmd)
}

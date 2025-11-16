package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listDecoderManifestsCmd = &cobra.Command{
	Use:   "list-decoder-manifests",
	Short: "Lists decoder manifests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listDecoderManifestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listDecoderManifestsCmd).Standalone()

		iotfleetwise_listDecoderManifestsCmd.Flags().String("list-response-scope", "", "When you set the `listResponseScope` parameter to `METADATA_ONLY`, the list response includes: decoder manifest name, Amazon Resource Name (ARN), creation time, and last modification time.")
		iotfleetwise_listDecoderManifestsCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listDecoderManifestsCmd.Flags().String("model-manifest-arn", "", "The Amazon Resource Name (ARN) of a vehicle model (model manifest) associated with the decoder manifest.")
		iotfleetwise_listDecoderManifestsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listDecoderManifestsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listDecoderManifestSignalsCmd = &cobra.Command{
	Use:   "list-decoder-manifest-signals",
	Short: "A list of information about signal decoders specified in a decoder manifest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listDecoderManifestSignalsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listDecoderManifestSignalsCmd).Standalone()

		iotfleetwise_listDecoderManifestSignalsCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listDecoderManifestSignalsCmd.Flags().String("name", "", "The name of the decoder manifest to list information about.")
		iotfleetwise_listDecoderManifestSignalsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
		iotfleetwise_listDecoderManifestSignalsCmd.MarkFlagRequired("name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listDecoderManifestSignalsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getDecoderManifestCmd = &cobra.Command{
	Use:   "get-decoder-manifest",
	Short: "Retrieves information about a created decoder manifest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getDecoderManifestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_getDecoderManifestCmd).Standalone()

		iotfleetwise_getDecoderManifestCmd.Flags().String("name", "", "The name of the decoder manifest to retrieve information about.")
		iotfleetwise_getDecoderManifestCmd.MarkFlagRequired("name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_getDecoderManifestCmd)
}

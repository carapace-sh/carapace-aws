package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_getDataEndpointCmd = &cobra.Command{
	Use:   "get-data-endpoint",
	Short: "Gets an endpoint for a specified stream for either reading or writing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_getDataEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_getDataEndpointCmd).Standalone()

		kinesisvideo_getDataEndpointCmd.Flags().String("apiname", "", "The name of the API action for which to get an endpoint.")
		kinesisvideo_getDataEndpointCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream that you want to get the endpoint for.")
		kinesisvideo_getDataEndpointCmd.Flags().String("stream-name", "", "The name of the stream that you want to get the endpoint for.")
		kinesisvideo_getDataEndpointCmd.MarkFlagRequired("apiname")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_getDataEndpointCmd)
}

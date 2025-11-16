package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_describeMappedResourceConfigurationCmd = &cobra.Command{
	Use:   "describe-mapped-resource-configuration",
	Short: "Returns the most current information about the stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_describeMappedResourceConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_describeMappedResourceConfigurationCmd).Standalone()

		kinesisvideo_describeMappedResourceConfigurationCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		kinesisvideo_describeMappedResourceConfigurationCmd.Flags().String("next-token", "", "The token to provide in your next request, to get another batch of results.")
		kinesisvideo_describeMappedResourceConfigurationCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream.")
		kinesisvideo_describeMappedResourceConfigurationCmd.Flags().String("stream-name", "", "The name of the stream.")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_describeMappedResourceConfigurationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getStreamingDistributionConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-streaming-distribution-config2020_05_31",
	Short: "Get the configuration information about a streaming distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getStreamingDistributionConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getStreamingDistributionConfig2020_05_31Cmd).Standalone()

		cloudfront_getStreamingDistributionConfig2020_05_31Cmd.Flags().String("id", "", "The streaming distribution's ID.")
		cloudfront_getStreamingDistributionConfig2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getStreamingDistributionConfig2020_05_31Cmd)
}

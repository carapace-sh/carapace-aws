package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateStreamingDistribution2020_05_31Cmd = &cobra.Command{
	Use:   "update-streaming-distribution2020_05_31",
	Short: "Update a streaming distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateStreamingDistribution2020_05_31Cmd).Standalone()

	cloudfront_updateStreamingDistribution2020_05_31Cmd.Flags().String("id", "", "The streaming distribution's id.")
	cloudfront_updateStreamingDistribution2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the streaming distribution's configuration.")
	cloudfront_updateStreamingDistribution2020_05_31Cmd.Flags().String("streaming-distribution-config", "", "The streaming distribution's configuration information.")
	cloudfront_updateStreamingDistribution2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_updateStreamingDistribution2020_05_31Cmd.MarkFlagRequired("streaming-distribution-config")
	cloudfrontCmd.AddCommand(cloudfront_updateStreamingDistribution2020_05_31Cmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getStreamingDistribution2020_05_31Cmd = &cobra.Command{
	Use:   "get-streaming-distribution2020_05_31",
	Short: "Gets information about a specified RTMP distribution, including the distribution configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getStreamingDistribution2020_05_31Cmd).Standalone()

	cloudfront_getStreamingDistribution2020_05_31Cmd.Flags().String("id", "", "The streaming distribution's ID.")
	cloudfront_getStreamingDistribution2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getStreamingDistribution2020_05_31Cmd)
}

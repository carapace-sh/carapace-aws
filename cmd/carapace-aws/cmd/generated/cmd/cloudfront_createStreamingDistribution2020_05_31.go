package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createStreamingDistribution2020_05_31Cmd = &cobra.Command{
	Use:   "create-streaming-distribution2020_05_31",
	Short: "This API is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createStreamingDistribution2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createStreamingDistribution2020_05_31Cmd).Standalone()

		cloudfront_createStreamingDistribution2020_05_31Cmd.Flags().String("streaming-distribution-config", "", "The streaming distribution's configuration information.")
		cloudfront_createStreamingDistribution2020_05_31Cmd.MarkFlagRequired("streaming-distribution-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_createStreamingDistribution2020_05_31Cmd)
}

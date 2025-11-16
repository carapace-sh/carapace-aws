package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateDistribution2020_05_31Cmd = &cobra.Command{
	Use:   "update-distribution2020_05_31",
	Short: "Updates the configuration for a CloudFront distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateDistribution2020_05_31Cmd).Standalone()

	cloudfront_updateDistribution2020_05_31Cmd.Flags().String("distribution-config", "", "The distribution's configuration information.")
	cloudfront_updateDistribution2020_05_31Cmd.Flags().String("id", "", "The distribution's id.")
	cloudfront_updateDistribution2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the distribution's configuration.")
	cloudfront_updateDistribution2020_05_31Cmd.MarkFlagRequired("distribution-config")
	cloudfront_updateDistribution2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_updateDistribution2020_05_31Cmd)
}

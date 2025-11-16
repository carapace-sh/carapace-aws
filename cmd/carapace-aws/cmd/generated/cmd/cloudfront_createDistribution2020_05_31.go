package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createDistribution2020_05_31Cmd = &cobra.Command{
	Use:   "create-distribution2020_05_31",
	Short: "Creates a CloudFront distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createDistribution2020_05_31Cmd).Standalone()

	cloudfront_createDistribution2020_05_31Cmd.Flags().String("distribution-config", "", "The distribution's configuration information.")
	cloudfront_createDistribution2020_05_31Cmd.MarkFlagRequired("distribution-config")
	cloudfrontCmd.AddCommand(cloudfront_createDistribution2020_05_31Cmd)
}

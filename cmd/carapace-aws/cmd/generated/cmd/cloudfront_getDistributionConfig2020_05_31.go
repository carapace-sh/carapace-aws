package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getDistributionConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-distribution-config2020_05_31",
	Short: "Get the configuration information about a distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getDistributionConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getDistributionConfig2020_05_31Cmd).Standalone()

		cloudfront_getDistributionConfig2020_05_31Cmd.Flags().String("id", "", "The distribution's ID.")
		cloudfront_getDistributionConfig2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getDistributionConfig2020_05_31Cmd)
}

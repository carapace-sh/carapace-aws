package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createDistributionWithTags2020_05_31Cmd = &cobra.Command{
	Use:   "create-distribution-with-tags2020_05_31",
	Short: "Create a new distribution with tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createDistributionWithTags2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createDistributionWithTags2020_05_31Cmd).Standalone()

		cloudfront_createDistributionWithTags2020_05_31Cmd.Flags().String("distribution-config-with-tags", "", "The distribution's configuration information.")
		cloudfront_createDistributionWithTags2020_05_31Cmd.MarkFlagRequired("distribution-config-with-tags")
	})
	cloudfrontCmd.AddCommand(cloudfront_createDistributionWithTags2020_05_31Cmd)
}

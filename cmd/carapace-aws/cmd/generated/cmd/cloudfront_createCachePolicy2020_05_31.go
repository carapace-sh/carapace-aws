package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createCachePolicy2020_05_31Cmd = &cobra.Command{
	Use:   "create-cache-policy2020_05_31",
	Short: "Creates a cache policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createCachePolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createCachePolicy2020_05_31Cmd).Standalone()

		cloudfront_createCachePolicy2020_05_31Cmd.Flags().String("cache-policy-config", "", "A cache policy configuration.")
		cloudfront_createCachePolicy2020_05_31Cmd.MarkFlagRequired("cache-policy-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_createCachePolicy2020_05_31Cmd)
}

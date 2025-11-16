package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateCachePolicy2020_05_31Cmd = &cobra.Command{
	Use:   "update-cache-policy2020_05_31",
	Short: "Updates a cache policy configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateCachePolicy2020_05_31Cmd).Standalone()

	cloudfront_updateCachePolicy2020_05_31Cmd.Flags().String("cache-policy-config", "", "A cache policy configuration.")
	cloudfront_updateCachePolicy2020_05_31Cmd.Flags().String("id", "", "The unique identifier for the cache policy that you are updating.")
	cloudfront_updateCachePolicy2020_05_31Cmd.Flags().String("if-match", "", "The version of the cache policy that you are updating.")
	cloudfront_updateCachePolicy2020_05_31Cmd.MarkFlagRequired("cache-policy-config")
	cloudfront_updateCachePolicy2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_updateCachePolicy2020_05_31Cmd)
}

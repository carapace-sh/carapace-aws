package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getCachePolicyConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-cache-policy-config2020_05_31",
	Short: "Gets a cache policy configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getCachePolicyConfig2020_05_31Cmd).Standalone()

	cloudfront_getCachePolicyConfig2020_05_31Cmd.Flags().String("id", "", "The unique identifier for the cache policy.")
	cloudfront_getCachePolicyConfig2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getCachePolicyConfig2020_05_31Cmd)
}

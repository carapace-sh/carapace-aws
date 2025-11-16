package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getCachePolicy2020_05_31Cmd = &cobra.Command{
	Use:   "get-cache-policy2020_05_31",
	Short: "Gets a cache policy, including the following metadata:\n\n- The policy's identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getCachePolicy2020_05_31Cmd).Standalone()

	cloudfront_getCachePolicy2020_05_31Cmd.Flags().String("id", "", "The unique identifier for the cache policy.")
	cloudfront_getCachePolicy2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getCachePolicy2020_05_31Cmd)
}

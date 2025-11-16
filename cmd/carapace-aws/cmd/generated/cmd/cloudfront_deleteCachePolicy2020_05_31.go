package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteCachePolicy2020_05_31Cmd = &cobra.Command{
	Use:   "delete-cache-policy2020_05_31",
	Short: "Deletes a cache policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteCachePolicy2020_05_31Cmd).Standalone()

	cloudfront_deleteCachePolicy2020_05_31Cmd.Flags().String("id", "", "The unique identifier for the cache policy that you are deleting.")
	cloudfront_deleteCachePolicy2020_05_31Cmd.Flags().String("if-match", "", "The version of the cache policy that you are deleting.")
	cloudfront_deleteCachePolicy2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_deleteCachePolicy2020_05_31Cmd)
}

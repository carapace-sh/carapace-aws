package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteStreamingDistribution2020_05_31Cmd = &cobra.Command{
	Use:   "delete-streaming-distribution2020_05_31",
	Short: "Delete a streaming distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteStreamingDistribution2020_05_31Cmd).Standalone()

	cloudfront_deleteStreamingDistribution2020_05_31Cmd.Flags().String("id", "", "The distribution ID.")
	cloudfront_deleteStreamingDistribution2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when you disabled the streaming distribution.")
	cloudfront_deleteStreamingDistribution2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_deleteStreamingDistribution2020_05_31Cmd)
}

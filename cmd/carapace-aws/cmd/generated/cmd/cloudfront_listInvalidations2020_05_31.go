package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listInvalidations2020_05_31Cmd = &cobra.Command{
	Use:   "list-invalidations2020_05_31",
	Short: "Lists invalidation batches.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listInvalidations2020_05_31Cmd).Standalone()

	cloudfront_listInvalidations2020_05_31Cmd.Flags().String("distribution-id", "", "The distribution's ID.")
	cloudfront_listInvalidations2020_05_31Cmd.Flags().String("marker", "", "Use this parameter when paginating results to indicate where to begin in your list of invalidation batches.")
	cloudfront_listInvalidations2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of invalidation batches that you want in the response body.")
	cloudfront_listInvalidations2020_05_31Cmd.MarkFlagRequired("distribution-id")
	cloudfrontCmd.AddCommand(cloudfront_listInvalidations2020_05_31Cmd)
}

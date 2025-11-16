package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributions2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions2020_05_31",
	Short: "List CloudFront distributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributions2020_05_31Cmd).Standalone()

	cloudfront_listDistributions2020_05_31Cmd.Flags().String("marker", "", "Use this when paginating results to indicate where to begin in your list of distributions.")
	cloudfront_listDistributions2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distributions you want in the response body.")
	cloudfrontCmd.AddCommand(cloudfront_listDistributions2020_05_31Cmd)
}

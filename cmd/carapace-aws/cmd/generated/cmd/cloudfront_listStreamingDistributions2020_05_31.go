package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listStreamingDistributions2020_05_31Cmd = &cobra.Command{
	Use:   "list-streaming-distributions2020_05_31",
	Short: "List streaming distributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listStreamingDistributions2020_05_31Cmd).Standalone()

	cloudfront_listStreamingDistributions2020_05_31Cmd.Flags().String("marker", "", "The value that you provided for the `Marker` request parameter.")
	cloudfront_listStreamingDistributions2020_05_31Cmd.Flags().String("max-items", "", "The value that you provided for the `MaxItems` request parameter.")
	cloudfrontCmd.AddCommand(cloudfront_listStreamingDistributions2020_05_31Cmd)
}

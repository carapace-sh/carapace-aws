package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listOriginAccessControls2020_05_31Cmd = &cobra.Command{
	Use:   "list-origin-access-controls2020_05_31",
	Short: "Gets the list of CloudFront origin access controls (OACs) in this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listOriginAccessControls2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listOriginAccessControls2020_05_31Cmd).Standalone()

		cloudfront_listOriginAccessControls2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of origin access controls.")
		cloudfront_listOriginAccessControls2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of origin access controls that you want in the response.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listOriginAccessControls2020_05_31Cmd)
}

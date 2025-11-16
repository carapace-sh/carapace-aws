package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listPublicKeys2020_05_31Cmd = &cobra.Command{
	Use:   "list-public-keys2020_05_31",
	Short: "List all public keys that have been added to CloudFront for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listPublicKeys2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listPublicKeys2020_05_31Cmd).Standalone()

		cloudfront_listPublicKeys2020_05_31Cmd.Flags().String("marker", "", "Use this when paginating results to indicate where to begin in your list of public keys.")
		cloudfront_listPublicKeys2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of public keys you want in the response body.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listPublicKeys2020_05_31Cmd)
}

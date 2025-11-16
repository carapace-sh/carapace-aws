package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteCloudFrontOriginAccessIdentity2020_05_31Cmd = &cobra.Command{
	Use:   "delete-cloud-front-origin-access-identity2020_05_31",
	Short: "Delete an origin access identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteCloudFrontOriginAccessIdentity2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_deleteCloudFrontOriginAccessIdentity2020_05_31Cmd).Standalone()

		cloudfront_deleteCloudFrontOriginAccessIdentity2020_05_31Cmd.Flags().String("id", "", "The origin access identity's ID.")
		cloudfront_deleteCloudFrontOriginAccessIdentity2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header you received from a previous `GET` or `PUT` request.")
		cloudfront_deleteCloudFrontOriginAccessIdentity2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_deleteCloudFrontOriginAccessIdentity2020_05_31Cmd)
}

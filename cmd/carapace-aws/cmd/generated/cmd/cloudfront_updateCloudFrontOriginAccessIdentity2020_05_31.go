package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd = &cobra.Command{
	Use:   "update-cloud-front-origin-access-identity2020_05_31",
	Short: "Update an origin access identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd).Standalone()

	cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd.Flags().String("cloud-front-origin-access-identity-config", "", "The identity's configuration information.")
	cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd.Flags().String("id", "", "The identity's id.")
	cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the identity's configuration.")
	cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd.MarkFlagRequired("cloud-front-origin-access-identity-config")
	cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_updateCloudFrontOriginAccessIdentity2020_05_31Cmd)
}

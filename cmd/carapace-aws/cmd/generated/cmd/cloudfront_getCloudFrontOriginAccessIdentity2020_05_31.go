package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getCloudFrontOriginAccessIdentity2020_05_31Cmd = &cobra.Command{
	Use:   "get-cloud-front-origin-access-identity2020_05_31",
	Short: "Get the information about an origin access identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getCloudFrontOriginAccessIdentity2020_05_31Cmd).Standalone()

	cloudfront_getCloudFrontOriginAccessIdentity2020_05_31Cmd.Flags().String("id", "", "The identity's ID.")
	cloudfront_getCloudFrontOriginAccessIdentity2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getCloudFrontOriginAccessIdentity2020_05_31Cmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createCloudFrontOriginAccessIdentity2020_05_31Cmd = &cobra.Command{
	Use:   "create-cloud-front-origin-access-identity2020_05_31",
	Short: "Creates a new origin access identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createCloudFrontOriginAccessIdentity2020_05_31Cmd).Standalone()

	cloudfront_createCloudFrontOriginAccessIdentity2020_05_31Cmd.Flags().String("cloud-front-origin-access-identity-config", "", "The current configuration information for the identity.")
	cloudfront_createCloudFrontOriginAccessIdentity2020_05_31Cmd.MarkFlagRequired("cloud-front-origin-access-identity-config")
	cloudfrontCmd.AddCommand(cloudfront_createCloudFrontOriginAccessIdentity2020_05_31Cmd)
}

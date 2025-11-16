package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateOriginAccessControl2020_05_31Cmd = &cobra.Command{
	Use:   "update-origin-access-control2020_05_31",
	Short: "Updates a CloudFront origin access control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateOriginAccessControl2020_05_31Cmd).Standalone()

	cloudfront_updateOriginAccessControl2020_05_31Cmd.Flags().String("id", "", "The unique identifier of the origin access control that you are updating.")
	cloudfront_updateOriginAccessControl2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the origin access control that you are updating.")
	cloudfront_updateOriginAccessControl2020_05_31Cmd.Flags().String("origin-access-control-config", "", "An origin access control.")
	cloudfront_updateOriginAccessControl2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_updateOriginAccessControl2020_05_31Cmd.MarkFlagRequired("origin-access-control-config")
	cloudfrontCmd.AddCommand(cloudfront_updateOriginAccessControl2020_05_31Cmd)
}

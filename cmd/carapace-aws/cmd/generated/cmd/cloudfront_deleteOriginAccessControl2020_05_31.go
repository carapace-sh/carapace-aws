package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteOriginAccessControl2020_05_31Cmd = &cobra.Command{
	Use:   "delete-origin-access-control2020_05_31",
	Short: "Deletes a CloudFront origin access control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteOriginAccessControl2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_deleteOriginAccessControl2020_05_31Cmd).Standalone()

		cloudfront_deleteOriginAccessControl2020_05_31Cmd.Flags().String("id", "", "The unique identifier of the origin access control that you are deleting.")
		cloudfront_deleteOriginAccessControl2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the origin access control that you are deleting.")
		cloudfront_deleteOriginAccessControl2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_deleteOriginAccessControl2020_05_31Cmd)
}

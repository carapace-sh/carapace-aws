package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getOriginAccessControlConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-origin-access-control-config2020_05_31",
	Short: "Gets a CloudFront origin access control configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getOriginAccessControlConfig2020_05_31Cmd).Standalone()

	cloudfront_getOriginAccessControlConfig2020_05_31Cmd.Flags().String("id", "", "The unique identifier of the origin access control.")
	cloudfront_getOriginAccessControlConfig2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getOriginAccessControlConfig2020_05_31Cmd)
}

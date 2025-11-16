package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createOriginAccessControl2020_05_31Cmd = &cobra.Command{
	Use:   "create-origin-access-control2020_05_31",
	Short: "Creates a new origin access control in CloudFront.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createOriginAccessControl2020_05_31Cmd).Standalone()

	cloudfront_createOriginAccessControl2020_05_31Cmd.Flags().String("origin-access-control-config", "", "Contains the origin access control.")
	cloudfront_createOriginAccessControl2020_05_31Cmd.MarkFlagRequired("origin-access-control-config")
	cloudfrontCmd.AddCommand(cloudfront_createOriginAccessControl2020_05_31Cmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createVpcOrigin2020_05_31Cmd = &cobra.Command{
	Use:   "create-vpc-origin2020_05_31",
	Short: "Create an Amazon CloudFront VPC origin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createVpcOrigin2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createVpcOrigin2020_05_31Cmd).Standalone()

		cloudfront_createVpcOrigin2020_05_31Cmd.Flags().String("tags", "", "")
		cloudfront_createVpcOrigin2020_05_31Cmd.Flags().String("vpc-origin-endpoint-config", "", "The VPC origin endpoint configuration.")
		cloudfront_createVpcOrigin2020_05_31Cmd.MarkFlagRequired("vpc-origin-endpoint-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_createVpcOrigin2020_05_31Cmd)
}

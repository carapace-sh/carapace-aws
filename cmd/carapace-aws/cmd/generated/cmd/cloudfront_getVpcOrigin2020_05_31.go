package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getVpcOrigin2020_05_31Cmd = &cobra.Command{
	Use:   "get-vpc-origin2020_05_31",
	Short: "Get the details of an Amazon CloudFront VPC origin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getVpcOrigin2020_05_31Cmd).Standalone()

	cloudfront_getVpcOrigin2020_05_31Cmd.Flags().String("id", "", "The VPC origin ID.")
	cloudfront_getVpcOrigin2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getVpcOrigin2020_05_31Cmd)
}

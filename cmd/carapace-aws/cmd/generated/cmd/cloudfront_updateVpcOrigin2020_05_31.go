package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateVpcOrigin2020_05_31Cmd = &cobra.Command{
	Use:   "update-vpc-origin2020_05_31",
	Short: "Update an Amazon CloudFront VPC origin in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateVpcOrigin2020_05_31Cmd).Standalone()

	cloudfront_updateVpcOrigin2020_05_31Cmd.Flags().String("id", "", "The VPC origin ID.")
	cloudfront_updateVpcOrigin2020_05_31Cmd.Flags().String("if-match", "", "The VPC origin to update, if a match occurs.")
	cloudfront_updateVpcOrigin2020_05_31Cmd.Flags().String("vpc-origin-endpoint-config", "", "The VPC origin endpoint configuration.")
	cloudfront_updateVpcOrigin2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_updateVpcOrigin2020_05_31Cmd.MarkFlagRequired("if-match")
	cloudfront_updateVpcOrigin2020_05_31Cmd.MarkFlagRequired("vpc-origin-endpoint-config")
	cloudfrontCmd.AddCommand(cloudfront_updateVpcOrigin2020_05_31Cmd)
}

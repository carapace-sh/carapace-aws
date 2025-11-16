package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteVpcOrigin2020_05_31Cmd = &cobra.Command{
	Use:   "delete-vpc-origin2020_05_31",
	Short: "Delete an Amazon CloudFront VPC origin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteVpcOrigin2020_05_31Cmd).Standalone()

	cloudfront_deleteVpcOrigin2020_05_31Cmd.Flags().String("id", "", "The VPC origin ID.")
	cloudfront_deleteVpcOrigin2020_05_31Cmd.Flags().String("if-match", "", "The version identifier of the VPC origin to delete.")
	cloudfront_deleteVpcOrigin2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_deleteVpcOrigin2020_05_31Cmd.MarkFlagRequired("if-match")
	cloudfrontCmd.AddCommand(cloudfront_deleteVpcOrigin2020_05_31Cmd)
}

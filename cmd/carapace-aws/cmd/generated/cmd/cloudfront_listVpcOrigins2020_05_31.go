package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listVpcOrigins2020_05_31Cmd = &cobra.Command{
	Use:   "list-vpc-origins2020_05_31",
	Short: "List the CloudFront VPC origins in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listVpcOrigins2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listVpcOrigins2020_05_31Cmd).Standalone()

		cloudfront_listVpcOrigins2020_05_31Cmd.Flags().String("marker", "", "The marker associated with the VPC origins list.")
		cloudfront_listVpcOrigins2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of items included in the list.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listVpcOrigins2020_05_31Cmd)
}

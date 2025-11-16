package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByWebAclid2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-web-aclid2020_05_31",
	Short: "List the distributions that are associated with a specified WAF web ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByWebAclid2020_05_31Cmd).Standalone()

	cloudfront_listDistributionsByWebAclid2020_05_31Cmd.Flags().String("marker", "", "Use `Marker` and `MaxItems` to control pagination of results.")
	cloudfront_listDistributionsByWebAclid2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distributions that you want CloudFront to return in the response body.")
	cloudfront_listDistributionsByWebAclid2020_05_31Cmd.Flags().String("web-aclid", "", "The ID of the WAF web ACL that you want to list the associated distributions.")
	cloudfront_listDistributionsByWebAclid2020_05_31Cmd.MarkFlagRequired("web-aclid")
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByWebAclid2020_05_31Cmd)
}

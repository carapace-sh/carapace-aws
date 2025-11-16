package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd = &cobra.Command{
	Use:   "list-distribution-tenants-by-customization2020_05_31",
	Short: "Lists distribution tenants by the customization that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd).Standalone()

		cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd.Flags().String("certificate-arn", "", "Filter by the ARN of the associated ACM certificate.")
		cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd.Flags().String("marker", "", "The marker for the next set of results.")
		cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distribution tenants to return by the specified customization.")
		cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd.Flags().String("web-aclarn", "", "Filter by the ARN of the associated WAF web ACL.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listDistributionTenantsByCustomization2020_05_31Cmd)
}

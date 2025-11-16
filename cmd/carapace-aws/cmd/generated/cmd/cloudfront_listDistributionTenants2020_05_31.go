package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionTenants2020_05_31Cmd = &cobra.Command{
	Use:   "list-distribution-tenants2020_05_31",
	Short: "Lists the distribution tenants in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionTenants2020_05_31Cmd).Standalone()

	cloudfront_listDistributionTenants2020_05_31Cmd.Flags().String("association-filter", "", "")
	cloudfront_listDistributionTenants2020_05_31Cmd.Flags().String("marker", "", "The marker for the next set of results.")
	cloudfront_listDistributionTenants2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distribution tenants to return.")
	cloudfrontCmd.AddCommand(cloudfront_listDistributionTenants2020_05_31Cmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_viewBillingCmd = &cobra.Command{
	Use:   "view-billing",
	Short: "Returns all the domain-related billing records for the current Amazon Web Services account for a specified period",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_viewBillingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_viewBillingCmd).Standalone()

		route53domains_viewBillingCmd.Flags().String("end", "", "The end date and time for the time period for which you want a list of billing records.")
		route53domains_viewBillingCmd.Flags().String("marker", "", "For an initial request for a list of billing records, omit this element.")
		route53domains_viewBillingCmd.Flags().String("max-items", "", "The number of billing records to be returned.")
		route53domains_viewBillingCmd.Flags().String("start", "", "The beginning date and time for the time period for which you want a list of billing records.")
	})
	route53domainsCmd.AddCommand(route53domains_viewBillingCmd)
}

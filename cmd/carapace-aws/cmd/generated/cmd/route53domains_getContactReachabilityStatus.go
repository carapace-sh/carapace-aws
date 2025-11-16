package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_getContactReachabilityStatusCmd = &cobra.Command{
	Use:   "get-contact-reachability-status",
	Short: "For operations that require confirmation that the email address for the registrant contact is valid, such as registering a new domain, this operation returns information about whether the registrant contact has responded.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_getContactReachabilityStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_getContactReachabilityStatusCmd).Standalone()

		route53domains_getContactReachabilityStatusCmd.Flags().String("domain-name", "", "The name of the domain for which you want to know whether the registrant contact has confirmed that the email address is valid.")
	})
	route53domainsCmd.AddCommand(route53domains_getContactReachabilityStatusCmd)
}

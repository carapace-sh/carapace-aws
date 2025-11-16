package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_resendContactReachabilityEmailCmd = &cobra.Command{
	Use:   "resend-contact-reachability-email",
	Short: "For operations that require confirmation that the email address for the registrant contact is valid, such as registering a new domain, this operation resends the confirmation email to the current email address for the registrant contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_resendContactReachabilityEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_resendContactReachabilityEmailCmd).Standalone()

		route53domains_resendContactReachabilityEmailCmd.Flags().String("domain-name", "", "The name of the domain for which you want Route 53 to resend a confirmation email to the registrant contact.")
	})
	route53domainsCmd.AddCommand(route53domains_resendContactReachabilityEmailCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_createLbcookieStickinessPolicyCmd = &cobra.Command{
	Use:   "create-lbcookie-stickiness-policy",
	Short: "Generates a stickiness policy with sticky session lifetimes controlled by the lifetime of the browser (user-agent) or a specified expiration period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_createLbcookieStickinessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_createLbcookieStickinessPolicyCmd).Standalone()

		elb_createLbcookieStickinessPolicyCmd.Flags().String("cookie-expiration-period", "", "The time period, in seconds, after which the cookie should be considered stale.")
		elb_createLbcookieStickinessPolicyCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_createLbcookieStickinessPolicyCmd.Flags().String("policy-name", "", "The name of the policy being created.")
		elb_createLbcookieStickinessPolicyCmd.MarkFlagRequired("load-balancer-name")
		elb_createLbcookieStickinessPolicyCmd.MarkFlagRequired("policy-name")
	})
	elbCmd.AddCommand(elb_createLbcookieStickinessPolicyCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_createAppCookieStickinessPolicyCmd = &cobra.Command{
	Use:   "create-app-cookie-stickiness-policy",
	Short: "Generates a stickiness policy with sticky session lifetimes that follow that of an application-generated cookie.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_createAppCookieStickinessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_createAppCookieStickinessPolicyCmd).Standalone()

		elb_createAppCookieStickinessPolicyCmd.Flags().String("cookie-name", "", "The name of the application cookie used for stickiness.")
		elb_createAppCookieStickinessPolicyCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_createAppCookieStickinessPolicyCmd.Flags().String("policy-name", "", "The name of the policy being created.")
		elb_createAppCookieStickinessPolicyCmd.MarkFlagRequired("cookie-name")
		elb_createAppCookieStickinessPolicyCmd.MarkFlagRequired("load-balancer-name")
		elb_createAppCookieStickinessPolicyCmd.MarkFlagRequired("policy-name")
	})
	elbCmd.AddCommand(elb_createAppCookieStickinessPolicyCmd)
}

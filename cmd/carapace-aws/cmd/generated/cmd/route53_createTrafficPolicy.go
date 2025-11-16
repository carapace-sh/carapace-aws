package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createTrafficPolicyCmd = &cobra.Command{
	Use:   "create-traffic-policy",
	Short: "Creates a traffic policy, which you use to create multiple DNS resource record sets for one domain name (such as example.com) or one subdomain name (such as www.example.com).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createTrafficPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_createTrafficPolicyCmd).Standalone()

		route53_createTrafficPolicyCmd.Flags().String("comment", "", "(Optional) Any comments that you want to include about the traffic policy.")
		route53_createTrafficPolicyCmd.Flags().String("document", "", "The definition of this traffic policy in JSON format.")
		route53_createTrafficPolicyCmd.Flags().String("name", "", "The name of the traffic policy.")
		route53_createTrafficPolicyCmd.MarkFlagRequired("document")
		route53_createTrafficPolicyCmd.MarkFlagRequired("name")
	})
	route53Cmd.AddCommand(route53_createTrafficPolicyCmd)
}

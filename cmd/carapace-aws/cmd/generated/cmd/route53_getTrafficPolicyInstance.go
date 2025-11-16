package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getTrafficPolicyInstanceCmd = &cobra.Command{
	Use:   "get-traffic-policy-instance",
	Short: "Gets information about a specified traffic policy instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getTrafficPolicyInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getTrafficPolicyInstanceCmd).Standalone()

		route53_getTrafficPolicyInstanceCmd.Flags().String("id", "", "The ID of the traffic policy instance that you want to get information about.")
		route53_getTrafficPolicyInstanceCmd.MarkFlagRequired("id")
	})
	route53Cmd.AddCommand(route53_getTrafficPolicyInstanceCmd)
}

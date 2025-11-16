package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getTrafficPolicyCmd = &cobra.Command{
	Use:   "get-traffic-policy",
	Short: "Gets information about a specific traffic policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getTrafficPolicyCmd).Standalone()

	route53_getTrafficPolicyCmd.Flags().String("id", "", "The ID of the traffic policy that you want to get information about.")
	route53_getTrafficPolicyCmd.Flags().String("version", "", "The version number of the traffic policy that you want to get information about.")
	route53_getTrafficPolicyCmd.MarkFlagRequired("id")
	route53_getTrafficPolicyCmd.MarkFlagRequired("version")
	route53Cmd.AddCommand(route53_getTrafficPolicyCmd)
}

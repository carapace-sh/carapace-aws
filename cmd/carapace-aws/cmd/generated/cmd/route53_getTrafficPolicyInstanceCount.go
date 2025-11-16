package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getTrafficPolicyInstanceCountCmd = &cobra.Command{
	Use:   "get-traffic-policy-instance-count",
	Short: "Gets the number of traffic policy instances that are associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getTrafficPolicyInstanceCountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getTrafficPolicyInstanceCountCmd).Standalone()

	})
	route53Cmd.AddCommand(route53_getTrafficPolicyInstanceCountCmd)
}

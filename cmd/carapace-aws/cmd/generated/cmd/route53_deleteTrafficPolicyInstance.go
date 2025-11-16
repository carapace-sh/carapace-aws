package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteTrafficPolicyInstanceCmd = &cobra.Command{
	Use:   "delete-traffic-policy-instance",
	Short: "Deletes a traffic policy instance and all of the resource record sets that Amazon Route 53 created when you created the instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteTrafficPolicyInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_deleteTrafficPolicyInstanceCmd).Standalone()

		route53_deleteTrafficPolicyInstanceCmd.Flags().String("id", "", "The ID of the traffic policy instance that you want to delete.")
		route53_deleteTrafficPolicyInstanceCmd.MarkFlagRequired("id")
	})
	route53Cmd.AddCommand(route53_deleteTrafficPolicyInstanceCmd)
}

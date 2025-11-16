package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteTrafficPolicyCmd = &cobra.Command{
	Use:   "delete-traffic-policy",
	Short: "Deletes a traffic policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteTrafficPolicyCmd).Standalone()

	route53_deleteTrafficPolicyCmd.Flags().String("id", "", "The ID of the traffic policy that you want to delete.")
	route53_deleteTrafficPolicyCmd.Flags().String("version", "", "The version number of the traffic policy that you want to delete.")
	route53_deleteTrafficPolicyCmd.MarkFlagRequired("id")
	route53_deleteTrafficPolicyCmd.MarkFlagRequired("version")
	route53Cmd.AddCommand(route53_deleteTrafficPolicyCmd)
}

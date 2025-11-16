package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createTrafficPolicyVersionCmd = &cobra.Command{
	Use:   "create-traffic-policy-version",
	Short: "Creates a new version of an existing traffic policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createTrafficPolicyVersionCmd).Standalone()

	route53_createTrafficPolicyVersionCmd.Flags().String("comment", "", "The comment that you specified in the `CreateTrafficPolicyVersion` request, if any.")
	route53_createTrafficPolicyVersionCmd.Flags().String("document", "", "The definition of this version of the traffic policy, in JSON format.")
	route53_createTrafficPolicyVersionCmd.Flags().String("id", "", "The ID of the traffic policy for which you want to create a new version.")
	route53_createTrafficPolicyVersionCmd.MarkFlagRequired("document")
	route53_createTrafficPolicyVersionCmd.MarkFlagRequired("id")
	route53Cmd.AddCommand(route53_createTrafficPolicyVersionCmd)
}

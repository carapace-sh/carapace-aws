package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeSslpoliciesCmd = &cobra.Command{
	Use:   "describe-sslpolicies",
	Short: "Describes the specified policies or all policies used for SSL negotiation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeSslpoliciesCmd).Standalone()

	elbv2_describeSslpoliciesCmd.Flags().String("load-balancer-type", "", "The type of load balancer.")
	elbv2_describeSslpoliciesCmd.Flags().String("marker", "", "The marker for the next set of results.")
	elbv2_describeSslpoliciesCmd.Flags().String("names", "", "The names of the policies.")
	elbv2_describeSslpoliciesCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	elbv2Cmd.AddCommand(elbv2_describeSslpoliciesCmd)
}

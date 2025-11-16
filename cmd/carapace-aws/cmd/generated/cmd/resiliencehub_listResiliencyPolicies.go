package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listResiliencyPoliciesCmd = &cobra.Command{
	Use:   "list-resiliency-policies",
	Short: "Lists the resiliency policies for the Resilience Hub applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listResiliencyPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listResiliencyPoliciesCmd).Standalone()

		resiliencehub_listResiliencyPoliciesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
		resiliencehub_listResiliencyPoliciesCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listResiliencyPoliciesCmd.Flags().String("policy-name", "", "Name of the resiliency policy.")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listResiliencyPoliciesCmd)
}

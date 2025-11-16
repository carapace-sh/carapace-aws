package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listSuggestedResiliencyPoliciesCmd = &cobra.Command{
	Use:   "list-suggested-resiliency-policies",
	Short: "Lists the suggested resiliency policies for the Resilience Hub applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listSuggestedResiliencyPoliciesCmd).Standalone()

	resiliencehub_listSuggestedResiliencyPoliciesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	resiliencehub_listSuggestedResiliencyPoliciesCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehubCmd.AddCommand(resiliencehub_listSuggestedResiliencyPoliciesCmd)
}

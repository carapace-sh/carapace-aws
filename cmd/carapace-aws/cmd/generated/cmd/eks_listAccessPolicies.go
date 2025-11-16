package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listAccessPoliciesCmd = &cobra.Command{
	Use:   "list-access-policies",
	Short: "Lists the available access policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listAccessPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_listAccessPoliciesCmd).Standalone()

		eks_listAccessPoliciesCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
		eks_listAccessPoliciesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
	})
	eksCmd.AddCommand(eks_listAccessPoliciesCmd)
}

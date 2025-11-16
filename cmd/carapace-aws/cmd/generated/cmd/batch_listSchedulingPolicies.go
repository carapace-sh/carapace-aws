package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_listSchedulingPoliciesCmd = &cobra.Command{
	Use:   "list-scheduling-policies",
	Short: "Returns a list of Batch scheduling policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_listSchedulingPoliciesCmd).Standalone()

	batch_listSchedulingPoliciesCmd.Flags().String("max-results", "", "The maximum number of results that's returned by `ListSchedulingPolicies` in paginated output.")
	batch_listSchedulingPoliciesCmd.Flags().String("next-token", "", "The `nextToken` value that's returned from a previous paginated `ListSchedulingPolicies` request where `maxResults` was used and the results exceeded the value of that parameter.")
	batchCmd.AddCommand(batch_listSchedulingPoliciesCmd)
}

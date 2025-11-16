package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listAccessEntriesCmd = &cobra.Command{
	Use:   "list-access-entries",
	Short: "Lists the access entries for your cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listAccessEntriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_listAccessEntriesCmd).Standalone()

		eks_listAccessEntriesCmd.Flags().String("associated-policy-arn", "", "The ARN of an `AccessPolicy`.")
		eks_listAccessEntriesCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_listAccessEntriesCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
		eks_listAccessEntriesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
		eks_listAccessEntriesCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_listAccessEntriesCmd)
}

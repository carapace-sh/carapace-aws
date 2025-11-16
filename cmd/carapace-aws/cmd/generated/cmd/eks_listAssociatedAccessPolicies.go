package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listAssociatedAccessPoliciesCmd = &cobra.Command{
	Use:   "list-associated-access-policies",
	Short: "Lists the access policies associated with an access entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listAssociatedAccessPoliciesCmd).Standalone()

	eks_listAssociatedAccessPoliciesCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_listAssociatedAccessPoliciesCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
	eks_listAssociatedAccessPoliciesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
	eks_listAssociatedAccessPoliciesCmd.Flags().String("principal-arn", "", "The ARN of the IAM principal for the `AccessEntry`.")
	eks_listAssociatedAccessPoliciesCmd.MarkFlagRequired("cluster-name")
	eks_listAssociatedAccessPoliciesCmd.MarkFlagRequired("principal-arn")
	eksCmd.AddCommand(eks_listAssociatedAccessPoliciesCmd)
}

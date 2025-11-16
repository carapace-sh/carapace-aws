package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_listPolicyVersionsCmd = &cobra.Command{
	Use:   "list-policy-versions",
	Short: "Returns a list of the versions for policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_listPolicyVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_listPolicyVersionsCmd).Standalone()

		mpa_listPolicyVersionsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		mpa_listPolicyVersionsCmd.Flags().String("next-token", "", "If present, indicates that more output is available than is included in the current response.")
		mpa_listPolicyVersionsCmd.Flags().String("policy-arn", "", "Amazon Resource Name (ARN) for the policy.")
		mpa_listPolicyVersionsCmd.MarkFlagRequired("policy-arn")
	})
	mpaCmd.AddCommand(mpa_listPolicyVersionsCmd)
}

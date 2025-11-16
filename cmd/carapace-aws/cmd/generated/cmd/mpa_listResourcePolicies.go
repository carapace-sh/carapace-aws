package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_listResourcePoliciesCmd = &cobra.Command{
	Use:   "list-resource-policies",
	Short: "Returns a list of policies for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_listResourcePoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_listResourcePoliciesCmd).Standalone()

		mpa_listResourcePoliciesCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		mpa_listResourcePoliciesCmd.Flags().String("next-token", "", "If present, indicates that more output is available than is included in the current response.")
		mpa_listResourcePoliciesCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) for the resource.")
		mpa_listResourcePoliciesCmd.MarkFlagRequired("resource-arn")
	})
	mpaCmd.AddCommand(mpa_listResourcePoliciesCmd)
}

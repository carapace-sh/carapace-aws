package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getResourcePoliciesCmd = &cobra.Command{
	Use:   "get-resource-policies",
	Short: "Returns an array of the `Policy` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getResourcePoliciesCmd).Standalone()

	ssm_getResourcePoliciesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_getResourcePoliciesCmd.Flags().String("next-token", "", "A token to start the list.")
	ssm_getResourcePoliciesCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource to which the policies are attached.")
	ssm_getResourcePoliciesCmd.MarkFlagRequired("resource-arn")
	ssmCmd.AddCommand(ssm_getResourcePoliciesCmd)
}

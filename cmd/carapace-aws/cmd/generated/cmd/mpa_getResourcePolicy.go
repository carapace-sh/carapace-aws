package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Returns details about a policy for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_getResourcePolicyCmd).Standalone()

		mpa_getResourcePolicyCmd.Flags().String("policy-name", "", "Name of the policy.")
		mpa_getResourcePolicyCmd.Flags().String("policy-type", "", "The type of policy.")
		mpa_getResourcePolicyCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) for the resource.")
		mpa_getResourcePolicyCmd.MarkFlagRequired("policy-name")
		mpa_getResourcePolicyCmd.MarkFlagRequired("policy-type")
		mpa_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	mpaCmd.AddCommand(mpa_getResourcePolicyCmd)
}

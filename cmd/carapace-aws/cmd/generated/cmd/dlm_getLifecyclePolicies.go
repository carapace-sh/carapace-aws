package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlm_getLifecyclePoliciesCmd = &cobra.Command{
	Use:   "get-lifecycle-policies",
	Short: "Gets summary information about all or the specified data lifecycle policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlm_getLifecyclePoliciesCmd).Standalone()

	dlm_getLifecyclePoliciesCmd.Flags().String("default-policy-type", "", "**\\[Default policies only]** Specifies the type of default policy to get.")
	dlm_getLifecyclePoliciesCmd.Flags().String("policy-ids", "", "The identifiers of the data lifecycle policies.")
	dlm_getLifecyclePoliciesCmd.Flags().String("resource-types", "", "The resource type.")
	dlm_getLifecyclePoliciesCmd.Flags().String("state", "", "The activation state.")
	dlm_getLifecyclePoliciesCmd.Flags().String("tags-to-add", "", "The tags to add to objects created by the policy.")
	dlm_getLifecyclePoliciesCmd.Flags().String("target-tags", "", "The target tag for a policy.")
	dlmCmd.AddCommand(dlm_getLifecyclePoliciesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_updateLifecyclePolicyCmd = &cobra.Command{
	Use:   "update-lifecycle-policy",
	Short: "Update the specified lifecycle policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_updateLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_updateLifecyclePolicyCmd).Standalone()

		imagebuilder_updateLifecyclePolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_updateLifecyclePolicyCmd.Flags().String("description", "", "Optional description for the lifecycle policy.")
		imagebuilder_updateLifecyclePolicyCmd.Flags().String("execution-role", "", "The name or Amazon Resource Name (ARN) of the IAM role that Image Builder uses to update the lifecycle policy.")
		imagebuilder_updateLifecyclePolicyCmd.Flags().String("lifecycle-policy-arn", "", "The Amazon Resource Name (ARN) of the lifecycle policy resource.")
		imagebuilder_updateLifecyclePolicyCmd.Flags().String("policy-details", "", "The configuration details for a lifecycle policy resource.")
		imagebuilder_updateLifecyclePolicyCmd.Flags().String("resource-selection", "", "Selection criteria for resources that the lifecycle policy applies to.")
		imagebuilder_updateLifecyclePolicyCmd.Flags().String("resource-type", "", "The type of image resource that the lifecycle policy applies to.")
		imagebuilder_updateLifecyclePolicyCmd.Flags().String("status", "", "Indicates whether the lifecycle policy resource is enabled.")
		imagebuilder_updateLifecyclePolicyCmd.MarkFlagRequired("client-token")
		imagebuilder_updateLifecyclePolicyCmd.MarkFlagRequired("execution-role")
		imagebuilder_updateLifecyclePolicyCmd.MarkFlagRequired("lifecycle-policy-arn")
		imagebuilder_updateLifecyclePolicyCmd.MarkFlagRequired("policy-details")
		imagebuilder_updateLifecyclePolicyCmd.MarkFlagRequired("resource-selection")
		imagebuilder_updateLifecyclePolicyCmd.MarkFlagRequired("resource-type")
	})
	imagebuilderCmd.AddCommand(imagebuilder_updateLifecyclePolicyCmd)
}

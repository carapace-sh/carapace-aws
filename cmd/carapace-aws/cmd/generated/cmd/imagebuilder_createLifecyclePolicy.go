package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createLifecyclePolicyCmd = &cobra.Command{
	Use:   "create-lifecycle-policy",
	Short: "Create a lifecycle policy resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_createLifecyclePolicyCmd).Standalone()

		imagebuilder_createLifecyclePolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("description", "", "Optional description for the lifecycle policy.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("execution-role", "", "The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to run lifecycle actions.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("name", "", "The name of the lifecycle policy to create.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("policy-details", "", "Configuration details for the lifecycle policy rules.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("resource-selection", "", "Selection criteria for the resources that the lifecycle policy applies to.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("resource-type", "", "The type of Image Builder resource that the lifecycle policy applies to.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("status", "", "Indicates whether the lifecycle policy resource is enabled.")
		imagebuilder_createLifecyclePolicyCmd.Flags().String("tags", "", "Tags to apply to the lifecycle policy resource.")
		imagebuilder_createLifecyclePolicyCmd.MarkFlagRequired("client-token")
		imagebuilder_createLifecyclePolicyCmd.MarkFlagRequired("execution-role")
		imagebuilder_createLifecyclePolicyCmd.MarkFlagRequired("name")
		imagebuilder_createLifecyclePolicyCmd.MarkFlagRequired("policy-details")
		imagebuilder_createLifecyclePolicyCmd.MarkFlagRequired("resource-selection")
		imagebuilder_createLifecyclePolicyCmd.MarkFlagRequired("resource-type")
	})
	imagebuilderCmd.AddCommand(imagebuilder_createLifecyclePolicyCmd)
}

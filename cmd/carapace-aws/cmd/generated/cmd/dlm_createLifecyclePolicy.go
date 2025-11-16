package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlm_createLifecyclePolicyCmd = &cobra.Command{
	Use:   "create-lifecycle-policy",
	Short: "Creates an Amazon Data Lifecycle Manager lifecycle policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlm_createLifecyclePolicyCmd).Standalone()

	dlm_createLifecyclePolicyCmd.Flags().String("copy-tags", "", "**\\[Default policies only]** Indicates whether the policy should copy tags from the source resource to the snapshot or AMI.")
	dlm_createLifecyclePolicyCmd.Flags().String("create-interval", "", "**\\[Default policies only]** Specifies how often the policy should run and create snapshots or AMIs.")
	dlm_createLifecyclePolicyCmd.Flags().String("cross-region-copy-targets", "", "**\\[Default policies only]** Specifies destination Regions for snapshot or AMI copies.")
	dlm_createLifecyclePolicyCmd.Flags().String("default-policy", "", "**\\[Default policies only]** Specify the type of default policy to create.")
	dlm_createLifecyclePolicyCmd.Flags().String("description", "", "A description of the lifecycle policy.")
	dlm_createLifecyclePolicyCmd.Flags().String("exclusions", "", "**\\[Default policies only]** Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs.")
	dlm_createLifecyclePolicyCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.")
	dlm_createLifecyclePolicyCmd.Flags().String("extend-deletion", "", "**\\[Default policies only]** Defines the snapshot or AMI retention behavior for the policy if the source volume or instance is deleted, or if the policy enters the error, disabled, or deleted state.")
	dlm_createLifecyclePolicyCmd.Flags().String("policy-details", "", "The configuration details of the lifecycle policy.")
	dlm_createLifecyclePolicyCmd.Flags().String("retain-interval", "", "**\\[Default policies only]** Specifies how long the policy should retain snapshots or AMIs before deleting them.")
	dlm_createLifecyclePolicyCmd.Flags().String("state", "", "The activation state of the lifecycle policy after creation.")
	dlm_createLifecyclePolicyCmd.Flags().String("tags", "", "The tags to apply to the lifecycle policy during creation.")
	dlm_createLifecyclePolicyCmd.MarkFlagRequired("description")
	dlm_createLifecyclePolicyCmd.MarkFlagRequired("execution-role-arn")
	dlm_createLifecyclePolicyCmd.MarkFlagRequired("state")
	dlmCmd.AddCommand(dlm_createLifecyclePolicyCmd)
}

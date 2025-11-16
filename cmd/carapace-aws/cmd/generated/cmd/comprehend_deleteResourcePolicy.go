package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource-based policy that is attached to a custom model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_deleteResourcePolicyCmd).Standalone()

		comprehend_deleteResourcePolicyCmd.Flags().String("policy-revision-id", "", "The revision ID of the policy to delete.")
		comprehend_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the custom model version that has the policy to delete.")
		comprehend_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	comprehendCmd.AddCommand(comprehend_deleteResourcePolicyCmd)
}

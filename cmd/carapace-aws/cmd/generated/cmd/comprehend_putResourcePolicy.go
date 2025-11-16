package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based policy to a custom model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_putResourcePolicyCmd).Standalone()

		comprehend_putResourcePolicyCmd.Flags().String("policy-revision-id", "", "The revision ID that Amazon Comprehend assigned to the policy that you are updating.")
		comprehend_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the custom model to attach the policy to.")
		comprehend_putResourcePolicyCmd.Flags().String("resource-policy", "", "The JSON resource-based policy to attach to your custom model.")
		comprehend_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
		comprehend_putResourcePolicyCmd.MarkFlagRequired("resource-policy")
	})
	comprehendCmd.AddCommand(comprehend_putResourcePolicyCmd)
}

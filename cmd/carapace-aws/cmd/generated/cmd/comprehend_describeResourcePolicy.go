package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeResourcePolicyCmd = &cobra.Command{
	Use:   "describe-resource-policy",
	Short: "Gets the details of a resource-based policy that is attached to a custom model, including the JSON body of the policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeResourcePolicyCmd).Standalone()

	comprehend_describeResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the custom model version that has the resource policy.")
	comprehend_describeResourcePolicyCmd.MarkFlagRequired("resource-arn")
	comprehendCmd.AddCommand(comprehend_describeResourcePolicyCmd)
}

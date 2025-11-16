package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Returns the resource-based policy document attached to the resource in `JSON` format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_getResourcePolicyCmd).Standalone()

	billing_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the billing view resource to which the policy is attached to.")
	billing_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	billingCmd.AddCommand(billing_getResourcePolicyCmd)
}

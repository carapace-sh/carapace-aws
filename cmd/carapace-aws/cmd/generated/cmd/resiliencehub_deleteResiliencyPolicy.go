package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_deleteResiliencyPolicyCmd = &cobra.Command{
	Use:   "delete-resiliency-policy",
	Short: "Deletes a resiliency policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_deleteResiliencyPolicyCmd).Standalone()

	resiliencehub_deleteResiliencyPolicyCmd.Flags().String("client-token", "", "Used for an idempotency token.")
	resiliencehub_deleteResiliencyPolicyCmd.Flags().String("policy-arn", "", "Amazon Resource Name (ARN) of the resiliency policy.")
	resiliencehub_deleteResiliencyPolicyCmd.MarkFlagRequired("policy-arn")
	resiliencehubCmd.AddCommand(resiliencehub_deleteResiliencyPolicyCmd)
}

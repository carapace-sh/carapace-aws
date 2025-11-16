package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeResiliencyPolicyCmd = &cobra.Command{
	Use:   "describe-resiliency-policy",
	Short: "Describes a specified resiliency policy for an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeResiliencyPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_describeResiliencyPolicyCmd).Standalone()

		resiliencehub_describeResiliencyPolicyCmd.Flags().String("policy-arn", "", "Amazon Resource Name (ARN) of the resiliency policy.")
		resiliencehub_describeResiliencyPolicyCmd.MarkFlagRequired("policy-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_describeResiliencyPolicyCmd)
}

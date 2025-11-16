package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_putPolicyCmd = &cobra.Command{
	Use:   "put-policy",
	Short: "Updates the resource-based policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_putPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_putPolicyCmd).Standalone()

		entityresolution_putPolicyCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource for which the policy needs to be updated.")
		entityresolution_putPolicyCmd.Flags().String("policy", "", "The resource-based policy.")
		entityresolution_putPolicyCmd.Flags().String("token", "", "A unique identifier for the current revision of the policy.")
		entityresolution_putPolicyCmd.MarkFlagRequired("arn")
		entityresolution_putPolicyCmd.MarkFlagRequired("policy")
	})
	entityresolutionCmd.AddCommand(entityresolution_putPolicyCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryControlConfig_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Get information about the resource policy for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryControlConfig_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryControlConfig_getResourcePolicyCmd).Standalone()

		route53RecoveryControlConfig_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		route53RecoveryControlConfig_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	route53RecoveryControlConfigCmd.AddCommand(route53RecoveryControlConfig_getResourcePolicyCmd)
}

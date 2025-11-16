package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_detachPolicyCmd = &cobra.Command{
	Use:   "detach-policy",
	Short: "Detaches a policy from the specified target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_detachPolicyCmd).Standalone()

	iot_detachPolicyCmd.Flags().String("policy-name", "", "The policy to detach.")
	iot_detachPolicyCmd.Flags().String("target", "", "The target from which the policy will be detached.")
	iot_detachPolicyCmd.MarkFlagRequired("policy-name")
	iot_detachPolicyCmd.MarkFlagRequired("target")
	iotCmd.AddCommand(iot_detachPolicyCmd)
}

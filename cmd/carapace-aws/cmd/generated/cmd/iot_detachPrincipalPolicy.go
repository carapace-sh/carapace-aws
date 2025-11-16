package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_detachPrincipalPolicyCmd = &cobra.Command{
	Use:   "detach-principal-policy",
	Short: "Removes the specified policy from the specified certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_detachPrincipalPolicyCmd).Standalone()

	iot_detachPrincipalPolicyCmd.Flags().String("policy-name", "", "The name of the policy to detach.")
	iot_detachPrincipalPolicyCmd.Flags().String("principal", "", "The principal.")
	iot_detachPrincipalPolicyCmd.MarkFlagRequired("policy-name")
	iot_detachPrincipalPolicyCmd.MarkFlagRequired("principal")
	iotCmd.AddCommand(iot_detachPrincipalPolicyCmd)
}

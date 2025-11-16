package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getEffectivePoliciesCmd = &cobra.Command{
	Use:   "get-effective-policies",
	Short: "Gets a list of the policies that have an effect on the authorization behavior of the specified device when it connects to the IoT device gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getEffectivePoliciesCmd).Standalone()

	iot_getEffectivePoliciesCmd.Flags().String("cognito-identity-pool-id", "", "The Cognito identity pool ID.")
	iot_getEffectivePoliciesCmd.Flags().String("principal", "", "The principal.")
	iot_getEffectivePoliciesCmd.Flags().String("thing-name", "", "The thing name.")
	iotCmd.AddCommand(iot_getEffectivePoliciesCmd)
}

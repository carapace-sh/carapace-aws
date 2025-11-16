package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_setDefaultPolicyVersionCmd = &cobra.Command{
	Use:   "set-default-policy-version",
	Short: "Sets the specified version of the specified policy as the policy's default (operative) version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_setDefaultPolicyVersionCmd).Standalone()

	iot_setDefaultPolicyVersionCmd.Flags().String("policy-name", "", "The policy name.")
	iot_setDefaultPolicyVersionCmd.Flags().String("policy-version-id", "", "The policy version ID.")
	iot_setDefaultPolicyVersionCmd.MarkFlagRequired("policy-name")
	iot_setDefaultPolicyVersionCmd.MarkFlagRequired("policy-version-id")
	iotCmd.AddCommand(iot_setDefaultPolicyVersionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deletePolicyVersionCmd = &cobra.Command{
	Use:   "delete-policy-version",
	Short: "Deletes the specified version of the specified policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deletePolicyVersionCmd).Standalone()

	iot_deletePolicyVersionCmd.Flags().String("policy-name", "", "The name of the policy.")
	iot_deletePolicyVersionCmd.Flags().String("policy-version-id", "", "The policy version ID.")
	iot_deletePolicyVersionCmd.MarkFlagRequired("policy-name")
	iot_deletePolicyVersionCmd.MarkFlagRequired("policy-version-id")
	iotCmd.AddCommand(iot_deletePolicyVersionCmd)
}

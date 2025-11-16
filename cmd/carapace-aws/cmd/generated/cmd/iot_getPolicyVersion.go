package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getPolicyVersionCmd = &cobra.Command{
	Use:   "get-policy-version",
	Short: "Gets information about the specified policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getPolicyVersionCmd).Standalone()

	iot_getPolicyVersionCmd.Flags().String("policy-name", "", "The name of the policy.")
	iot_getPolicyVersionCmd.Flags().String("policy-version-id", "", "The policy version ID.")
	iot_getPolicyVersionCmd.MarkFlagRequired("policy-name")
	iot_getPolicyVersionCmd.MarkFlagRequired("policy-version-id")
	iotCmd.AddCommand(iot_getPolicyVersionCmd)
}

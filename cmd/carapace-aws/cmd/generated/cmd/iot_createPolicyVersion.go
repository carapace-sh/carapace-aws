package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createPolicyVersionCmd = &cobra.Command{
	Use:   "create-policy-version",
	Short: "Creates a new version of the specified IoT policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createPolicyVersionCmd).Standalone()

	iot_createPolicyVersionCmd.Flags().String("policy-document", "", "The JSON document that describes the policy.")
	iot_createPolicyVersionCmd.Flags().String("policy-name", "", "The policy name.")
	iot_createPolicyVersionCmd.Flags().String("set-as-default", "", "Specifies whether the policy version is set as the default.")
	iot_createPolicyVersionCmd.MarkFlagRequired("policy-document")
	iot_createPolicyVersionCmd.MarkFlagRequired("policy-name")
	iotCmd.AddCommand(iot_createPolicyVersionCmd)
}

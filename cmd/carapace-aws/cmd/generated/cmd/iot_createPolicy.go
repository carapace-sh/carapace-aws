package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createPolicyCmd = &cobra.Command{
	Use:   "create-policy",
	Short: "Creates an IoT policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createPolicyCmd).Standalone()

		iot_createPolicyCmd.Flags().String("policy-document", "", "The JSON document that describes the policy.")
		iot_createPolicyCmd.Flags().String("policy-name", "", "The policy name.")
		iot_createPolicyCmd.Flags().String("tags", "", "Metadata which can be used to manage the policy.")
		iot_createPolicyCmd.MarkFlagRequired("policy-document")
		iot_createPolicyCmd.MarkFlagRequired("policy-name")
	})
	iotCmd.AddCommand(iot_createPolicyCmd)
}

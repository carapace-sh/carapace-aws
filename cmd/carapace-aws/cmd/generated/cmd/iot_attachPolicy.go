package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_attachPolicyCmd = &cobra.Command{
	Use:   "attach-policy",
	Short: "Attaches the specified policy to the specified principal (certificate or other credential).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_attachPolicyCmd).Standalone()

	iot_attachPolicyCmd.Flags().String("policy-name", "", "The name of the policy to attach.")
	iot_attachPolicyCmd.Flags().String("target", "", "The [identity](https://docs.aws.amazon.com/iot/latest/developerguide/security-iam.html) to which the policy is attached.")
	iot_attachPolicyCmd.MarkFlagRequired("policy-name")
	iot_attachPolicyCmd.MarkFlagRequired("target")
	iotCmd.AddCommand(iot_attachPolicyCmd)
}

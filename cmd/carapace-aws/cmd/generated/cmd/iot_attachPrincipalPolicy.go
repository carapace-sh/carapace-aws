package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_attachPrincipalPolicyCmd = &cobra.Command{
	Use:   "attach-principal-policy",
	Short: "Attaches the specified policy to the specified principal (certificate or other credential).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_attachPrincipalPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_attachPrincipalPolicyCmd).Standalone()

		iot_attachPrincipalPolicyCmd.Flags().String("policy-name", "", "The policy name.")
		iot_attachPrincipalPolicyCmd.Flags().String("principal", "", "The principal, which can be a certificate ARN (as returned from the CreateCertificate operation) or an Amazon Cognito ID.")
		iot_attachPrincipalPolicyCmd.MarkFlagRequired("policy-name")
		iot_attachPrincipalPolicyCmd.MarkFlagRequired("principal")
	})
	iotCmd.AddCommand(iot_attachPrincipalPolicyCmd)
}

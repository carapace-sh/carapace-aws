package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_testAuthorizationCmd = &cobra.Command{
	Use:   "test-authorization",
	Short: "Tests if a specified principal is authorized to perform an IoT action on a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_testAuthorizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_testAuthorizationCmd).Standalone()

		iot_testAuthorizationCmd.Flags().String("auth-infos", "", "A list of authorization info objects.")
		iot_testAuthorizationCmd.Flags().String("client-id", "", "The MQTT client ID.")
		iot_testAuthorizationCmd.Flags().String("cognito-identity-pool-id", "", "The Cognito identity pool ID.")
		iot_testAuthorizationCmd.Flags().String("policy-names-to-add", "", "When testing custom authorization, the policies specified here are treated as if they are attached to the principal being authorized.")
		iot_testAuthorizationCmd.Flags().String("policy-names-to-skip", "", "When testing custom authorization, the policies specified here are treated as if they are not attached to the principal being authorized.")
		iot_testAuthorizationCmd.Flags().String("principal", "", "The principal.")
		iot_testAuthorizationCmd.MarkFlagRequired("auth-infos")
	})
	iotCmd.AddCommand(iot_testAuthorizationCmd)
}

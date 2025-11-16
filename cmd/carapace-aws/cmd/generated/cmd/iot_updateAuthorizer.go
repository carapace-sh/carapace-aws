package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateAuthorizerCmd = &cobra.Command{
	Use:   "update-authorizer",
	Short: "Updates an authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateAuthorizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateAuthorizerCmd).Standalone()

		iot_updateAuthorizerCmd.Flags().String("authorizer-function-arn", "", "The ARN of the authorizer's Lambda function.")
		iot_updateAuthorizerCmd.Flags().String("authorizer-name", "", "The authorizer name.")
		iot_updateAuthorizerCmd.Flags().String("enable-caching-for-http", "", "When `true`, the result from the authorizer’s Lambda function is cached for the time specified in `refreshAfterInSeconds`.")
		iot_updateAuthorizerCmd.Flags().String("status", "", "The status of the update authorizer request.")
		iot_updateAuthorizerCmd.Flags().String("token-key-name", "", "The key used to extract the token from the HTTP headers.")
		iot_updateAuthorizerCmd.Flags().String("token-signing-public-keys", "", "The public keys used to verify the token signature.")
		iot_updateAuthorizerCmd.MarkFlagRequired("authorizer-name")
	})
	iotCmd.AddCommand(iot_updateAuthorizerCmd)
}

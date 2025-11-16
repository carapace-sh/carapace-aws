package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createAuthorizerCmd = &cobra.Command{
	Use:   "create-authorizer",
	Short: "Creates an authorizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createAuthorizerCmd).Standalone()

	iot_createAuthorizerCmd.Flags().String("authorizer-function-arn", "", "The ARN of the authorizer's Lambda function.")
	iot_createAuthorizerCmd.Flags().String("authorizer-name", "", "The authorizer name.")
	iot_createAuthorizerCmd.Flags().String("enable-caching-for-http", "", "When `true`, the result from the authorizer’s Lambda function is cached for clients that use persistent HTTP connections.")
	iot_createAuthorizerCmd.Flags().String("signing-disabled", "", "Specifies whether IoT validates the token signature in an authorization request.")
	iot_createAuthorizerCmd.Flags().String("status", "", "The status of the create authorizer request.")
	iot_createAuthorizerCmd.Flags().String("tags", "", "Metadata which can be used to manage the custom authorizer.")
	iot_createAuthorizerCmd.Flags().String("token-key-name", "", "The name of the token key used to extract the token from the HTTP headers.")
	iot_createAuthorizerCmd.Flags().String("token-signing-public-keys", "", "The public keys used to verify the digital signature returned by your custom authentication service.")
	iot_createAuthorizerCmd.MarkFlagRequired("authorizer-function-arn")
	iot_createAuthorizerCmd.MarkFlagRequired("authorizer-name")
	iotCmd.AddCommand(iot_createAuthorizerCmd)
}

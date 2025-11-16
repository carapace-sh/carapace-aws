package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_registerAppInstanceUserEndpointCmd = &cobra.Command{
	Use:   "register-app-instance-user-endpoint",
	Short: "Registers an endpoint under an Amazon Chime `AppInstanceUser`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_registerAppInstanceUserEndpointCmd).Standalone()

	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.Flags().String("allow-messages", "", "Boolean that controls whether the AppInstanceUserEndpoint is opted in to receive messages.")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.Flags().String("client-request-token", "", "The unique ID assigned to the request.")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.Flags().String("endpoint-attributes", "", "The attributes of an `Endpoint`.")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.Flags().String("name", "", "The name of the `AppInstanceUserEndpoint`.")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.Flags().String("resource-arn", "", "The ARN of the resource to which the endpoint belongs.")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.Flags().String("type", "", "The type of the `AppInstanceUserEndpoint`.")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.MarkFlagRequired("app-instance-user-arn")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.MarkFlagRequired("client-request-token")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.MarkFlagRequired("endpoint-attributes")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.MarkFlagRequired("resource-arn")
	chimeSdkIdentity_registerAppInstanceUserEndpointCmd.MarkFlagRequired("type")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_registerAppInstanceUserEndpointCmd)
}

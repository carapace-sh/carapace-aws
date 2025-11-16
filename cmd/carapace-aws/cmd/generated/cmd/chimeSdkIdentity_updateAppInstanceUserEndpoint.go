package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_updateAppInstanceUserEndpointCmd = &cobra.Command{
	Use:   "update-app-instance-user-endpoint",
	Short: "Updates the details of an `AppInstanceUserEndpoint`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_updateAppInstanceUserEndpointCmd).Standalone()

	chimeSdkIdentity_updateAppInstanceUserEndpointCmd.Flags().String("allow-messages", "", "Boolean that controls whether the `AppInstanceUserEndpoint` is opted in to receive messages.")
	chimeSdkIdentity_updateAppInstanceUserEndpointCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
	chimeSdkIdentity_updateAppInstanceUserEndpointCmd.Flags().String("endpoint-id", "", "The unique identifier of the `AppInstanceUserEndpoint`.")
	chimeSdkIdentity_updateAppInstanceUserEndpointCmd.Flags().String("name", "", "The name of the `AppInstanceUserEndpoint`.")
	chimeSdkIdentity_updateAppInstanceUserEndpointCmd.MarkFlagRequired("app-instance-user-arn")
	chimeSdkIdentity_updateAppInstanceUserEndpointCmd.MarkFlagRequired("endpoint-id")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_updateAppInstanceUserEndpointCmd)
}

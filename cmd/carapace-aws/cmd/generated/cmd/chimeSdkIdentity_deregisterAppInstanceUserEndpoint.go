package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_deregisterAppInstanceUserEndpointCmd = &cobra.Command{
	Use:   "deregister-app-instance-user-endpoint",
	Short: "Deregisters an `AppInstanceUserEndpoint`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_deregisterAppInstanceUserEndpointCmd).Standalone()

	chimeSdkIdentity_deregisterAppInstanceUserEndpointCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
	chimeSdkIdentity_deregisterAppInstanceUserEndpointCmd.Flags().String("endpoint-id", "", "The unique identifier of the `AppInstanceUserEndpoint`.")
	chimeSdkIdentity_deregisterAppInstanceUserEndpointCmd.MarkFlagRequired("app-instance-user-arn")
	chimeSdkIdentity_deregisterAppInstanceUserEndpointCmd.MarkFlagRequired("endpoint-id")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_deregisterAppInstanceUserEndpointCmd)
}

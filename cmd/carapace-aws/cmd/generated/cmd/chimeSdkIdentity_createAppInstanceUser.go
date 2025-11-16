package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_createAppInstanceUserCmd = &cobra.Command{
	Use:   "create-app-instance-user",
	Short: "Creates a user under an Amazon Chime `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_createAppInstanceUserCmd).Standalone()

	chimeSdkIdentity_createAppInstanceUserCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance` request.")
	chimeSdkIdentity_createAppInstanceUserCmd.Flags().String("app-instance-user-id", "", "The user ID of the `AppInstance`.")
	chimeSdkIdentity_createAppInstanceUserCmd.Flags().String("client-request-token", "", "The unique ID of the request.")
	chimeSdkIdentity_createAppInstanceUserCmd.Flags().String("expiration-settings", "", "Settings that control the interval after which the `AppInstanceUser` is automatically deleted.")
	chimeSdkIdentity_createAppInstanceUserCmd.Flags().String("metadata", "", "The request's metadata.")
	chimeSdkIdentity_createAppInstanceUserCmd.Flags().String("name", "", "The user's name.")
	chimeSdkIdentity_createAppInstanceUserCmd.Flags().String("tags", "", "Tags assigned to the `AppInstanceUser`.")
	chimeSdkIdentity_createAppInstanceUserCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentity_createAppInstanceUserCmd.MarkFlagRequired("app-instance-user-id")
	chimeSdkIdentity_createAppInstanceUserCmd.MarkFlagRequired("client-request-token")
	chimeSdkIdentity_createAppInstanceUserCmd.MarkFlagRequired("name")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_createAppInstanceUserCmd)
}

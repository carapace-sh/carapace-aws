package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_createAppInstanceCmd = &cobra.Command{
	Use:   "create-app-instance",
	Short: "Creates an Amazon Chime SDK messaging `AppInstance` under an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_createAppInstanceCmd).Standalone()

	chimeSdkIdentity_createAppInstanceCmd.Flags().String("client-request-token", "", "The unique ID of the request.")
	chimeSdkIdentity_createAppInstanceCmd.Flags().String("metadata", "", "The metadata of the `AppInstance`.")
	chimeSdkIdentity_createAppInstanceCmd.Flags().String("name", "", "The name of the `AppInstance`.")
	chimeSdkIdentity_createAppInstanceCmd.Flags().String("tags", "", "Tags assigned to the `AppInstance`.")
	chimeSdkIdentity_createAppInstanceCmd.MarkFlagRequired("client-request-token")
	chimeSdkIdentity_createAppInstanceCmd.MarkFlagRequired("name")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_createAppInstanceCmd)
}

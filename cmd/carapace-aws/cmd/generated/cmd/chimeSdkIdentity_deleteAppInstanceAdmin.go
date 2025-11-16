package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_deleteAppInstanceAdminCmd = &cobra.Command{
	Use:   "delete-app-instance-admin",
	Short: "Demotes an `AppInstanceAdmin` to an `AppInstanceUser` or `AppInstanceBot`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_deleteAppInstanceAdminCmd).Standalone()

	chimeSdkIdentity_deleteAppInstanceAdminCmd.Flags().String("app-instance-admin-arn", "", "The ARN of the `AppInstance`'s administrator.")
	chimeSdkIdentity_deleteAppInstanceAdminCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
	chimeSdkIdentity_deleteAppInstanceAdminCmd.MarkFlagRequired("app-instance-admin-arn")
	chimeSdkIdentity_deleteAppInstanceAdminCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_deleteAppInstanceAdminCmd)
}

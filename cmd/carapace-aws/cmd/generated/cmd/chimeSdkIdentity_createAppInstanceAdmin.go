package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_createAppInstanceAdminCmd = &cobra.Command{
	Use:   "create-app-instance-admin",
	Short: "Promotes an `AppInstanceUser` or `AppInstanceBot` to an `AppInstanceAdmin`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_createAppInstanceAdminCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_createAppInstanceAdminCmd).Standalone()

		chimeSdkIdentity_createAppInstanceAdminCmd.Flags().String("app-instance-admin-arn", "", "The ARN of the administrator of the current `AppInstance`.")
		chimeSdkIdentity_createAppInstanceAdminCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
		chimeSdkIdentity_createAppInstanceAdminCmd.MarkFlagRequired("app-instance-admin-arn")
		chimeSdkIdentity_createAppInstanceAdminCmd.MarkFlagRequired("app-instance-arn")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_createAppInstanceAdminCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_putAppInstanceUserExpirationSettingsCmd = &cobra.Command{
	Use:   "put-app-instance-user-expiration-settings",
	Short: "Sets the number of days before the `AppInstanceUser` is automatically deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_putAppInstanceUserExpirationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_putAppInstanceUserExpirationSettingsCmd).Standalone()

		chimeSdkIdentity_putAppInstanceUserExpirationSettingsCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
		chimeSdkIdentity_putAppInstanceUserExpirationSettingsCmd.Flags().String("expiration-settings", "", "Settings that control the interval after which an `AppInstanceUser` is automatically deleted.")
		chimeSdkIdentity_putAppInstanceUserExpirationSettingsCmd.MarkFlagRequired("app-instance-user-arn")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_putAppInstanceUserExpirationSettingsCmd)
}

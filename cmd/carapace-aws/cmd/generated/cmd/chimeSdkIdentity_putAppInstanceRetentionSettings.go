package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_putAppInstanceRetentionSettingsCmd = &cobra.Command{
	Use:   "put-app-instance-retention-settings",
	Short: "Sets the amount of time in days that a given `AppInstance` retains data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_putAppInstanceRetentionSettingsCmd).Standalone()

	chimeSdkIdentity_putAppInstanceRetentionSettingsCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
	chimeSdkIdentity_putAppInstanceRetentionSettingsCmd.Flags().String("app-instance-retention-settings", "", "The time in days to retain data.")
	chimeSdkIdentity_putAppInstanceRetentionSettingsCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentity_putAppInstanceRetentionSettingsCmd.MarkFlagRequired("app-instance-retention-settings")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_putAppInstanceRetentionSettingsCmd)
}

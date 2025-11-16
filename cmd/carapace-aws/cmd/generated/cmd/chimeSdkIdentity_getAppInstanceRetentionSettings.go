package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_getAppInstanceRetentionSettingsCmd = &cobra.Command{
	Use:   "get-app-instance-retention-settings",
	Short: "Gets the retention settings for an `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_getAppInstanceRetentionSettingsCmd).Standalone()

	chimeSdkIdentity_getAppInstanceRetentionSettingsCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
	chimeSdkIdentity_getAppInstanceRetentionSettingsCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_getAppInstanceRetentionSettingsCmd)
}

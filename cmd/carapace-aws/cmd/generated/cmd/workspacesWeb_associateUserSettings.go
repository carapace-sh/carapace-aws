package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateUserSettingsCmd = &cobra.Command{
	Use:   "associate-user-settings",
	Short: "Associates a user settings resource with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateUserSettingsCmd).Standalone()

	workspacesWeb_associateUserSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_associateUserSettingsCmd.Flags().String("user-settings-arn", "", "The ARN of the user settings.")
	workspacesWeb_associateUserSettingsCmd.MarkFlagRequired("portal-arn")
	workspacesWeb_associateUserSettingsCmd.MarkFlagRequired("user-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_associateUserSettingsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteUserSettingsCmd = &cobra.Command{
	Use:   "delete-user-settings",
	Short: "Deletes user settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteUserSettingsCmd).Standalone()

	workspacesWeb_deleteUserSettingsCmd.Flags().String("user-settings-arn", "", "The ARN of the user settings.")
	workspacesWeb_deleteUserSettingsCmd.MarkFlagRequired("user-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteUserSettingsCmd)
}

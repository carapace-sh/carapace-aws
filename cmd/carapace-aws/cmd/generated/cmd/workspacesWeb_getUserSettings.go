package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getUserSettingsCmd = &cobra.Command{
	Use:   "get-user-settings",
	Short: "Gets user settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getUserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_getUserSettingsCmd).Standalone()

		workspacesWeb_getUserSettingsCmd.Flags().String("user-settings-arn", "", "The ARN of the user settings.")
		workspacesWeb_getUserSettingsCmd.MarkFlagRequired("user-settings-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_getUserSettingsCmd)
}

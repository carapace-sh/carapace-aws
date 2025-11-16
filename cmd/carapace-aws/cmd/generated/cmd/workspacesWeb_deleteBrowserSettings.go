package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteBrowserSettingsCmd = &cobra.Command{
	Use:   "delete-browser-settings",
	Short: "Deletes browser settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteBrowserSettingsCmd).Standalone()

	workspacesWeb_deleteBrowserSettingsCmd.Flags().String("browser-settings-arn", "", "The ARN of the browser settings.")
	workspacesWeb_deleteBrowserSettingsCmd.MarkFlagRequired("browser-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteBrowserSettingsCmd)
}

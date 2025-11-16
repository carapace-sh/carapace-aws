package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateBrowserSettingsCmd = &cobra.Command{
	Use:   "disassociate-browser-settings",
	Short: "Disassociates browser settings from a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateBrowserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_disassociateBrowserSettingsCmd).Standalone()

		workspacesWeb_disassociateBrowserSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_disassociateBrowserSettingsCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateBrowserSettingsCmd)
}

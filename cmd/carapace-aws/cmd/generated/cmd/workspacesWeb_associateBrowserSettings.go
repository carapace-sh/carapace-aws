package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateBrowserSettingsCmd = &cobra.Command{
	Use:   "associate-browser-settings",
	Short: "Associates a browser settings resource with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateBrowserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_associateBrowserSettingsCmd).Standalone()

		workspacesWeb_associateBrowserSettingsCmd.Flags().String("browser-settings-arn", "", "The ARN of the browser settings.")
		workspacesWeb_associateBrowserSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_associateBrowserSettingsCmd.MarkFlagRequired("browser-settings-arn")
		workspacesWeb_associateBrowserSettingsCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_associateBrowserSettingsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateBrowserSettingsCmd = &cobra.Command{
	Use:   "update-browser-settings",
	Short: "Updates browser settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateBrowserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_updateBrowserSettingsCmd).Standalone()

		workspacesWeb_updateBrowserSettingsCmd.Flags().String("browser-policy", "", "A JSON string containing Chrome Enterprise policies that will be applied to all streaming sessions.")
		workspacesWeb_updateBrowserSettingsCmd.Flags().String("browser-settings-arn", "", "The ARN of the browser settings.")
		workspacesWeb_updateBrowserSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesWeb_updateBrowserSettingsCmd.Flags().String("web-content-filtering-policy", "", "The policy that specifies which URLs end users are allowed to access or which URLs or domain categories they are restricted from accessing for enhanced security.")
		workspacesWeb_updateBrowserSettingsCmd.MarkFlagRequired("browser-settings-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_updateBrowserSettingsCmd)
}

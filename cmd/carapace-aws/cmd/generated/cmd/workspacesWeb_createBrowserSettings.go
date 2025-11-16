package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createBrowserSettingsCmd = &cobra.Command{
	Use:   "create-browser-settings",
	Short: "Creates a browser settings resource that can be associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createBrowserSettingsCmd).Standalone()

	workspacesWeb_createBrowserSettingsCmd.Flags().String("additional-encryption-context", "", "Additional encryption context of the browser settings.")
	workspacesWeb_createBrowserSettingsCmd.Flags().String("browser-policy", "", "A JSON string containing Chrome Enterprise policies that will be applied to all streaming sessions.")
	workspacesWeb_createBrowserSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_createBrowserSettingsCmd.Flags().String("customer-managed-key", "", "The custom managed key of the browser settings.")
	workspacesWeb_createBrowserSettingsCmd.Flags().String("tags", "", "The tags to add to the browser settings resource.")
	workspacesWeb_createBrowserSettingsCmd.Flags().String("web-content-filtering-policy", "", "The policy that specifies which URLs end users are allowed to access or which URLs or domain categories they are restricted from accessing for enhanced security.")
	workspacesWebCmd.AddCommand(workspacesWeb_createBrowserSettingsCmd)
}

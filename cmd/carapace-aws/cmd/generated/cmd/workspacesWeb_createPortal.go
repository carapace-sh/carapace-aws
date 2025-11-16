package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createPortalCmd = &cobra.Command{
	Use:   "create-portal",
	Short: "Creates a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createPortalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_createPortalCmd).Standalone()

		workspacesWeb_createPortalCmd.Flags().String("additional-encryption-context", "", "The additional encryption context of the portal.")
		workspacesWeb_createPortalCmd.Flags().String("authentication-type", "", "The type of authentication integration points used when signing into the web portal.")
		workspacesWeb_createPortalCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesWeb_createPortalCmd.Flags().String("customer-managed-key", "", "The customer managed key of the web portal.")
		workspacesWeb_createPortalCmd.Flags().String("display-name", "", "The name of the web portal.")
		workspacesWeb_createPortalCmd.Flags().String("instance-type", "", "The type and resources of the underlying instance.")
		workspacesWeb_createPortalCmd.Flags().String("max-concurrent-sessions", "", "The maximum number of concurrent sessions for the portal.")
		workspacesWeb_createPortalCmd.Flags().String("tags", "", "The tags to add to the web portal.")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_createPortalCmd)
}

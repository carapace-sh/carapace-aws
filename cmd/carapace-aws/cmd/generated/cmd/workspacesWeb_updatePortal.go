package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updatePortalCmd = &cobra.Command{
	Use:   "update-portal",
	Short: "Updates a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updatePortalCmd).Standalone()

	workspacesWeb_updatePortalCmd.Flags().String("authentication-type", "", "The type of authentication integration points used when signing into the web portal.")
	workspacesWeb_updatePortalCmd.Flags().String("display-name", "", "The name of the web portal.")
	workspacesWeb_updatePortalCmd.Flags().String("instance-type", "", "The type and resources of the underlying instance.")
	workspacesWeb_updatePortalCmd.Flags().String("max-concurrent-sessions", "", "The maximum number of concurrent sessions for the portal.")
	workspacesWeb_updatePortalCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_updatePortalCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_updatePortalCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getPortalCmd = &cobra.Command{
	Use:   "get-portal",
	Short: "Gets the web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getPortalCmd).Standalone()

	workspacesWeb_getPortalCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_getPortalCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_getPortalCmd)
}

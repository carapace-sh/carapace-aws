package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deletePortalCmd = &cobra.Command{
	Use:   "delete-portal",
	Short: "Deletes a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deletePortalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_deletePortalCmd).Standalone()

		workspacesWeb_deletePortalCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_deletePortalCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_deletePortalCmd)
}

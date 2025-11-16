package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_deleteEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesThinClient_deleteEnvironmentCmd).Standalone()

		workspacesThinClient_deleteEnvironmentCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesThinClient_deleteEnvironmentCmd.Flags().String("id", "", "The ID of the environment to delete.")
		workspacesThinClient_deleteEnvironmentCmd.MarkFlagRequired("id")
	})
	workspacesThinClientCmd.AddCommand(workspacesThinClient_deleteEnvironmentCmd)
}

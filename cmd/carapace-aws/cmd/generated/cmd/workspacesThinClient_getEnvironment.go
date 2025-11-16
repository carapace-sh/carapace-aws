package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Returns information for an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_getEnvironmentCmd).Standalone()

	workspacesThinClient_getEnvironmentCmd.Flags().String("id", "", "The ID of the environment for which to return information.")
	workspacesThinClient_getEnvironmentCmd.MarkFlagRequired("id")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_getEnvironmentCmd)
}

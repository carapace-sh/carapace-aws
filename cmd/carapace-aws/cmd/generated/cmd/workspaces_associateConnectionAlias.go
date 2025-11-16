package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_associateConnectionAliasCmd = &cobra.Command{
	Use:   "associate-connection-alias",
	Short: "Associates the specified connection alias with the specified directory to enable cross-Region redirection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_associateConnectionAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_associateConnectionAliasCmd).Standalone()

		workspaces_associateConnectionAliasCmd.Flags().String("alias-id", "", "The identifier of the connection alias.")
		workspaces_associateConnectionAliasCmd.Flags().String("resource-id", "", "The identifier of the directory to associate the connection alias with.")
		workspaces_associateConnectionAliasCmd.MarkFlagRequired("alias-id")
		workspaces_associateConnectionAliasCmd.MarkFlagRequired("resource-id")
	})
	workspacesCmd.AddCommand(workspaces_associateConnectionAliasCmd)
}

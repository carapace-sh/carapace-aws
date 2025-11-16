package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createConnectionAliasCmd = &cobra.Command{
	Use:   "create-connection-alias",
	Short: "Creates the specified connection alias for use with cross-Region redirection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createConnectionAliasCmd).Standalone()

	workspaces_createConnectionAliasCmd.Flags().String("connection-string", "", "A connection string in the form of a fully qualified domain name (FQDN), such as `www.example.com`.")
	workspaces_createConnectionAliasCmd.Flags().String("tags", "", "The tags to associate with the connection alias.")
	workspaces_createConnectionAliasCmd.MarkFlagRequired("connection-string")
	workspacesCmd.AddCommand(workspaces_createConnectionAliasCmd)
}

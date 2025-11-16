package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_disassociateConnectionAliasCmd = &cobra.Command{
	Use:   "disassociate-connection-alias",
	Short: "Disassociates a connection alias from a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_disassociateConnectionAliasCmd).Standalone()

	workspaces_disassociateConnectionAliasCmd.Flags().String("alias-id", "", "The identifier of the connection alias to disassociate.")
	workspaces_disassociateConnectionAliasCmd.MarkFlagRequired("alias-id")
	workspacesCmd.AddCommand(workspaces_disassociateConnectionAliasCmd)
}

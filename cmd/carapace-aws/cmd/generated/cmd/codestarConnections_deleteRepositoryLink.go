package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_deleteRepositoryLinkCmd = &cobra.Command{
	Use:   "delete-repository-link",
	Short: "Deletes the association between your connection and a specified external Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_deleteRepositoryLinkCmd).Standalone()

	codestarConnections_deleteRepositoryLinkCmd.Flags().String("repository-link-id", "", "The ID of the repository link to be deleted.")
	codestarConnections_deleteRepositoryLinkCmd.MarkFlagRequired("repository-link-id")
	codestarConnectionsCmd.AddCommand(codestarConnections_deleteRepositoryLinkCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_getRepositoryLinkCmd = &cobra.Command{
	Use:   "get-repository-link",
	Short: "Returns details about a repository link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_getRepositoryLinkCmd).Standalone()

	codestarConnections_getRepositoryLinkCmd.Flags().String("repository-link-id", "", "The ID of the repository link to get.")
	codestarConnections_getRepositoryLinkCmd.MarkFlagRequired("repository-link-id")
	codestarConnectionsCmd.AddCommand(codestarConnections_getRepositoryLinkCmd)
}

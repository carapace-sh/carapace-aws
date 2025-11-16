package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_createRepositoryLinkCmd = &cobra.Command{
	Use:   "create-repository-link",
	Short: "Creates a link to a specified external Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_createRepositoryLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_createRepositoryLinkCmd).Standalone()

		codestarConnections_createRepositoryLinkCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of the connection to be associated with the repository link.")
		codestarConnections_createRepositoryLinkCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) encryption key for the repository to be associated with the repository link.")
		codestarConnections_createRepositoryLinkCmd.Flags().String("owner-id", "", "The owner ID for the repository associated with a specific sync configuration, such as the owner ID in GitHub.")
		codestarConnections_createRepositoryLinkCmd.Flags().String("repository-name", "", "The name of the repository to be associated with the repository link.")
		codestarConnections_createRepositoryLinkCmd.Flags().String("tags", "", "The tags for the repository to be associated with the repository link.")
		codestarConnections_createRepositoryLinkCmd.MarkFlagRequired("connection-arn")
		codestarConnections_createRepositoryLinkCmd.MarkFlagRequired("owner-id")
		codestarConnections_createRepositoryLinkCmd.MarkFlagRequired("repository-name")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_createRepositoryLinkCmd)
}

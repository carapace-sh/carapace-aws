package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_updateRepositoryLinkCmd = &cobra.Command{
	Use:   "update-repository-link",
	Short: "Updates the association between your connection and a specified external Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_updateRepositoryLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_updateRepositoryLinkCmd).Standalone()

		codestarConnections_updateRepositoryLinkCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of the connection for the repository link to be updated.")
		codestarConnections_updateRepositoryLinkCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of the encryption key for the repository link to be updated.")
		codestarConnections_updateRepositoryLinkCmd.Flags().String("repository-link-id", "", "The ID of the repository link to be updated.")
		codestarConnections_updateRepositoryLinkCmd.MarkFlagRequired("repository-link-id")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_updateRepositoryLinkCmd)
}

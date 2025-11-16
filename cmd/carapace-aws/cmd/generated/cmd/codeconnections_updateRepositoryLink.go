package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_updateRepositoryLinkCmd = &cobra.Command{
	Use:   "update-repository-link",
	Short: "Updates the association between your connection and a specified external Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_updateRepositoryLinkCmd).Standalone()

	codeconnections_updateRepositoryLinkCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of the connection for the repository link to be updated.")
	codeconnections_updateRepositoryLinkCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of the encryption key for the repository link to be updated.")
	codeconnections_updateRepositoryLinkCmd.Flags().String("repository-link-id", "", "The ID of the repository link to be updated.")
	codeconnections_updateRepositoryLinkCmd.MarkFlagRequired("repository-link-id")
	codeconnectionsCmd.AddCommand(codeconnections_updateRepositoryLinkCmd)
}

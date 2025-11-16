package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_createRepositoryLinkCmd = &cobra.Command{
	Use:   "create-repository-link",
	Short: "Creates a link to a specified external Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_createRepositoryLinkCmd).Standalone()

	codeconnections_createRepositoryLinkCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of the connection to be associated with the repository link.")
	codeconnections_createRepositoryLinkCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) encryption key for the repository to be associated with the repository link.")
	codeconnections_createRepositoryLinkCmd.Flags().String("owner-id", "", "The owner ID for the repository associated with a specific sync configuration, such as the owner ID in GitHub.")
	codeconnections_createRepositoryLinkCmd.Flags().String("repository-name", "", "The name of the repository to be associated with the repository link.")
	codeconnections_createRepositoryLinkCmd.Flags().String("tags", "", "The tags for the repository to be associated with the repository link.")
	codeconnections_createRepositoryLinkCmd.MarkFlagRequired("connection-arn")
	codeconnections_createRepositoryLinkCmd.MarkFlagRequired("owner-id")
	codeconnections_createRepositoryLinkCmd.MarkFlagRequired("repository-name")
	codeconnectionsCmd.AddCommand(codeconnections_createRepositoryLinkCmd)
}

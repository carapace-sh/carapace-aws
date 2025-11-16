package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_deleteRepositoryCmd = &cobra.Command{
	Use:   "delete-repository",
	Short: "Deletes a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_deleteRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_deleteRepositoryCmd).Standalone()

		ecr_deleteRepositoryCmd.Flags().String("force", "", "If true, deleting the repository force deletes the contents of the repository.")
		ecr_deleteRepositoryCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository to delete.")
		ecr_deleteRepositoryCmd.Flags().String("repository-name", "", "The name of the repository to delete.")
		ecr_deleteRepositoryCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_deleteRepositoryCmd)
}

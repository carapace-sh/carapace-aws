package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_deleteRepositoryCmd = &cobra.Command{
	Use:   "delete-repository",
	Short: "Deletes a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_deleteRepositoryCmd).Standalone()

	ecrPublic_deleteRepositoryCmd.Flags().String("force", "", "The force option can be used to delete a repository that contains images.")
	ecrPublic_deleteRepositoryCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the public registry that contains the repository to delete.")
	ecrPublic_deleteRepositoryCmd.Flags().String("repository-name", "", "The name of the repository to delete.")
	ecrPublic_deleteRepositoryCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_deleteRepositoryCmd)
}

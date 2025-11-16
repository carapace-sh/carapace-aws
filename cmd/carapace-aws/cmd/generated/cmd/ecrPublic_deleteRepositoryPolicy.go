package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_deleteRepositoryPolicyCmd = &cobra.Command{
	Use:   "delete-repository-policy",
	Short: "Deletes the repository policy that's associated with the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_deleteRepositoryPolicyCmd).Standalone()

	ecrPublic_deleteRepositoryPolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the public registry that contains the repository policy to delete.")
	ecrPublic_deleteRepositoryPolicyCmd.Flags().String("repository-name", "", "The name of the repository that's associated with the repository policy to delete.")
	ecrPublic_deleteRepositoryPolicyCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_deleteRepositoryPolicyCmd)
}

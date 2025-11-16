package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_deleteRepositoryPolicyCmd = &cobra.Command{
	Use:   "delete-repository-policy",
	Short: "Deletes the repository policy associated with the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_deleteRepositoryPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_deleteRepositoryPolicyCmd).Standalone()

		ecr_deleteRepositoryPolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository policy to delete.")
		ecr_deleteRepositoryPolicyCmd.Flags().String("repository-name", "", "The name of the repository that is associated with the repository policy to delete.")
		ecr_deleteRepositoryPolicyCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_deleteRepositoryPolicyCmd)
}

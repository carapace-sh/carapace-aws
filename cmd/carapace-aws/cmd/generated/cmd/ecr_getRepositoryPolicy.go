package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getRepositoryPolicyCmd = &cobra.Command{
	Use:   "get-repository-policy",
	Short: "Retrieves the repository policy for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getRepositoryPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_getRepositoryPolicyCmd).Standalone()

		ecr_getRepositoryPolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository.")
		ecr_getRepositoryPolicyCmd.Flags().String("repository-name", "", "The name of the repository with the policy to retrieve.")
		ecr_getRepositoryPolicyCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_getRepositoryPolicyCmd)
}

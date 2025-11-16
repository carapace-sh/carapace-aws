package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_setRepositoryPolicyCmd = &cobra.Command{
	Use:   "set-repository-policy",
	Short: "Applies a repository policy to the specified repository to control access permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_setRepositoryPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_setRepositoryPolicyCmd).Standalone()

		ecr_setRepositoryPolicyCmd.Flags().String("force", "", "If the policy you are attempting to set on a repository policy would prevent you from setting another policy in the future, you must force the [SetRepositoryPolicy]() operation.")
		ecr_setRepositoryPolicyCmd.Flags().String("policy-text", "", "The JSON repository policy text to apply to the repository.")
		ecr_setRepositoryPolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository.")
		ecr_setRepositoryPolicyCmd.Flags().String("repository-name", "", "The name of the repository to receive the policy.")
		ecr_setRepositoryPolicyCmd.MarkFlagRequired("policy-text")
		ecr_setRepositoryPolicyCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_setRepositoryPolicyCmd)
}

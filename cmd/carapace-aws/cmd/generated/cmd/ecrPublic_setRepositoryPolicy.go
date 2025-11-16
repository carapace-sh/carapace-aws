package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_setRepositoryPolicyCmd = &cobra.Command{
	Use:   "set-repository-policy",
	Short: "Applies a repository policy to the specified public repository to control access permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_setRepositoryPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_setRepositoryPolicyCmd).Standalone()

		ecrPublic_setRepositoryPolicyCmd.Flags().String("force", "", "If the policy that you want to set on a repository policy would prevent you from setting another policy in the future, you must force the [SetRepositoryPolicy]() operation.")
		ecrPublic_setRepositoryPolicyCmd.Flags().String("policy-text", "", "The JSON repository policy text to apply to the repository.")
		ecrPublic_setRepositoryPolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the registry that contains the repository.")
		ecrPublic_setRepositoryPolicyCmd.Flags().String("repository-name", "", "The name of the repository to receive the policy.")
		ecrPublic_setRepositoryPolicyCmd.MarkFlagRequired("policy-text")
		ecrPublic_setRepositoryPolicyCmd.MarkFlagRequired("repository-name")
	})
	ecrPublicCmd.AddCommand(ecrPublic_setRepositoryPolicyCmd)
}

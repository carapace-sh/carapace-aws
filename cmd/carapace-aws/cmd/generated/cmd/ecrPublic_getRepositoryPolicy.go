package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_getRepositoryPolicyCmd = &cobra.Command{
	Use:   "get-repository-policy",
	Short: "Retrieves the repository policy for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_getRepositoryPolicyCmd).Standalone()

	ecrPublic_getRepositoryPolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the public registry that contains the repository.")
	ecrPublic_getRepositoryPolicyCmd.Flags().String("repository-name", "", "The name of the repository with the policy to retrieve.")
	ecrPublic_getRepositoryPolicyCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_getRepositoryPolicyCmd)
}

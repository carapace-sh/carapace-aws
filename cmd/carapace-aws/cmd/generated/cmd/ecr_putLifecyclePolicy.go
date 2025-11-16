package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putLifecyclePolicyCmd = &cobra.Command{
	Use:   "put-lifecycle-policy",
	Short: "Creates or updates the lifecycle policy for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putLifecyclePolicyCmd).Standalone()

	ecr_putLifecyclePolicyCmd.Flags().String("lifecycle-policy-text", "", "The JSON repository policy text to apply to the repository.")
	ecr_putLifecyclePolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository.")
	ecr_putLifecyclePolicyCmd.Flags().String("repository-name", "", "The name of the repository to receive the policy.")
	ecr_putLifecyclePolicyCmd.MarkFlagRequired("lifecycle-policy-text")
	ecr_putLifecyclePolicyCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_putLifecyclePolicyCmd)
}

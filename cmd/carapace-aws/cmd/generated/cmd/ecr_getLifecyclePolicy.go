package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getLifecyclePolicyCmd = &cobra.Command{
	Use:   "get-lifecycle-policy",
	Short: "Retrieves the lifecycle policy for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_getLifecyclePolicyCmd).Standalone()

		ecr_getLifecyclePolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository.")
		ecr_getLifecyclePolicyCmd.Flags().String("repository-name", "", "The name of the repository.")
		ecr_getLifecyclePolicyCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_getLifecyclePolicyCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_deleteLifecyclePolicyCmd = &cobra.Command{
	Use:   "delete-lifecycle-policy",
	Short: "Deletes the lifecycle policy associated with the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_deleteLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_deleteLifecyclePolicyCmd).Standalone()

		ecr_deleteLifecyclePolicyCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository.")
		ecr_deleteLifecyclePolicyCmd.Flags().String("repository-name", "", "The name of the repository.")
		ecr_deleteLifecyclePolicyCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_deleteLifecyclePolicyCmd)
}

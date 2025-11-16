package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_startLifecyclePolicyPreviewCmd = &cobra.Command{
	Use:   "start-lifecycle-policy-preview",
	Short: "Starts a preview of a lifecycle policy for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_startLifecyclePolicyPreviewCmd).Standalone()

	ecr_startLifecyclePolicyPreviewCmd.Flags().String("lifecycle-policy-text", "", "The policy to be evaluated against.")
	ecr_startLifecyclePolicyPreviewCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository.")
	ecr_startLifecyclePolicyPreviewCmd.Flags().String("repository-name", "", "The name of the repository to be evaluated.")
	ecr_startLifecyclePolicyPreviewCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_startLifecyclePolicyPreviewCmd)
}

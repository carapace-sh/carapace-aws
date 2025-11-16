package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteLifecyclePolicyCmd = &cobra.Command{
	Use:   "delete-lifecycle-policy",
	Short: "Delete the specified lifecycle policy resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteLifecyclePolicyCmd).Standalone()

	imagebuilder_deleteLifecyclePolicyCmd.Flags().String("lifecycle-policy-arn", "", "The Amazon Resource Name (ARN) of the lifecycle policy resource to delete.")
	imagebuilder_deleteLifecyclePolicyCmd.MarkFlagRequired("lifecycle-policy-arn")
	imagebuilderCmd.AddCommand(imagebuilder_deleteLifecyclePolicyCmd)
}

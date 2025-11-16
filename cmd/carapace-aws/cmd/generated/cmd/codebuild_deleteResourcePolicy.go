package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource policy that is identified by its resource ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_deleteResourcePolicyCmd).Standalone()

		codebuild_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the resource that is associated with the resource policy.")
		codebuild_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	codebuildCmd.AddCommand(codebuild_deleteResourcePolicyCmd)
}

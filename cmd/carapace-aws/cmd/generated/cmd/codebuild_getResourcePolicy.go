package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Gets a resource policy that is identified by its resource ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_getResourcePolicyCmd).Standalone()

	codebuild_getResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the resource that is associated with the resource policy.")
	codebuild_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	codebuildCmd.AddCommand(codebuild_getResourcePolicyCmd)
}

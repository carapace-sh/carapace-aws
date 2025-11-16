package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_updateProjectVisibilityCmd = &cobra.Command{
	Use:   "update-project-visibility",
	Short: "Changes the public visibility for a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_updateProjectVisibilityCmd).Standalone()

	codebuild_updateProjectVisibilityCmd.Flags().String("project-arn", "", "The Amazon Resource Name (ARN) of the build project.")
	codebuild_updateProjectVisibilityCmd.Flags().String("project-visibility", "", "")
	codebuild_updateProjectVisibilityCmd.Flags().String("resource-access-role", "", "The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.")
	codebuild_updateProjectVisibilityCmd.MarkFlagRequired("project-arn")
	codebuild_updateProjectVisibilityCmd.MarkFlagRequired("project-visibility")
	codebuildCmd.AddCommand(codebuild_updateProjectVisibilityCmd)
}

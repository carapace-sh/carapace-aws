package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a build project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_createProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_createProjectCmd).Standalone()

		codebuild_createProjectCmd.Flags().String("artifacts", "", "Information about the build output artifacts for the build project.")
		codebuild_createProjectCmd.Flags().String("auto-retry-limit", "", "The maximum number of additional automatic retries after a failed build.")
		codebuild_createProjectCmd.Flags().String("badge-enabled", "", "Set this to true to generate a publicly accessible URL for your project's build badge.")
		codebuild_createProjectCmd.Flags().String("build-batch-config", "", "A [ProjectBuildBatchConfig]() object that defines the batch build options for the project.")
		codebuild_createProjectCmd.Flags().String("cache", "", "Stores recently used information so that it can be quickly accessed at a later time.")
		codebuild_createProjectCmd.Flags().String("concurrent-build-limit", "", "The maximum number of concurrent builds that are allowed for this project.")
		codebuild_createProjectCmd.Flags().String("description", "", "A description that makes the build project easy to identify.")
		codebuild_createProjectCmd.Flags().String("encryption-key", "", "The Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.")
		codebuild_createProjectCmd.Flags().String("environment", "", "Information about the build environment for the build project.")
		codebuild_createProjectCmd.Flags().String("file-system-locations", "", "An array of `ProjectFileSystemLocation` objects for a CodeBuild build project.")
		codebuild_createProjectCmd.Flags().String("logs-config", "", "Information about logs for the build project.")
		codebuild_createProjectCmd.Flags().String("name", "", "The name of the build project.")
		codebuild_createProjectCmd.Flags().String("queued-timeout-in-minutes", "", "The number of minutes a build is allowed to be queued before it times out.")
		codebuild_createProjectCmd.Flags().String("secondary-artifacts", "", "An array of `ProjectArtifacts` objects.")
		codebuild_createProjectCmd.Flags().String("secondary-source-versions", "", "An array of `ProjectSourceVersion` objects.")
		codebuild_createProjectCmd.Flags().String("secondary-sources", "", "An array of `ProjectSource` objects.")
		codebuild_createProjectCmd.Flags().String("service-role", "", "The ARN of the IAM role that enables CodeBuild to interact with dependent Amazon Web Services services on behalf of the Amazon Web Services account.")
		codebuild_createProjectCmd.Flags().String("source", "", "Information about the build input source code for the build project.")
		codebuild_createProjectCmd.Flags().String("source-version", "", "A version of the build input to be built for this project.")
		codebuild_createProjectCmd.Flags().String("tags", "", "A list of tag key and value pairs associated with this build project.")
		codebuild_createProjectCmd.Flags().String("timeout-in-minutes", "", "How long, in minutes, from 5 to 2160 (36 hours), for CodeBuild to wait before it times out any build that has not been marked as completed.")
		codebuild_createProjectCmd.Flags().String("vpc-config", "", "VpcConfig enables CodeBuild to access resources in an Amazon VPC.")
		codebuild_createProjectCmd.MarkFlagRequired("artifacts")
		codebuild_createProjectCmd.MarkFlagRequired("environment")
		codebuild_createProjectCmd.MarkFlagRequired("name")
		codebuild_createProjectCmd.MarkFlagRequired("service-role")
		codebuild_createProjectCmd.MarkFlagRequired("source")
	})
	codebuildCmd.AddCommand(codebuild_createProjectCmd)
}

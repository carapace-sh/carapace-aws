package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Changes the settings of a build project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_updateProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_updateProjectCmd).Standalone()

		codebuild_updateProjectCmd.Flags().String("artifacts", "", "Information to be changed about the build output artifacts for the build project.")
		codebuild_updateProjectCmd.Flags().String("auto-retry-limit", "", "The maximum number of additional automatic retries after a failed build.")
		codebuild_updateProjectCmd.Flags().String("badge-enabled", "", "Set this to true to generate a publicly accessible URL for your project's build badge.")
		codebuild_updateProjectCmd.Flags().String("build-batch-config", "", "")
		codebuild_updateProjectCmd.Flags().String("cache", "", "Stores recently used information so that it can be quickly accessed at a later time.")
		codebuild_updateProjectCmd.Flags().String("concurrent-build-limit", "", "The maximum number of concurrent builds that are allowed for this project.")
		codebuild_updateProjectCmd.Flags().String("description", "", "A new or replacement description of the build project.")
		codebuild_updateProjectCmd.Flags().String("encryption-key", "", "The Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.")
		codebuild_updateProjectCmd.Flags().String("environment", "", "Information to be changed about the build environment for the build project.")
		codebuild_updateProjectCmd.Flags().String("file-system-locations", "", "An array of `ProjectFileSystemLocation` objects for a CodeBuild build project.")
		codebuild_updateProjectCmd.Flags().String("logs-config", "", "Information about logs for the build project.")
		codebuild_updateProjectCmd.Flags().String("name", "", "The name of the build project.")
		codebuild_updateProjectCmd.Flags().String("queued-timeout-in-minutes", "", "The number of minutes a build is allowed to be queued before it times out.")
		codebuild_updateProjectCmd.Flags().String("secondary-artifacts", "", "An array of `ProjectArtifact` objects.")
		codebuild_updateProjectCmd.Flags().String("secondary-source-versions", "", "An array of `ProjectSourceVersion` objects.")
		codebuild_updateProjectCmd.Flags().String("secondary-sources", "", "An array of `ProjectSource` objects.")
		codebuild_updateProjectCmd.Flags().String("service-role", "", "The replacement ARN of the IAM role that enables CodeBuild to interact with dependent Amazon Web Services services on behalf of the Amazon Web Services account.")
		codebuild_updateProjectCmd.Flags().String("source", "", "Information to be changed about the build input source code for the build project.")
		codebuild_updateProjectCmd.Flags().String("source-version", "", "A version of the build input to be built for this project.")
		codebuild_updateProjectCmd.Flags().String("tags", "", "An updated list of tag key and value pairs associated with this build project.")
		codebuild_updateProjectCmd.Flags().String("timeout-in-minutes", "", "The replacement value in minutes, from 5 to 2160 (36 hours), for CodeBuild to wait before timing out any related build that did not get marked as completed.")
		codebuild_updateProjectCmd.Flags().String("vpc-config", "", "VpcConfig enables CodeBuild to access resources in an Amazon VPC.")
		codebuild_updateProjectCmd.MarkFlagRequired("name")
	})
	codebuildCmd.AddCommand(codebuild_updateProjectCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_startBuildCmd = &cobra.Command{
	Use:   "start-build",
	Short: "Starts running a build with the settings defined in the project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_startBuildCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_startBuildCmd).Standalone()

		codebuild_startBuildCmd.Flags().String("artifacts-override", "", "Build output artifact settings that override, for this build only, the latest ones already defined in the build project.")
		codebuild_startBuildCmd.Flags().String("auto-retry-limit-override", "", "The maximum number of additional automatic retries after a failed build.")
		codebuild_startBuildCmd.Flags().String("build-status-config-override", "", "Contains information that defines how the build project reports the build status to the source provider.")
		codebuild_startBuildCmd.Flags().String("buildspec-override", "", "A buildspec file declaration that overrides the latest one defined in the build project, for this build only.")
		codebuild_startBuildCmd.Flags().String("cache-override", "", "A ProjectCache object specified for this build that overrides the one defined in the build project.")
		codebuild_startBuildCmd.Flags().String("certificate-override", "", "The name of a certificate for this build that overrides the one specified in the build project.")
		codebuild_startBuildCmd.Flags().String("compute-type-override", "", "The name of a compute type for this build that overrides the one specified in the build project.")
		codebuild_startBuildCmd.Flags().String("debug-session-enabled", "", "Specifies if session debugging is enabled for this build.")
		codebuild_startBuildCmd.Flags().String("encryption-key-override", "", "The Key Management Service customer master key (CMK) that overrides the one specified in the build project.")
		codebuild_startBuildCmd.Flags().String("environment-type-override", "", "A container type for this build that overrides the one specified in the build project.")
		codebuild_startBuildCmd.Flags().String("environment-variables-override", "", "A set of environment variables that overrides, for this build only, the latest ones already defined in the build project.")
		codebuild_startBuildCmd.Flags().String("fleet-override", "", "A ProjectFleet object specified for this build that overrides the one defined in the build project.")
		codebuild_startBuildCmd.Flags().String("git-clone-depth-override", "", "The user-defined depth of history, with a minimum value of 0, that overrides, for this build only, any previous depth of history defined in the build project.")
		codebuild_startBuildCmd.Flags().String("git-submodules-config-override", "", "Information about the Git submodules configuration for this build of an CodeBuild build project.")
		codebuild_startBuildCmd.Flags().String("idempotency-token", "", "A unique, case sensitive identifier you provide to ensure the idempotency of the StartBuild request.")
		codebuild_startBuildCmd.Flags().String("image-override", "", "The name of an image for this build that overrides the one specified in the build project.")
		codebuild_startBuildCmd.Flags().String("image-pull-credentials-type-override", "", "The type of credentials CodeBuild uses to pull images in your build.")
		codebuild_startBuildCmd.Flags().String("insecure-ssl-override", "", "Enable this flag to override the insecure SSL setting that is specified in the build project.")
		codebuild_startBuildCmd.Flags().String("logs-config-override", "", "Log settings for this build that override the log settings defined in the build project.")
		codebuild_startBuildCmd.Flags().String("privileged-mode-override", "", "Enable this flag to override privileged mode in the build project.")
		codebuild_startBuildCmd.Flags().String("project-name", "", "The name of the CodeBuild build project to start running a build.")
		codebuild_startBuildCmd.Flags().String("queued-timeout-in-minutes-override", "", "The number of minutes a build is allowed to be queued before it times out.")
		codebuild_startBuildCmd.Flags().String("registry-credential-override", "", "The credentials for access to a private registry.")
		codebuild_startBuildCmd.Flags().String("report-build-status-override", "", "Set to true to report to your source provider the status of a build's start and completion.")
		codebuild_startBuildCmd.Flags().String("secondary-artifacts-override", "", "An array of `ProjectArtifacts` objects.")
		codebuild_startBuildCmd.Flags().String("secondary-sources-override", "", "An array of `ProjectSource` objects.")
		codebuild_startBuildCmd.Flags().String("secondary-sources-version-override", "", "An array of `ProjectSourceVersion` objects that specify one or more versions of the project's secondary sources to be used for this build only.")
		codebuild_startBuildCmd.Flags().String("service-role-override", "", "The name of a service role for this build that overrides the one specified in the build project.")
		codebuild_startBuildCmd.Flags().String("source-auth-override", "", "An authorization type for this build that overrides the one defined in the build project.")
		codebuild_startBuildCmd.Flags().String("source-location-override", "", "A location that overrides, for this build, the source location for the one defined in the build project.")
		codebuild_startBuildCmd.Flags().String("source-type-override", "", "A source input type, for this build, that overrides the source input defined in the build project.")
		codebuild_startBuildCmd.Flags().String("source-version", "", "The version of the build input to be built, for this build only.")
		codebuild_startBuildCmd.Flags().String("timeout-in-minutes-override", "", "The number of build timeout minutes, from 5 to 2160 (36 hours), that overrides, for this build only, the latest setting already defined in the build project.")
		codebuild_startBuildCmd.MarkFlagRequired("project-name")
	})
	codebuildCmd.AddCommand(codebuild_startBuildCmd)
}

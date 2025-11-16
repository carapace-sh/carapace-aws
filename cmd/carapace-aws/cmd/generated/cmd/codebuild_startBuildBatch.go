package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_startBuildBatchCmd = &cobra.Command{
	Use:   "start-build-batch",
	Short: "Starts a batch build for a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_startBuildBatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_startBuildBatchCmd).Standalone()

		codebuild_startBuildBatchCmd.Flags().String("artifacts-override", "", "An array of `ProjectArtifacts` objects that contains information about the build output artifact overrides for the build project.")
		codebuild_startBuildBatchCmd.Flags().String("build-batch-config-override", "", "A `BuildBatchConfigOverride` object that contains batch build configuration overrides.")
		codebuild_startBuildBatchCmd.Flags().String("build-timeout-in-minutes-override", "", "Overrides the build timeout specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("buildspec-override", "", "A buildspec file declaration that overrides, for this build only, the latest one already defined in the build project.")
		codebuild_startBuildBatchCmd.Flags().String("cache-override", "", "A `ProjectCache` object that specifies cache overrides.")
		codebuild_startBuildBatchCmd.Flags().String("certificate-override", "", "The name of a certificate for this batch build that overrides the one specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("compute-type-override", "", "The name of a compute type for this batch build that overrides the one specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("debug-session-enabled", "", "Specifies if session debugging is enabled for this batch build.")
		codebuild_startBuildBatchCmd.Flags().String("encryption-key-override", "", "The Key Management Service customer master key (CMK) that overrides the one specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("environment-type-override", "", "A container type for this batch build that overrides the one specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("environment-variables-override", "", "An array of `EnvironmentVariable` objects that override, or add to, the environment variables defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("git-clone-depth-override", "", "The user-defined depth of history, with a minimum value of 0, that overrides, for this batch build only, any previous depth of history defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("git-submodules-config-override", "", "A `GitSubmodulesConfig` object that overrides the Git submodules configuration for this batch build.")
		codebuild_startBuildBatchCmd.Flags().String("idempotency-token", "", "A unique, case sensitive identifier you provide to ensure the idempotency of the `StartBuildBatch` request.")
		codebuild_startBuildBatchCmd.Flags().String("image-override", "", "The name of an image for this batch build that overrides the one specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("image-pull-credentials-type-override", "", "The type of credentials CodeBuild uses to pull images in your batch build.")
		codebuild_startBuildBatchCmd.Flags().String("insecure-ssl-override", "", "Enable this flag to override the insecure SSL setting that is specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("logs-config-override", "", "A `LogsConfig` object that override the log settings defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("privileged-mode-override", "", "Enable this flag to override privileged mode in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("project-name", "", "The name of the project.")
		codebuild_startBuildBatchCmd.Flags().String("queued-timeout-in-minutes-override", "", "The number of minutes a batch build is allowed to be queued before it times out.")
		codebuild_startBuildBatchCmd.Flags().String("registry-credential-override", "", "A `RegistryCredential` object that overrides credentials for access to a private registry.")
		codebuild_startBuildBatchCmd.Flags().String("report-build-batch-status-override", "", "Set to `true` to report to your source provider the status of a batch build's start and completion.")
		codebuild_startBuildBatchCmd.Flags().String("secondary-artifacts-override", "", "An array of `ProjectArtifacts` objects that override the secondary artifacts defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("secondary-sources-override", "", "An array of `ProjectSource` objects that override the secondary sources defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("secondary-sources-version-override", "", "An array of `ProjectSourceVersion` objects that override the secondary source versions in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("service-role-override", "", "The name of a service role for this batch build that overrides the one specified in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("source-auth-override", "", "A `SourceAuth` object that overrides the one defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("source-location-override", "", "A location that overrides, for this batch build, the source location defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("source-type-override", "", "The source input type that overrides the source input defined in the batch build project.")
		codebuild_startBuildBatchCmd.Flags().String("source-version", "", "The version of the batch build input to be built, for this build only.")
		codebuild_startBuildBatchCmd.MarkFlagRequired("project-name")
	})
	codebuildCmd.AddCommand(codebuild_startBuildBatchCmd)
}

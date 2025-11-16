package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_createBranchCmd = &cobra.Command{
	Use:   "create-branch",
	Short: "Creates a new branch for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_createBranchCmd).Standalone()

	amplify_createBranchCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_createBranchCmd.Flags().String("backend", "", "The backend for a `Branch` of an Amplify app.")
	amplify_createBranchCmd.Flags().String("backend-environment-arn", "", "The Amazon Resource Name (ARN) for a backend environment that is part of a Gen 1 Amplify app.")
	amplify_createBranchCmd.Flags().String("basic-auth-credentials", "", "The basic authorization credentials for the branch.")
	amplify_createBranchCmd.Flags().String("branch-name", "", "The name for the branch.")
	amplify_createBranchCmd.Flags().String("build-spec", "", "The build specification (build spec) for the branch.")
	amplify_createBranchCmd.Flags().String("compute-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to assign to a branch of an SSR app.")
	amplify_createBranchCmd.Flags().String("description", "", "The description for the branch.")
	amplify_createBranchCmd.Flags().String("display-name", "", "The display name for a branch.")
	amplify_createBranchCmd.Flags().String("enable-auto-build", "", "Enables auto building for the branch.")
	amplify_createBranchCmd.Flags().String("enable-basic-auth", "", "Enables basic authorization for the branch.")
	amplify_createBranchCmd.Flags().String("enable-notification", "", "Enables notifications for the branch.")
	amplify_createBranchCmd.Flags().String("enable-performance-mode", "", "Enables performance mode for the branch.")
	amplify_createBranchCmd.Flags().String("enable-pull-request-preview", "", "Enables pull request previews for this branch.")
	amplify_createBranchCmd.Flags().String("enable-skew-protection", "", "Specifies whether the skew protection feature is enabled for the branch.")
	amplify_createBranchCmd.Flags().String("environment-variables", "", "The environment variables for the branch.")
	amplify_createBranchCmd.Flags().String("framework", "", "The framework for the branch.")
	amplify_createBranchCmd.Flags().String("pull-request-environment-name", "", "The Amplify environment name for the pull request.")
	amplify_createBranchCmd.Flags().String("stage", "", "Describes the current stage for the branch.")
	amplify_createBranchCmd.Flags().String("tags", "", "The tag for the branch.")
	amplify_createBranchCmd.Flags().String("ttl", "", "The content Time To Live (TTL) for the website in seconds.")
	amplify_createBranchCmd.MarkFlagRequired("app-id")
	amplify_createBranchCmd.MarkFlagRequired("branch-name")
	amplifyCmd.AddCommand(amplify_createBranchCmd)
}

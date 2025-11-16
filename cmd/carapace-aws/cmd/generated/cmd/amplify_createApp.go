package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "Creates a new Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_createAppCmd).Standalone()

	amplify_createAppCmd.Flags().String("access-token", "", "The personal access token for a GitHub repository for an Amplify app.")
	amplify_createAppCmd.Flags().String("auto-branch-creation-config", "", "The automated branch creation configuration for an Amplify app.")
	amplify_createAppCmd.Flags().String("auto-branch-creation-patterns", "", "The automated branch creation glob patterns for an Amplify app.")
	amplify_createAppCmd.Flags().String("basic-auth-credentials", "", "The credentials for basic authorization for an Amplify app.")
	amplify_createAppCmd.Flags().String("build-spec", "", "The build specification (build spec) for an Amplify app.")
	amplify_createAppCmd.Flags().String("cache-config", "", "The cache configuration for the Amplify app.")
	amplify_createAppCmd.Flags().String("compute-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to assign to an SSR app.")
	amplify_createAppCmd.Flags().String("custom-headers", "", "The custom HTTP headers for an Amplify app.")
	amplify_createAppCmd.Flags().String("custom-rules", "", "The custom rewrite and redirect rules for an Amplify app.")
	amplify_createAppCmd.Flags().String("description", "", "The description of the Amplify app.")
	amplify_createAppCmd.Flags().String("enable-auto-branch-creation", "", "Enables automated branch creation for an Amplify app.")
	amplify_createAppCmd.Flags().String("enable-basic-auth", "", "Enables basic authorization for an Amplify app.")
	amplify_createAppCmd.Flags().String("enable-branch-auto-build", "", "Enables the auto building of branches for an Amplify app.")
	amplify_createAppCmd.Flags().String("enable-branch-auto-deletion", "", "Automatically disconnects a branch in the Amplify console when you delete a branch from your Git repository.")
	amplify_createAppCmd.Flags().String("environment-variables", "", "The environment variables map for an Amplify app.")
	amplify_createAppCmd.Flags().String("iam-service-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role for the Amplify app.")
	amplify_createAppCmd.Flags().String("job-config", "", "Describes the configuration details that apply to the jobs for an Amplify app.")
	amplify_createAppCmd.Flags().String("name", "", "The name of the Amplify app.")
	amplify_createAppCmd.Flags().String("oauth-token", "", "The OAuth token for a third-party source control system for an Amplify app.")
	amplify_createAppCmd.Flags().String("platform", "", "The platform for the Amplify app.")
	amplify_createAppCmd.Flags().String("repository", "", "The Git repository for the Amplify app.")
	amplify_createAppCmd.Flags().String("tags", "", "The tag for an Amplify app.")
	amplify_createAppCmd.MarkFlagRequired("name")
	amplifyCmd.AddCommand(amplify_createAppCmd)
}

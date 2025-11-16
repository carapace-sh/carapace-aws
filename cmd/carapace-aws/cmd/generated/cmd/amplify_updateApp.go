package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_updateAppCmd = &cobra.Command{
	Use:   "update-app",
	Short: "Updates an existing Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_updateAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_updateAppCmd).Standalone()

		amplify_updateAppCmd.Flags().String("access-token", "", "The personal access token for a GitHub repository for an Amplify app.")
		amplify_updateAppCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_updateAppCmd.Flags().String("auto-branch-creation-config", "", "The automated branch creation configuration for an Amplify app.")
		amplify_updateAppCmd.Flags().String("auto-branch-creation-patterns", "", "Describes the automated branch creation glob patterns for an Amplify app.")
		amplify_updateAppCmd.Flags().String("basic-auth-credentials", "", "The basic authorization credentials for an Amplify app.")
		amplify_updateAppCmd.Flags().String("build-spec", "", "The build specification (build spec) for an Amplify app.")
		amplify_updateAppCmd.Flags().String("cache-config", "", "The cache configuration for the Amplify app.")
		amplify_updateAppCmd.Flags().String("compute-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to assign to an SSR app.")
		amplify_updateAppCmd.Flags().String("custom-headers", "", "The custom HTTP headers for an Amplify app.")
		amplify_updateAppCmd.Flags().String("custom-rules", "", "The custom redirect and rewrite rules for an Amplify app.")
		amplify_updateAppCmd.Flags().String("description", "", "The description for an Amplify app.")
		amplify_updateAppCmd.Flags().String("enable-auto-branch-creation", "", "Enables automated branch creation for an Amplify app.")
		amplify_updateAppCmd.Flags().String("enable-basic-auth", "", "Enables basic authorization for an Amplify app.")
		amplify_updateAppCmd.Flags().String("enable-branch-auto-build", "", "Enables branch auto-building for an Amplify app.")
		amplify_updateAppCmd.Flags().String("enable-branch-auto-deletion", "", "Automatically disconnects a branch in the Amplify console when you delete a branch from your Git repository.")
		amplify_updateAppCmd.Flags().String("environment-variables", "", "The environment variables for an Amplify app.")
		amplify_updateAppCmd.Flags().String("iam-service-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role for the Amplify app.")
		amplify_updateAppCmd.Flags().String("job-config", "", "Describes the configuration details that apply to the jobs for an Amplify app.")
		amplify_updateAppCmd.Flags().String("name", "", "The name for an Amplify app.")
		amplify_updateAppCmd.Flags().String("oauth-token", "", "The OAuth token for a third-party source control system for an Amplify app.")
		amplify_updateAppCmd.Flags().String("platform", "", "The platform for the Amplify app.")
		amplify_updateAppCmd.Flags().String("repository", "", "The name of the Git repository for an Amplify app.")
		amplify_updateAppCmd.MarkFlagRequired("app-id")
	})
	amplifyCmd.AddCommand(amplify_updateAppCmd)
}

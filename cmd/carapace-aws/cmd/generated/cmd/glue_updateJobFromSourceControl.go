package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateJobFromSourceControlCmd = &cobra.Command{
	Use:   "update-job-from-source-control",
	Short: "Synchronizes a job from the source control repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateJobFromSourceControlCmd).Standalone()

	glue_updateJobFromSourceControlCmd.Flags().String("auth-strategy", "", "The type of authentication, which can be an authentication token stored in Amazon Web Services Secrets Manager, or a personal access token.")
	glue_updateJobFromSourceControlCmd.Flags().String("auth-token", "", "The value of the authorization token.")
	glue_updateJobFromSourceControlCmd.Flags().String("branch-name", "", "An optional branch in the remote repository.")
	glue_updateJobFromSourceControlCmd.Flags().String("commit-id", "", "A commit ID for a commit in the remote repository.")
	glue_updateJobFromSourceControlCmd.Flags().String("folder", "", "An optional folder in the remote repository.")
	glue_updateJobFromSourceControlCmd.Flags().String("job-name", "", "The name of the Glue job to be synchronized to or from the remote repository.")
	glue_updateJobFromSourceControlCmd.Flags().String("provider", "", "The provider for the remote repository.")
	glue_updateJobFromSourceControlCmd.Flags().String("repository-name", "", "The name of the remote repository that contains the job artifacts.")
	glue_updateJobFromSourceControlCmd.Flags().String("repository-owner", "", "The owner of the remote repository that contains the job artifacts.")
	glueCmd.AddCommand(glue_updateJobFromSourceControlCmd)
}

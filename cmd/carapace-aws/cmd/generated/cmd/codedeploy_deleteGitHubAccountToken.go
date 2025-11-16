package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_deleteGitHubAccountTokenCmd = &cobra.Command{
	Use:   "delete-git-hub-account-token",
	Short: "Deletes a GitHub account connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_deleteGitHubAccountTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_deleteGitHubAccountTokenCmd).Standalone()

		codedeploy_deleteGitHubAccountTokenCmd.Flags().String("token-name", "", "The name of the GitHub account connection to delete.")
	})
	codedeployCmd.AddCommand(codedeploy_deleteGitHubAccountTokenCmd)
}

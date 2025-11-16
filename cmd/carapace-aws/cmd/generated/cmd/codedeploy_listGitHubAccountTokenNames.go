package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listGitHubAccountTokenNamesCmd = &cobra.Command{
	Use:   "list-git-hub-account-token-names",
	Short: "Lists the names of stored connections to GitHub accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listGitHubAccountTokenNamesCmd).Standalone()

	codedeploy_listGitHubAccountTokenNamesCmd.Flags().String("next-token", "", "An identifier returned from the previous `ListGitHubAccountTokenNames` call.")
	codedeployCmd.AddCommand(codedeploy_listGitHubAccountTokenNamesCmd)
}

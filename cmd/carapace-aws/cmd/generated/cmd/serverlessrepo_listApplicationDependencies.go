package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_listApplicationDependenciesCmd = &cobra.Command{
	Use:   "list-application-dependencies",
	Short: "Retrieves the list of applications nested in the containing application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_listApplicationDependenciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepo_listApplicationDependenciesCmd).Standalone()

		serverlessrepo_listApplicationDependenciesCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
		serverlessrepo_listApplicationDependenciesCmd.Flags().String("max-items", "", "The total number of items to return.")
		serverlessrepo_listApplicationDependenciesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		serverlessrepo_listApplicationDependenciesCmd.Flags().String("semantic-version", "", "The semantic version of the application to get.")
		serverlessrepo_listApplicationDependenciesCmd.MarkFlagRequired("application-id")
	})
	serverlessrepoCmd.AddCommand(serverlessrepo_listApplicationDependenciesCmd)
}

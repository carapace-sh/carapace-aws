package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listSourceRepositoriesCmd = &cobra.Command{
	Use:   "list-source-repositories",
	Short: "Retrieves a list of source repositories in a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listSourceRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_listSourceRepositoriesCmd).Standalone()

		codecatalyst_listSourceRepositoriesCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
		codecatalyst_listSourceRepositoriesCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
		codecatalyst_listSourceRepositoriesCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_listSourceRepositoriesCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_listSourceRepositoriesCmd.MarkFlagRequired("project-name")
		codecatalyst_listSourceRepositoriesCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_listSourceRepositoriesCmd)
}

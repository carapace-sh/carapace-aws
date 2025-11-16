package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_deleteSourceRepositoryCmd = &cobra.Command{
	Use:   "delete-source-repository",
	Short: "Deletes a source repository in Amazon CodeCatalyst.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_deleteSourceRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_deleteSourceRepositoryCmd).Standalone()

		codecatalyst_deleteSourceRepositoryCmd.Flags().String("name", "", "The name of the source repository.")
		codecatalyst_deleteSourceRepositoryCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_deleteSourceRepositoryCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_deleteSourceRepositoryCmd.MarkFlagRequired("name")
		codecatalyst_deleteSourceRepositoryCmd.MarkFlagRequired("project-name")
		codecatalyst_deleteSourceRepositoryCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_deleteSourceRepositoryCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_createSourceRepositoryCmd = &cobra.Command{
	Use:   "create-source-repository",
	Short: "Creates an empty Git-based source repository in a specified project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_createSourceRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_createSourceRepositoryCmd).Standalone()

		codecatalyst_createSourceRepositoryCmd.Flags().String("description", "", "The description of the source repository.")
		codecatalyst_createSourceRepositoryCmd.Flags().String("name", "", "The name of the source repository.")
		codecatalyst_createSourceRepositoryCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_createSourceRepositoryCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_createSourceRepositoryCmd.MarkFlagRequired("name")
		codecatalyst_createSourceRepositoryCmd.MarkFlagRequired("project-name")
		codecatalyst_createSourceRepositoryCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_createSourceRepositoryCmd)
}

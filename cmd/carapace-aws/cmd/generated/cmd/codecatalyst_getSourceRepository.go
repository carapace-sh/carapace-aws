package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getSourceRepositoryCmd = &cobra.Command{
	Use:   "get-source-repository",
	Short: "Returns information about a source repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getSourceRepositoryCmd).Standalone()

	codecatalyst_getSourceRepositoryCmd.Flags().String("name", "", "The name of the source repository.")
	codecatalyst_getSourceRepositoryCmd.Flags().String("project-name", "", "The name of the project in the space.")
	codecatalyst_getSourceRepositoryCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_getSourceRepositoryCmd.MarkFlagRequired("name")
	codecatalyst_getSourceRepositoryCmd.MarkFlagRequired("project-name")
	codecatalyst_getSourceRepositoryCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_getSourceRepositoryCmd)
}

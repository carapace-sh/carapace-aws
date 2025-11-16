package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateRepositoryNameCmd = &cobra.Command{
	Use:   "update-repository-name",
	Short: "Renames a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateRepositoryNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_updateRepositoryNameCmd).Standalone()

		codecommit_updateRepositoryNameCmd.Flags().String("new-name", "", "The new name for the repository.")
		codecommit_updateRepositoryNameCmd.Flags().String("old-name", "", "The current name of the repository.")
		codecommit_updateRepositoryNameCmd.MarkFlagRequired("new-name")
		codecommit_updateRepositoryNameCmd.MarkFlagRequired("old-name")
	})
	codecommitCmd.AddCommand(codecommit_updateRepositoryNameCmd)
}

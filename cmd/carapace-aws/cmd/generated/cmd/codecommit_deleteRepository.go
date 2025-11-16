package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_deleteRepositoryCmd = &cobra.Command{
	Use:   "delete-repository",
	Short: "Deletes a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_deleteRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_deleteRepositoryCmd).Standalone()

		codecommit_deleteRepositoryCmd.Flags().String("repository-name", "", "The name of the repository to delete.")
		codecommit_deleteRepositoryCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_deleteRepositoryCmd)
}

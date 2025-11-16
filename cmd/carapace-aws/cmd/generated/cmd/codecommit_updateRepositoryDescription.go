package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateRepositoryDescriptionCmd = &cobra.Command{
	Use:   "update-repository-description",
	Short: "Sets or changes the comment or description for a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateRepositoryDescriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_updateRepositoryDescriptionCmd).Standalone()

		codecommit_updateRepositoryDescriptionCmd.Flags().String("repository-description", "", "The new comment or description for the specified repository.")
		codecommit_updateRepositoryDescriptionCmd.Flags().String("repository-name", "", "The name of the repository to set or change the comment or description for.")
		codecommit_updateRepositoryDescriptionCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_updateRepositoryDescriptionCmd)
}

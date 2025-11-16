package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getRepositoryCmd = &cobra.Command{
	Use:   "get-repository",
	Short: "Returns information about a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getRepositoryCmd).Standalone()

	codecommit_getRepositoryCmd.Flags().String("repository-name", "", "The name of the repository to get information about.")
	codecommit_getRepositoryCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getRepositoryCmd)
}

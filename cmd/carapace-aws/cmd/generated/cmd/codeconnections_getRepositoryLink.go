package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_getRepositoryLinkCmd = &cobra.Command{
	Use:   "get-repository-link",
	Short: "Returns details about a repository link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_getRepositoryLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_getRepositoryLinkCmd).Standalone()

		codeconnections_getRepositoryLinkCmd.Flags().String("repository-link-id", "", "The ID of the repository link to get.")
		codeconnections_getRepositoryLinkCmd.MarkFlagRequired("repository-link-id")
	})
	codeconnectionsCmd.AddCommand(codeconnections_getRepositoryLinkCmd)
}

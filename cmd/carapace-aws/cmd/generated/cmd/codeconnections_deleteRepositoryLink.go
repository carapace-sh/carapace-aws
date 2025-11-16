package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_deleteRepositoryLinkCmd = &cobra.Command{
	Use:   "delete-repository-link",
	Short: "Deletes the association between your connection and a specified external Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_deleteRepositoryLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_deleteRepositoryLinkCmd).Standalone()

		codeconnections_deleteRepositoryLinkCmd.Flags().String("repository-link-id", "", "The ID of the repository link to be deleted.")
		codeconnections_deleteRepositoryLinkCmd.MarkFlagRequired("repository-link-id")
	})
	codeconnectionsCmd.AddCommand(codeconnections_deleteRepositoryLinkCmd)
}

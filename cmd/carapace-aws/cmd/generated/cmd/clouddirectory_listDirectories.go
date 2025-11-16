package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listDirectoriesCmd = &cobra.Command{
	Use:   "list-directories",
	Short: "Lists directories created within an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listDirectoriesCmd).Standalone()

	clouddirectory_listDirectoriesCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
	clouddirectory_listDirectoriesCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listDirectoriesCmd.Flags().String("state", "", "The state of the directories in the list.")
	clouddirectoryCmd.AddCommand(clouddirectory_listDirectoriesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getFolderCmd = &cobra.Command{
	Use:   "get-folder",
	Short: "Returns the contents of a specified folder in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getFolderCmd).Standalone()

	codecommit_getFolderCmd.Flags().String("commit-specifier", "", "A fully qualified reference used to identify a commit that contains the version of the folder's content to return.")
	codecommit_getFolderCmd.Flags().String("folder-path", "", "The fully qualified path to the folder whose contents are returned, including the folder name.")
	codecommit_getFolderCmd.Flags().String("repository-name", "", "The name of the repository.")
	codecommit_getFolderCmd.MarkFlagRequired("folder-path")
	codecommit_getFolderCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getFolderCmd)
}

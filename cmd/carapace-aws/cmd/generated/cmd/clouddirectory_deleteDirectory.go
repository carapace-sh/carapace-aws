package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_deleteDirectoryCmd = &cobra.Command{
	Use:   "delete-directory",
	Short: "Deletes a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_deleteDirectoryCmd).Standalone()

	clouddirectory_deleteDirectoryCmd.Flags().String("directory-arn", "", "The ARN of the directory to delete.")
	clouddirectory_deleteDirectoryCmd.MarkFlagRequired("directory-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_deleteDirectoryCmd)
}

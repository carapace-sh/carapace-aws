package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getDirectoryCmd = &cobra.Command{
	Use:   "get-directory",
	Short: "Retrieves metadata about a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getDirectoryCmd).Standalone()

	clouddirectory_getDirectoryCmd.Flags().String("directory-arn", "", "The ARN of the directory.")
	clouddirectory_getDirectoryCmd.MarkFlagRequired("directory-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_getDirectoryCmd)
}

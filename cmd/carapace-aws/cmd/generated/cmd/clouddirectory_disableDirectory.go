package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_disableDirectoryCmd = &cobra.Command{
	Use:   "disable-directory",
	Short: "Disables the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_disableDirectoryCmd).Standalone()

	clouddirectory_disableDirectoryCmd.Flags().String("directory-arn", "", "The ARN of the directory to disable.")
	clouddirectory_disableDirectoryCmd.MarkFlagRequired("directory-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_disableDirectoryCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getFileCmd = &cobra.Command{
	Use:   "get-file",
	Short: "Returns the base-64 encoded contents of a specified file and its metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getFileCmd).Standalone()

	codecommit_getFileCmd.Flags().String("commit-specifier", "", "The fully quaified reference that identifies the commit that contains the file.")
	codecommit_getFileCmd.Flags().String("file-path", "", "The fully qualified path to the file, including the full name and extension of the file.")
	codecommit_getFileCmd.Flags().String("repository-name", "", "The name of the repository that contains the file.")
	codecommit_getFileCmd.MarkFlagRequired("file-path")
	codecommit_getFileCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getFileCmd)
}

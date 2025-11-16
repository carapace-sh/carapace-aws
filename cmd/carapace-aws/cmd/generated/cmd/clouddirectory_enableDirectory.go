package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_enableDirectoryCmd = &cobra.Command{
	Use:   "enable-directory",
	Short: "Enables the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_enableDirectoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_enableDirectoryCmd).Standalone()

		clouddirectory_enableDirectoryCmd.Flags().String("directory-arn", "", "The ARN of the directory to enable.")
		clouddirectory_enableDirectoryCmd.MarkFlagRequired("directory-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_enableDirectoryCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_createDirectoryCmd = &cobra.Command{
	Use:   "create-directory",
	Short: "Creates a [Directory]() by copying the published schema into the directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_createDirectoryCmd).Standalone()

	clouddirectory_createDirectoryCmd.Flags().String("name", "", "The name of the [Directory]().")
	clouddirectory_createDirectoryCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) of the published schema that will be copied into the data [Directory]().")
	clouddirectory_createDirectoryCmd.MarkFlagRequired("name")
	clouddirectory_createDirectoryCmd.MarkFlagRequired("schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_createDirectoryCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listObjectParentPathsCmd = &cobra.Command{
	Use:   "list-object-parent-paths",
	Short: "Retrieves all available parent paths for any object type such as node, leaf node, policy node, and index node objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listObjectParentPathsCmd).Standalone()

	clouddirectory_listObjectParentPathsCmd.Flags().String("directory-arn", "", "The ARN of the directory to which the parent path applies.")
	clouddirectory_listObjectParentPathsCmd.Flags().String("max-results", "", "The maximum number of items to be retrieved in a single call.")
	clouddirectory_listObjectParentPathsCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listObjectParentPathsCmd.Flags().String("object-reference", "", "The reference that identifies the object whose parent paths are listed.")
	clouddirectory_listObjectParentPathsCmd.MarkFlagRequired("directory-arn")
	clouddirectory_listObjectParentPathsCmd.MarkFlagRequired("object-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_listObjectParentPathsCmd)
}

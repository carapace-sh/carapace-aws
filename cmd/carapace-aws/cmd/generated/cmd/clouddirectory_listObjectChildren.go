package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listObjectChildrenCmd = &cobra.Command{
	Use:   "list-object-children",
	Short: "Returns a paginated list of child objects that are associated with a given object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listObjectChildrenCmd).Standalone()

	clouddirectory_listObjectChildrenCmd.Flags().String("consistency-level", "", "Represents the manner and timing in which the successful write or update of an object is reflected in a subsequent read operation of that same object.")
	clouddirectory_listObjectChildrenCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where the object resides.")
	clouddirectory_listObjectChildrenCmd.Flags().String("max-results", "", "The maximum number of items to be retrieved in a single call.")
	clouddirectory_listObjectChildrenCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listObjectChildrenCmd.Flags().String("object-reference", "", "The reference that identifies the object for which child objects are being listed.")
	clouddirectory_listObjectChildrenCmd.MarkFlagRequired("directory-arn")
	clouddirectory_listObjectChildrenCmd.MarkFlagRequired("object-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_listObjectChildrenCmd)
}

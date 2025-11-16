package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listObjectParentsCmd = &cobra.Command{
	Use:   "list-object-parents",
	Short: "Lists parent objects that are associated with a given object in pagination fashion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listObjectParentsCmd).Standalone()

	clouddirectory_listObjectParentsCmd.Flags().String("consistency-level", "", "Represents the manner and timing in which the successful write or update of an object is reflected in a subsequent read operation of that same object.")
	clouddirectory_listObjectParentsCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where the object resides.")
	clouddirectory_listObjectParentsCmd.Flags().String("include-all-links-to-each-parent", "", "When set to True, returns all [ListObjectParentsResponse$ParentLinks]().")
	clouddirectory_listObjectParentsCmd.Flags().String("max-results", "", "The maximum number of items to be retrieved in a single call.")
	clouddirectory_listObjectParentsCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listObjectParentsCmd.Flags().String("object-reference", "", "The reference that identifies the object for which parent objects are being listed.")
	clouddirectory_listObjectParentsCmd.MarkFlagRequired("directory-arn")
	clouddirectory_listObjectParentsCmd.MarkFlagRequired("object-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_listObjectParentsCmd)
}

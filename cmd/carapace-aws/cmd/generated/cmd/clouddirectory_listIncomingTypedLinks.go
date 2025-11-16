package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listIncomingTypedLinksCmd = &cobra.Command{
	Use:   "list-incoming-typed-links",
	Short: "Returns a paginated list of all the incoming [TypedLinkSpecifier]() information for an object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listIncomingTypedLinksCmd).Standalone()

	clouddirectory_listIncomingTypedLinksCmd.Flags().String("consistency-level", "", "The consistency level to execute the request at.")
	clouddirectory_listIncomingTypedLinksCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) of the directory where you want to list the typed links.")
	clouddirectory_listIncomingTypedLinksCmd.Flags().String("filter-attribute-ranges", "", "Provides range filters for multiple attributes.")
	clouddirectory_listIncomingTypedLinksCmd.Flags().String("filter-typed-link", "", "Filters are interpreted in the order of the attributes on the typed link facet, not the order in which they are supplied to any API calls.")
	clouddirectory_listIncomingTypedLinksCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
	clouddirectory_listIncomingTypedLinksCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listIncomingTypedLinksCmd.Flags().String("object-reference", "", "Reference that identifies the object whose attributes will be listed.")
	clouddirectory_listIncomingTypedLinksCmd.MarkFlagRequired("directory-arn")
	clouddirectory_listIncomingTypedLinksCmd.MarkFlagRequired("object-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_listIncomingTypedLinksCmd)
}

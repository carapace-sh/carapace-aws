package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listOutgoingTypedLinksCmd = &cobra.Command{
	Use:   "list-outgoing-typed-links",
	Short: "Returns a paginated list of all the outgoing [TypedLinkSpecifier]() information for an object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listOutgoingTypedLinksCmd).Standalone()

	clouddirectory_listOutgoingTypedLinksCmd.Flags().String("consistency-level", "", "The consistency level to execute the request at.")
	clouddirectory_listOutgoingTypedLinksCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) of the directory where you want to list the typed links.")
	clouddirectory_listOutgoingTypedLinksCmd.Flags().String("filter-attribute-ranges", "", "Provides range filters for multiple attributes.")
	clouddirectory_listOutgoingTypedLinksCmd.Flags().String("filter-typed-link", "", "Filters are interpreted in the order of the attributes defined on the typed link facet, not the order they are supplied to any API calls.")
	clouddirectory_listOutgoingTypedLinksCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
	clouddirectory_listOutgoingTypedLinksCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listOutgoingTypedLinksCmd.Flags().String("object-reference", "", "A reference that identifies the object whose attributes will be listed.")
	clouddirectory_listOutgoingTypedLinksCmd.MarkFlagRequired("directory-arn")
	clouddirectory_listOutgoingTypedLinksCmd.MarkFlagRequired("object-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_listOutgoingTypedLinksCmd)
}

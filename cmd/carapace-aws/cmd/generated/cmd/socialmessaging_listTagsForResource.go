package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags associated with a resource, such as a phone number or WABA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_listTagsForResourceCmd).Standalone()

	socialmessaging_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to retrieve the tags from.")
	socialmessaging_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	socialmessagingCmd.AddCommand(socialmessaging_listTagsForResourceCmd)
}

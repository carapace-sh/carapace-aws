package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List the tags for an Amazon ECR Public resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_listTagsForResourceCmd).Standalone()

	ecrPublic_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource to list the tags for.")
	ecrPublic_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ecrPublicCmd.AddCommand(ecrPublic_listTagsForResourceCmd)
}

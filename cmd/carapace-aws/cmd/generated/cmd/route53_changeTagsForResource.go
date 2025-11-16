package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_changeTagsForResourceCmd = &cobra.Command{
	Use:   "change-tags-for-resource",
	Short: "Adds, edits, or deletes tags for a health check or a hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_changeTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_changeTagsForResourceCmd).Standalone()

		route53_changeTagsForResourceCmd.Flags().String("add-tags", "", "A complex type that contains a list of the tags that you want to add to the specified health check or hosted zone and/or the tags that you want to edit `Value` for.")
		route53_changeTagsForResourceCmd.Flags().String("remove-tag-keys", "", "A complex type that contains a list of the tags that you want to delete from the specified health check or hosted zone.")
		route53_changeTagsForResourceCmd.Flags().String("resource-id", "", "The ID of the resource for which you want to add, change, or delete tags.")
		route53_changeTagsForResourceCmd.Flags().String("resource-type", "", "The type of the resource.")
		route53_changeTagsForResourceCmd.MarkFlagRequired("resource-id")
		route53_changeTagsForResourceCmd.MarkFlagRequired("resource-type")
	})
	route53Cmd.AddCommand(route53_changeTagsForResourceCmd)
}

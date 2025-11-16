package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for one health check or hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listTagsForResourceCmd).Standalone()

	route53_listTagsForResourceCmd.Flags().String("resource-id", "", "The ID of the resource for which you want to retrieve tags.")
	route53_listTagsForResourceCmd.Flags().String("resource-type", "", "The type of the resource.")
	route53_listTagsForResourceCmd.MarkFlagRequired("resource-id")
	route53_listTagsForResourceCmd.MarkFlagRequired("resource-type")
	route53Cmd.AddCommand(route53_listTagsForResourceCmd)
}

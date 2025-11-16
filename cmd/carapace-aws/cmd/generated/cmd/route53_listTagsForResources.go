package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listTagsForResourcesCmd = &cobra.Command{
	Use:   "list-tags-for-resources",
	Short: "Lists tags for up to 10 health checks or hosted zones.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listTagsForResourcesCmd).Standalone()

	route53_listTagsForResourcesCmd.Flags().String("resource-ids", "", "A complex type that contains the ResourceId element for each resource for which you want to get a list of tags.")
	route53_listTagsForResourcesCmd.Flags().String("resource-type", "", "The type of the resources.")
	route53_listTagsForResourcesCmd.MarkFlagRequired("resource-ids")
	route53_listTagsForResourcesCmd.MarkFlagRequired("resource-type")
	route53Cmd.AddCommand(route53_listTagsForResourcesCmd)
}

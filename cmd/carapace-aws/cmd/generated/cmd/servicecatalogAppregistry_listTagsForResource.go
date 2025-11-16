package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags on the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_listTagsForResourceCmd).Standalone()

		servicecatalogAppregistry_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon resource name (ARN) that specifies the resource.")
		servicecatalogAppregistry_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_listTagsForResourceCmd)
}

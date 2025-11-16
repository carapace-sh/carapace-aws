package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalogAppregistry_untagResourceCmd).Standalone()

		servicecatalogAppregistry_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon resource name (ARN) that specifies the resource.")
		servicecatalogAppregistry_untagResourceCmd.Flags().String("tag-keys", "", "A list of the tag keys to remove from the specified resource.")
		servicecatalogAppregistry_untagResourceCmd.MarkFlagRequired("resource-arn")
		servicecatalogAppregistry_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_untagResourceCmd)
}

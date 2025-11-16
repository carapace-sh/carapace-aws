package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalogAppregistry_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalogAppregistry_tagResourceCmd).Standalone()

	servicecatalogAppregistry_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon resource name (ARN) that specifies the resource.")
	servicecatalogAppregistry_tagResourceCmd.Flags().String("tags", "", "The new or modified tags for the resource.")
	servicecatalogAppregistry_tagResourceCmd.MarkFlagRequired("resource-arn")
	servicecatalogAppregistry_tagResourceCmd.MarkFlagRequired("tags")
	servicecatalogAppregistryCmd.AddCommand(servicecatalogAppregistry_tagResourceCmd)
}

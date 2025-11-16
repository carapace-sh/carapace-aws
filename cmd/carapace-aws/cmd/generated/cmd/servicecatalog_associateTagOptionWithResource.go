package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_associateTagOptionWithResourceCmd = &cobra.Command{
	Use:   "associate-tag-option-with-resource",
	Short: "Associate the specified TagOption with the specified portfolio or product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_associateTagOptionWithResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_associateTagOptionWithResourceCmd).Standalone()

		servicecatalog_associateTagOptionWithResourceCmd.Flags().String("resource-id", "", "The resource identifier.")
		servicecatalog_associateTagOptionWithResourceCmd.Flags().String("tag-option-id", "", "The TagOption identifier.")
		servicecatalog_associateTagOptionWithResourceCmd.MarkFlagRequired("resource-id")
		servicecatalog_associateTagOptionWithResourceCmd.MarkFlagRequired("tag-option-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_associateTagOptionWithResourceCmd)
}

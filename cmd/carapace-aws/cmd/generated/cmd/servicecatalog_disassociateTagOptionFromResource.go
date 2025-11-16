package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_disassociateTagOptionFromResourceCmd = &cobra.Command{
	Use:   "disassociate-tag-option-from-resource",
	Short: "Disassociates the specified TagOption from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_disassociateTagOptionFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_disassociateTagOptionFromResourceCmd).Standalone()

		servicecatalog_disassociateTagOptionFromResourceCmd.Flags().String("resource-id", "", "The resource identifier.")
		servicecatalog_disassociateTagOptionFromResourceCmd.Flags().String("tag-option-id", "", "The TagOption identifier.")
		servicecatalog_disassociateTagOptionFromResourceCmd.MarkFlagRequired("resource-id")
		servicecatalog_disassociateTagOptionFromResourceCmd.MarkFlagRequired("tag-option-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_disassociateTagOptionFromResourceCmd)
}

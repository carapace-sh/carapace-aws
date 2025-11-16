package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updateTagOptionCmd = &cobra.Command{
	Use:   "update-tag-option",
	Short: "Updates the specified TagOption.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updateTagOptionCmd).Standalone()

	servicecatalog_updateTagOptionCmd.Flags().String("active", "", "The updated active state.")
	servicecatalog_updateTagOptionCmd.Flags().String("id", "", "The TagOption identifier.")
	servicecatalog_updateTagOptionCmd.Flags().String("value", "", "The updated value.")
	servicecatalog_updateTagOptionCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_updateTagOptionCmd)
}

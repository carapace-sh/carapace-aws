package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deleteTagOptionCmd = &cobra.Command{
	Use:   "delete-tag-option",
	Short: "Deletes the specified TagOption.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deleteTagOptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_deleteTagOptionCmd).Standalone()

		servicecatalog_deleteTagOptionCmd.Flags().String("id", "", "The TagOption identifier.")
		servicecatalog_deleteTagOptionCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_deleteTagOptionCmd)
}

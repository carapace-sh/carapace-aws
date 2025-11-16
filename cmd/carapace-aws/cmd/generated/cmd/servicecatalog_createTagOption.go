package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createTagOptionCmd = &cobra.Command{
	Use:   "create-tag-option",
	Short: "Creates a TagOption.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createTagOptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_createTagOptionCmd).Standalone()

		servicecatalog_createTagOptionCmd.Flags().String("key", "", "The TagOption key.")
		servicecatalog_createTagOptionCmd.Flags().String("value", "", "The TagOption value.")
		servicecatalog_createTagOptionCmd.MarkFlagRequired("key")
		servicecatalog_createTagOptionCmd.MarkFlagRequired("value")
	})
	servicecatalogCmd.AddCommand(servicecatalog_createTagOptionCmd)
}

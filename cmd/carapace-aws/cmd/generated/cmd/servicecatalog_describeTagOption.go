package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeTagOptionCmd = &cobra.Command{
	Use:   "describe-tag-option",
	Short: "Gets information about the specified TagOption.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeTagOptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_describeTagOptionCmd).Standalone()

		servicecatalog_describeTagOptionCmd.Flags().String("id", "", "The TagOption identifier.")
		servicecatalog_describeTagOptionCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_describeTagOptionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeServiceActionCmd = &cobra.Command{
	Use:   "describe-service-action",
	Short: "Describes a self-service action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeServiceActionCmd).Standalone()

	servicecatalog_describeServiceActionCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeServiceActionCmd.Flags().String("id", "", "The self-service action identifier.")
	servicecatalog_describeServiceActionCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_describeServiceActionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updateServiceActionCmd = &cobra.Command{
	Use:   "update-service-action",
	Short: "Updates a self-service action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updateServiceActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_updateServiceActionCmd).Standalone()

		servicecatalog_updateServiceActionCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_updateServiceActionCmd.Flags().String("definition", "", "A map that defines the self-service action.")
		servicecatalog_updateServiceActionCmd.Flags().String("description", "", "The self-service action description.")
		servicecatalog_updateServiceActionCmd.Flags().String("id", "", "The self-service action identifier.")
		servicecatalog_updateServiceActionCmd.Flags().String("name", "", "The self-service action name.")
		servicecatalog_updateServiceActionCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_updateServiceActionCmd)
}

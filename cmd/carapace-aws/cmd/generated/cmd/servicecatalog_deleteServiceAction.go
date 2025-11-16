package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deleteServiceActionCmd = &cobra.Command{
	Use:   "delete-service-action",
	Short: "Deletes a self-service action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deleteServiceActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_deleteServiceActionCmd).Standalone()

		servicecatalog_deleteServiceActionCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_deleteServiceActionCmd.Flags().String("id", "", "The self-service action identifier.")
		servicecatalog_deleteServiceActionCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalog_deleteServiceActionCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_deleteServiceActionCmd)
}

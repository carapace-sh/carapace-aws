package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createServiceActionCmd = &cobra.Command{
	Use:   "create-service-action",
	Short: "Creates a self-service action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createServiceActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_createServiceActionCmd).Standalone()

		servicecatalog_createServiceActionCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_createServiceActionCmd.Flags().String("definition", "", "The self-service action definition.")
		servicecatalog_createServiceActionCmd.Flags().String("definition-type", "", "The service action definition type.")
		servicecatalog_createServiceActionCmd.Flags().String("description", "", "The self-service action description.")
		servicecatalog_createServiceActionCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalog_createServiceActionCmd.Flags().String("name", "", "The self-service action name.")
		servicecatalog_createServiceActionCmd.MarkFlagRequired("definition")
		servicecatalog_createServiceActionCmd.MarkFlagRequired("definition-type")
		servicecatalog_createServiceActionCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_createServiceActionCmd.MarkFlagRequired("name")
	})
	servicecatalogCmd.AddCommand(servicecatalog_createServiceActionCmd)
}

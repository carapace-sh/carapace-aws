package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createConstraintCmd = &cobra.Command{
	Use:   "create-constraint",
	Short: "Creates a constraint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createConstraintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_createConstraintCmd).Standalone()

		servicecatalog_createConstraintCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_createConstraintCmd.Flags().String("description", "", "The description of the constraint.")
		servicecatalog_createConstraintCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalog_createConstraintCmd.Flags().String("parameters", "", "The constraint parameters, in JSON format.")
		servicecatalog_createConstraintCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
		servicecatalog_createConstraintCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_createConstraintCmd.Flags().String("type", "", "The type of constraint.")
		servicecatalog_createConstraintCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_createConstraintCmd.MarkFlagRequired("parameters")
		servicecatalog_createConstraintCmd.MarkFlagRequired("portfolio-id")
		servicecatalog_createConstraintCmd.MarkFlagRequired("product-id")
		servicecatalog_createConstraintCmd.MarkFlagRequired("type")
	})
	servicecatalogCmd.AddCommand(servicecatalog_createConstraintCmd)
}

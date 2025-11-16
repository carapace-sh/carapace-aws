package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeServiceActionExecutionParametersCmd = &cobra.Command{
	Use:   "describe-service-action-execution-parameters",
	Short: "Finds the default parameters for a specific self-service action on a specific provisioned product and returns a map of the results to the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeServiceActionExecutionParametersCmd).Standalone()

	servicecatalog_describeServiceActionExecutionParametersCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeServiceActionExecutionParametersCmd.Flags().String("provisioned-product-id", "", "The identifier of the provisioned product.")
	servicecatalog_describeServiceActionExecutionParametersCmd.Flags().String("service-action-id", "", "The self-service action identifier.")
	servicecatalog_describeServiceActionExecutionParametersCmd.MarkFlagRequired("provisioned-product-id")
	servicecatalog_describeServiceActionExecutionParametersCmd.MarkFlagRequired("service-action-id")
	servicecatalogCmd.AddCommand(servicecatalog_describeServiceActionExecutionParametersCmd)
}

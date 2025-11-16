package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateRequestValidatorCmd = &cobra.Command{
	Use:   "update-request-validator",
	Short: "Updates a RequestValidator of a given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateRequestValidatorCmd).Standalone()

	apigateway_updateRequestValidatorCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateRequestValidatorCmd.Flags().String("request-validator-id", "", "The identifier of RequestValidator to be updated.")
	apigateway_updateRequestValidatorCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateRequestValidatorCmd.MarkFlagRequired("request-validator-id")
	apigateway_updateRequestValidatorCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_updateRequestValidatorCmd)
}

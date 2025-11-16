package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteRequestValidatorCmd = &cobra.Command{
	Use:   "delete-request-validator",
	Short: "Deletes a RequestValidator of a given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteRequestValidatorCmd).Standalone()

	apigateway_deleteRequestValidatorCmd.Flags().String("request-validator-id", "", "The identifier of the RequestValidator to be deleted.")
	apigateway_deleteRequestValidatorCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_deleteRequestValidatorCmd.MarkFlagRequired("request-validator-id")
	apigateway_deleteRequestValidatorCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_deleteRequestValidatorCmd)
}

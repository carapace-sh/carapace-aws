package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createApiKeyCmd = &cobra.Command{
	Use:   "create-api-key",
	Short: "Create an ApiKey resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createApiKeyCmd).Standalone()

	apigateway_createApiKeyCmd.Flags().String("customer-id", "", "An Amazon Web Services Marketplace customer identifier, when integrating with the Amazon Web Services SaaS Marketplace.")
	apigateway_createApiKeyCmd.Flags().String("description", "", "The description of the ApiKey.")
	apigateway_createApiKeyCmd.Flags().Bool("enabled", false, "Specifies whether the ApiKey can be used by callers.")
	apigateway_createApiKeyCmd.Flags().Bool("generate-distinct-id", false, "Specifies whether (`true`) or not (`false`) the key identifier is distinct from the created API key value.")
	apigateway_createApiKeyCmd.Flags().String("name", "", "The name of the ApiKey.")
	apigateway_createApiKeyCmd.Flags().Bool("no-enabled", false, "Specifies whether the ApiKey can be used by callers.")
	apigateway_createApiKeyCmd.Flags().Bool("no-generate-distinct-id", false, "Specifies whether (`true`) or not (`false`) the key identifier is distinct from the created API key value.")
	apigateway_createApiKeyCmd.Flags().String("stage-keys", "", "DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.")
	apigateway_createApiKeyCmd.Flags().String("tags", "", "The key-value map of strings.")
	apigateway_createApiKeyCmd.Flags().String("value", "", "Specifies a value of the API key.")
	apigateway_createApiKeyCmd.Flag("no-enabled").Hidden = true
	apigateway_createApiKeyCmd.Flag("no-generate-distinct-id").Hidden = true
	apigatewayCmd.AddCommand(apigateway_createApiKeyCmd)
}

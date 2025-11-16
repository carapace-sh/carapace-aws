package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_importApiKeysCmd = &cobra.Command{
	Use:   "import-api-keys",
	Short: "Import API keys from an external source, such as a CSV-formatted file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_importApiKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_importApiKeysCmd).Standalone()

		apigateway_importApiKeysCmd.Flags().String("body", "", "The payload of the POST request to import API keys.")
		apigateway_importApiKeysCmd.Flags().Bool("fail-on-warnings", false, "A query parameter to indicate whether to rollback ApiKey importation (`true`) or not (`false`) when error is encountered.")
		apigateway_importApiKeysCmd.Flags().String("format", "", "A query parameter to specify the input format to imported API keys.")
		apigateway_importApiKeysCmd.Flags().Bool("no-fail-on-warnings", false, "A query parameter to indicate whether to rollback ApiKey importation (`true`) or not (`false`) when error is encountered.")
		apigateway_importApiKeysCmd.MarkFlagRequired("body")
		apigateway_importApiKeysCmd.MarkFlagRequired("format")
		apigateway_importApiKeysCmd.Flag("no-fail-on-warnings").Hidden = true
	})
	apigatewayCmd.AddCommand(apigateway_importApiKeysCmd)
}

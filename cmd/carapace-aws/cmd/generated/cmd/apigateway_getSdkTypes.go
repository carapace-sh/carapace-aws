package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getSdkTypesCmd = &cobra.Command{
	Use:   "get-sdk-types",
	Short: "Gets SDK types",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getSdkTypesCmd).Standalone()

	apigateway_getSdkTypesCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
	apigateway_getSdkTypesCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	apigatewayCmd.AddCommand(apigateway_getSdkTypesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getSdkTypeCmd = &cobra.Command{
	Use:   "get-sdk-type",
	Short: "Gets an SDK type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getSdkTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getSdkTypeCmd).Standalone()

		apigateway_getSdkTypeCmd.Flags().String("id", "", "The identifier of the queried SdkType instance.")
		apigateway_getSdkTypeCmd.MarkFlagRequired("id")
	})
	apigatewayCmd.AddCommand(apigateway_getSdkTypeCmd)
}

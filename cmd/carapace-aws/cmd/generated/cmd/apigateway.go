package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayCmd = &cobra.Command{
	Use:   "apigateway",
	Short: "Amazon API Gateway\n\nAmazon API Gateway helps developers deliver robust, secure, and scalable mobile and web application back ends.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayCmd).Standalone()

	})
	rootCmd.AddCommand(apigatewayCmd)
}

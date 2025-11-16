package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayCmd = &cobra.Command{
	Use:   "apigateway",
	Short: "Amazon API Gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayCmd).Standalone()

	})
	rootCmd.AddCommand(apigatewayCmd)
}

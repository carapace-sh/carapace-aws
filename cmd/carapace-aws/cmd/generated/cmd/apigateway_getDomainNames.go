package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDomainNamesCmd = &cobra.Command{
	Use:   "get-domain-names",
	Short: "Represents a collection of DomainName resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDomainNamesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getDomainNamesCmd).Standalone()

		apigateway_getDomainNamesCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getDomainNamesCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
		apigateway_getDomainNamesCmd.Flags().String("resource-owner", "", "The owner of the domain name access association.")
	})
	apigatewayCmd.AddCommand(apigateway_getDomainNamesCmd)
}

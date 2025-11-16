package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDomainNameCmd = &cobra.Command{
	Use:   "get-domain-name",
	Short: "Represents a domain name that is contained in a simpler, more intuitive URL that can be called.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDomainNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getDomainNameCmd).Standalone()

		apigateway_getDomainNameCmd.Flags().String("domain-name", "", "The name of the DomainName resource.")
		apigateway_getDomainNameCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
		apigateway_getDomainNameCmd.MarkFlagRequired("domain-name")
	})
	apigatewayCmd.AddCommand(apigateway_getDomainNameCmd)
}

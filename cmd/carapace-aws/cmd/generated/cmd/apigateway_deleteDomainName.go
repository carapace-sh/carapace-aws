package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteDomainNameCmd = &cobra.Command{
	Use:   "delete-domain-name",
	Short: "Deletes the DomainName resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteDomainNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteDomainNameCmd).Standalone()

		apigateway_deleteDomainNameCmd.Flags().String("domain-name", "", "The name of the DomainName resource to be deleted.")
		apigateway_deleteDomainNameCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
		apigateway_deleteDomainNameCmd.MarkFlagRequired("domain-name")
	})
	apigatewayCmd.AddCommand(apigateway_deleteDomainNameCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateDomainNameCmd = &cobra.Command{
	Use:   "update-domain-name",
	Short: "Changes information about the DomainName resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateDomainNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateDomainNameCmd).Standalone()

		apigateway_updateDomainNameCmd.Flags().String("domain-name", "", "The name of the DomainName resource to be changed.")
		apigateway_updateDomainNameCmd.Flags().String("domain-name-id", "", "The identifier for the domain name resource.")
		apigateway_updateDomainNameCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateDomainNameCmd.MarkFlagRequired("domain-name")
	})
	apigatewayCmd.AddCommand(apigateway_updateDomainNameCmd)
}

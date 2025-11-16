package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_updateDomainEndpointOptionsCmd = &cobra.Command{
	Use:   "update-domain-endpoint-options",
	Short: "Updates the domain's endpoint options, specifically whether all requests to the domain must arrive over HTTPS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_updateDomainEndpointOptionsCmd).Standalone()

	cloudsearch_updateDomainEndpointOptionsCmd.Flags().String("domain-endpoint-options", "", "Whether to require that all requests to the domain arrive over HTTPS.")
	cloudsearch_updateDomainEndpointOptionsCmd.Flags().String("domain-name", "", "A string that represents the name of a domain.")
	cloudsearch_updateDomainEndpointOptionsCmd.MarkFlagRequired("domain-endpoint-options")
	cloudsearch_updateDomainEndpointOptionsCmd.MarkFlagRequired("domain-name")
	cloudsearchCmd.AddCommand(cloudsearch_updateDomainEndpointOptionsCmd)
}

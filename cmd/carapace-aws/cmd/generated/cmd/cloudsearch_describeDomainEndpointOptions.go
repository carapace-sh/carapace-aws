package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeDomainEndpointOptionsCmd = &cobra.Command{
	Use:   "describe-domain-endpoint-options",
	Short: "Returns the domain's endpoint options, specifically whether all requests to the domain must arrive over HTTPS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeDomainEndpointOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_describeDomainEndpointOptionsCmd).Standalone()

		cloudsearch_describeDomainEndpointOptionsCmd.Flags().Bool("deployed", false, "Whether to retrieve the latest configuration (which might be in a Processing state) or the current, active configuration.")
		cloudsearch_describeDomainEndpointOptionsCmd.Flags().String("domain-name", "", "A string that represents the name of a domain.")
		cloudsearch_describeDomainEndpointOptionsCmd.Flags().Bool("no-deployed", false, "Whether to retrieve the latest configuration (which might be in a Processing state) or the current, active configuration.")
		cloudsearch_describeDomainEndpointOptionsCmd.MarkFlagRequired("domain-name")
		cloudsearch_describeDomainEndpointOptionsCmd.Flag("no-deployed").Hidden = true
	})
	cloudsearchCmd.AddCommand(cloudsearch_describeDomainEndpointOptionsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeDomainsCmd = &cobra.Command{
	Use:   "describe-domains",
	Short: "Gets information about the search domains owned by this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_describeDomainsCmd).Standalone()

		cloudsearch_describeDomainsCmd.Flags().String("domain-names", "", "The names of the domains you want to include in the response.")
	})
	cloudsearchCmd.AddCommand(cloudsearch_describeDomainsCmd)
}

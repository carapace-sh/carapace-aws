package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_listDomainNamesCmd = &cobra.Command{
	Use:   "list-domain-names",
	Short: "Lists all search domains owned by an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_listDomainNamesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_listDomainNamesCmd).Standalone()

	})
	cloudsearchCmd.AddCommand(cloudsearch_listDomainNamesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listDomainNamesCmd = &cobra.Command{
	Use:   "list-domain-names",
	Short: "Returns the name of all Elasticsearch domains owned by the current user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listDomainNamesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_listDomainNamesCmd).Standalone()

		es_listDomainNamesCmd.Flags().String("engine-type", "", "Optional parameter to filter the output by domain engine type.")
	})
	esCmd.AddCommand(es_listDomainNamesCmd)
}

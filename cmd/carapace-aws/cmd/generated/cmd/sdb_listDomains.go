package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "The `ListDomains` operation lists all domains associated with the Access Key ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_listDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sdb_listDomainsCmd).Standalone()

		sdb_listDomainsCmd.Flags().String("max-number-of-domains", "", "The maximum number of domain names you want returned.")
		sdb_listDomainsCmd.Flags().String("next-token", "", "A string informing Amazon SimpleDB where to start the next list of domain names.")
	})
	sdbCmd.AddCommand(sdb_listDomainsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "The `DeleteDomain` operation deletes a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_deleteDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sdb_deleteDomainCmd).Standalone()

		sdb_deleteDomainCmd.Flags().String("domain-name", "", "The name of the domain to delete.")
		sdb_deleteDomainCmd.MarkFlagRequired("domain-name")
	})
	sdbCmd.AddCommand(sdb_deleteDomainCmd)
}

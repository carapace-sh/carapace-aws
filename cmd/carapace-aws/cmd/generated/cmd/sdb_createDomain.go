package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "The `CreateDomain` operation creates a new domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_createDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sdb_createDomainCmd).Standalone()

		sdb_createDomainCmd.Flags().String("domain-name", "", "The name of the domain to create.")
		sdb_createDomainCmd.MarkFlagRequired("domain-name")
	})
	sdbCmd.AddCommand(sdb_createDomainCmd)
}

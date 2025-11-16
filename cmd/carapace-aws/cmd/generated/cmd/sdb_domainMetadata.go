package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_domainMetadataCmd = &cobra.Command{
	Use:   "domain-metadata",
	Short: "Returns information about the domain, including when the domain was created, the number of items and attributes in the domain, and the size of the attribute names and values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_domainMetadataCmd).Standalone()

	sdb_domainMetadataCmd.Flags().String("domain-name", "", "The name of the domain for which to display the metadata of.")
	sdb_domainMetadataCmd.MarkFlagRequired("domain-name")
	sdbCmd.AddCommand(sdb_domainMetadataCmd)
}

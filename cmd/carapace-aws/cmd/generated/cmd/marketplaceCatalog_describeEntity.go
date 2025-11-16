package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_describeEntityCmd = &cobra.Command{
	Use:   "describe-entity",
	Short: "Returns the metadata and content of the entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_describeEntityCmd).Standalone()

	marketplaceCatalog_describeEntityCmd.Flags().String("catalog", "", "Required.")
	marketplaceCatalog_describeEntityCmd.Flags().String("entity-id", "", "Required.")
	marketplaceCatalog_describeEntityCmd.MarkFlagRequired("catalog")
	marketplaceCatalog_describeEntityCmd.MarkFlagRequired("entity-id")
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_describeEntityCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_describeChangeSetCmd = &cobra.Command{
	Use:   "describe-change-set",
	Short: "Provides information about a given change set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_describeChangeSetCmd).Standalone()

	marketplaceCatalog_describeChangeSetCmd.Flags().String("catalog", "", "Required.")
	marketplaceCatalog_describeChangeSetCmd.Flags().String("change-set-id", "", "Required.")
	marketplaceCatalog_describeChangeSetCmd.MarkFlagRequired("catalog")
	marketplaceCatalog_describeChangeSetCmd.MarkFlagRequired("change-set-id")
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_describeChangeSetCmd)
}

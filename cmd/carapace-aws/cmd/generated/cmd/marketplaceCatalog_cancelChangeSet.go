package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_cancelChangeSetCmd = &cobra.Command{
	Use:   "cancel-change-set",
	Short: "Used to cancel an open change request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_cancelChangeSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceCatalog_cancelChangeSetCmd).Standalone()

		marketplaceCatalog_cancelChangeSetCmd.Flags().String("catalog", "", "Required.")
		marketplaceCatalog_cancelChangeSetCmd.Flags().String("change-set-id", "", "Required.")
		marketplaceCatalog_cancelChangeSetCmd.MarkFlagRequired("catalog")
		marketplaceCatalog_cancelChangeSetCmd.MarkFlagRequired("change-set-id")
	})
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_cancelChangeSetCmd)
}

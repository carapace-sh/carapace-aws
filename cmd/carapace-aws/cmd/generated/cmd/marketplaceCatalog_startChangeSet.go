package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_startChangeSetCmd = &cobra.Command{
	Use:   "start-change-set",
	Short: "Allows you to request changes for your entities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_startChangeSetCmd).Standalone()

	marketplaceCatalog_startChangeSetCmd.Flags().String("catalog", "", "The catalog related to the request.")
	marketplaceCatalog_startChangeSetCmd.Flags().String("change-set", "", "Array of `change` object.")
	marketplaceCatalog_startChangeSetCmd.Flags().String("change-set-name", "", "Optional case sensitive string of up to 100 ASCII characters.")
	marketplaceCatalog_startChangeSetCmd.Flags().String("change-set-tags", "", "A list of objects specifying each key name and value for the `ChangeSetTags` property.")
	marketplaceCatalog_startChangeSetCmd.Flags().String("client-request-token", "", "A unique token to identify the request to ensure idempotency.")
	marketplaceCatalog_startChangeSetCmd.Flags().String("intent", "", "The intent related to the request.")
	marketplaceCatalog_startChangeSetCmd.MarkFlagRequired("catalog")
	marketplaceCatalog_startChangeSetCmd.MarkFlagRequired("change-set")
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_startChangeSetCmd)
}

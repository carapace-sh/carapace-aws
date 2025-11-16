package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listInventoryEntriesCmd = &cobra.Command{
	Use:   "list-inventory-entries",
	Short: "A list of inventory items returned by the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listInventoryEntriesCmd).Standalone()

	ssm_listInventoryEntriesCmd.Flags().String("filters", "", "One or more filters.")
	ssm_listInventoryEntriesCmd.Flags().String("instance-id", "", "The managed node ID for which you want inventory information.")
	ssm_listInventoryEntriesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_listInventoryEntriesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_listInventoryEntriesCmd.Flags().String("type-name", "", "The type of inventory item for which you want information.")
	ssm_listInventoryEntriesCmd.MarkFlagRequired("instance-id")
	ssm_listInventoryEntriesCmd.MarkFlagRequired("type-name")
	ssmCmd.AddCommand(ssm_listInventoryEntriesCmd)
}

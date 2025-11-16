package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_putInventoryCmd = &cobra.Command{
	Use:   "put-inventory",
	Short: "Bulk update custom inventory items on one or more managed nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_putInventoryCmd).Standalone()

	ssm_putInventoryCmd.Flags().String("instance-id", "", "An managed node ID where you want to add or update inventory items.")
	ssm_putInventoryCmd.Flags().String("items", "", "The inventory items that you want to add or update on managed nodes.")
	ssm_putInventoryCmd.MarkFlagRequired("instance-id")
	ssm_putInventoryCmd.MarkFlagRequired("items")
	ssmCmd.AddCommand(ssm_putInventoryCmd)
}

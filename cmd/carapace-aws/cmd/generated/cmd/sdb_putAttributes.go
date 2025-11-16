package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_putAttributesCmd = &cobra.Command{
	Use:   "put-attributes",
	Short: "The PutAttributes operation creates or replaces attributes in an item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_putAttributesCmd).Standalone()

	sdb_putAttributesCmd.Flags().String("attributes", "", "The list of attributes.")
	sdb_putAttributesCmd.Flags().String("domain-name", "", "The name of the domain in which to perform the operation.")
	sdb_putAttributesCmd.Flags().String("expected", "", "The update condition which, if specified, determines whether the specified attributes will be updated or not.")
	sdb_putAttributesCmd.Flags().String("item-name", "", "The name of the item.")
	sdb_putAttributesCmd.MarkFlagRequired("attributes")
	sdb_putAttributesCmd.MarkFlagRequired("domain-name")
	sdb_putAttributesCmd.MarkFlagRequired("item-name")
	sdbCmd.AddCommand(sdb_putAttributesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_deleteAttributesCmd = &cobra.Command{
	Use:   "delete-attributes",
	Short: "Deletes one or more attributes associated with an item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_deleteAttributesCmd).Standalone()

	sdb_deleteAttributesCmd.Flags().String("attributes", "", "A list of Attributes.")
	sdb_deleteAttributesCmd.Flags().String("domain-name", "", "The name of the domain in which to perform the operation.")
	sdb_deleteAttributesCmd.Flags().String("expected", "", "The update condition which, if specified, determines whether the specified attributes will be deleted or not.")
	sdb_deleteAttributesCmd.Flags().String("item-name", "", "The name of the item.")
	sdb_deleteAttributesCmd.MarkFlagRequired("domain-name")
	sdb_deleteAttributesCmd.MarkFlagRequired("item-name")
	sdbCmd.AddCommand(sdb_deleteAttributesCmd)
}

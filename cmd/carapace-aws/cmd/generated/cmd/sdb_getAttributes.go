package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_getAttributesCmd = &cobra.Command{
	Use:   "get-attributes",
	Short: "Returns all of the attributes associated with the specified item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_getAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sdb_getAttributesCmd).Standalone()

		sdb_getAttributesCmd.Flags().String("attribute-names", "", "The names of the attributes.")
		sdb_getAttributesCmd.Flags().Bool("consistent-read", false, "Determines whether or not strong consistency should be enforced when data is read from SimpleDB.")
		sdb_getAttributesCmd.Flags().String("domain-name", "", "The name of the domain in which to perform the operation.")
		sdb_getAttributesCmd.Flags().String("item-name", "", "The name of the item.")
		sdb_getAttributesCmd.Flags().Bool("no-consistent-read", false, "Determines whether or not strong consistency should be enforced when data is read from SimpleDB.")
		sdb_getAttributesCmd.MarkFlagRequired("domain-name")
		sdb_getAttributesCmd.MarkFlagRequired("item-name")
		sdb_getAttributesCmd.Flag("no-consistent-read").Hidden = true
	})
	sdbCmd.AddCommand(sdb_getAttributesCmd)
}

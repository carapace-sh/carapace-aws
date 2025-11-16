package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_batchDeleteAttributesCmd = &cobra.Command{
	Use:   "batch-delete-attributes",
	Short: "Performs multiple DeleteAttributes operations in a single call, which reduces round trips and latencies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_batchDeleteAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sdb_batchDeleteAttributesCmd).Standalone()

		sdb_batchDeleteAttributesCmd.Flags().String("domain-name", "", "The name of the domain in which the attributes are being deleted.")
		sdb_batchDeleteAttributesCmd.Flags().String("items", "", "A list of items on which to perform the operation.")
		sdb_batchDeleteAttributesCmd.MarkFlagRequired("domain-name")
		sdb_batchDeleteAttributesCmd.MarkFlagRequired("items")
	})
	sdbCmd.AddCommand(sdb_batchDeleteAttributesCmd)
}

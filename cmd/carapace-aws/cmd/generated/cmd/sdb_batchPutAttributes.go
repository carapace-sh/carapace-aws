package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_batchPutAttributesCmd = &cobra.Command{
	Use:   "batch-put-attributes",
	Short: "The `BatchPutAttributes` operation creates or replaces attributes within one or more items.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_batchPutAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sdb_batchPutAttributesCmd).Standalone()

		sdb_batchPutAttributesCmd.Flags().String("domain-name", "", "The name of the domain in which the attributes are being stored.")
		sdb_batchPutAttributesCmd.Flags().String("items", "", "A list of items on which to perform the operation.")
		sdb_batchPutAttributesCmd.MarkFlagRequired("domain-name")
		sdb_batchPutAttributesCmd.MarkFlagRequired("items")
	})
	sdbCmd.AddCommand(sdb_batchPutAttributesCmd)
}

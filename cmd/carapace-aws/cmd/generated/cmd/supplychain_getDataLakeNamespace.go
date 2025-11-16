package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_getDataLakeNamespaceCmd = &cobra.Command{
	Use:   "get-data-lake-namespace",
	Short: "Enables you to programmatically view an Amazon Web Services Supply Chain data lake namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_getDataLakeNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_getDataLakeNamespaceCmd).Standalone()

		supplychain_getDataLakeNamespaceCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
		supplychain_getDataLakeNamespaceCmd.Flags().String("name", "", "The name of the namespace.")
		supplychain_getDataLakeNamespaceCmd.MarkFlagRequired("instance-id")
		supplychain_getDataLakeNamespaceCmd.MarkFlagRequired("name")
	})
	supplychainCmd.AddCommand(supplychain_getDataLakeNamespaceCmd)
}

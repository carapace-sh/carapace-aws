package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_deleteDataLakeNamespaceCmd = &cobra.Command{
	Use:   "delete-data-lake-namespace",
	Short: "Enables you to programmatically delete an Amazon Web Services Supply Chain data lake namespace and its underling datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_deleteDataLakeNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_deleteDataLakeNamespaceCmd).Standalone()

		supplychain_deleteDataLakeNamespaceCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
		supplychain_deleteDataLakeNamespaceCmd.Flags().String("name", "", "The name of the namespace.")
		supplychain_deleteDataLakeNamespaceCmd.MarkFlagRequired("instance-id")
		supplychain_deleteDataLakeNamespaceCmd.MarkFlagRequired("name")
	})
	supplychainCmd.AddCommand(supplychain_deleteDataLakeNamespaceCmd)
}

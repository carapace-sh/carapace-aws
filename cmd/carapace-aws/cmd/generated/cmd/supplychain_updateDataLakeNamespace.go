package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_updateDataLakeNamespaceCmd = &cobra.Command{
	Use:   "update-data-lake-namespace",
	Short: "Enables you to programmatically update an Amazon Web Services Supply Chain data lake namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_updateDataLakeNamespaceCmd).Standalone()

	supplychain_updateDataLakeNamespaceCmd.Flags().String("description", "", "The updated description of the data lake namespace.")
	supplychain_updateDataLakeNamespaceCmd.Flags().String("instance-id", "", "The Amazon Web Services Chain instance identifier.")
	supplychain_updateDataLakeNamespaceCmd.Flags().String("name", "", "The name of the namespace.")
	supplychain_updateDataLakeNamespaceCmd.MarkFlagRequired("instance-id")
	supplychain_updateDataLakeNamespaceCmd.MarkFlagRequired("name")
	supplychainCmd.AddCommand(supplychain_updateDataLakeNamespaceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_createDataLakeNamespaceCmd = &cobra.Command{
	Use:   "create-data-lake-namespace",
	Short: "Enables you to programmatically create an Amazon Web Services Supply Chain data lake namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_createDataLakeNamespaceCmd).Standalone()

	supplychain_createDataLakeNamespaceCmd.Flags().String("description", "", "The description of the namespace.")
	supplychain_createDataLakeNamespaceCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
	supplychain_createDataLakeNamespaceCmd.Flags().String("name", "", "The name of the namespace.")
	supplychain_createDataLakeNamespaceCmd.Flags().String("tags", "", "The tags of the namespace.")
	supplychain_createDataLakeNamespaceCmd.MarkFlagRequired("instance-id")
	supplychain_createDataLakeNamespaceCmd.MarkFlagRequired("name")
	supplychainCmd.AddCommand(supplychain_createDataLakeNamespaceCmd)
}

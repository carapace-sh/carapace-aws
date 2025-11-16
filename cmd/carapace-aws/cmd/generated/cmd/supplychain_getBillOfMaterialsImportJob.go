package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_getBillOfMaterialsImportJobCmd = &cobra.Command{
	Use:   "get-bill-of-materials-import-job",
	Short: "Get status and details of a BillOfMaterialsImportJob.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_getBillOfMaterialsImportJobCmd).Standalone()

	supplychain_getBillOfMaterialsImportJobCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
	supplychain_getBillOfMaterialsImportJobCmd.Flags().String("job-id", "", "The BillOfMaterialsImportJob identifier.")
	supplychain_getBillOfMaterialsImportJobCmd.MarkFlagRequired("instance-id")
	supplychain_getBillOfMaterialsImportJobCmd.MarkFlagRequired("job-id")
	supplychainCmd.AddCommand(supplychain_getBillOfMaterialsImportJobCmd)
}

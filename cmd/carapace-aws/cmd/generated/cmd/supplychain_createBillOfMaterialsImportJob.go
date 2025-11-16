package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_createBillOfMaterialsImportJobCmd = &cobra.Command{
	Use:   "create-bill-of-materials-import-job",
	Short: "CreateBillOfMaterialsImportJob creates an import job for the Product Bill Of Materials (BOM) entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_createBillOfMaterialsImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_createBillOfMaterialsImportJobCmd).Standalone()

		supplychain_createBillOfMaterialsImportJobCmd.Flags().String("client-token", "", "An idempotency token ensures the API request is only completed no more than once.")
		supplychain_createBillOfMaterialsImportJobCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
		supplychain_createBillOfMaterialsImportJobCmd.Flags().String("s3uri", "", "The S3 URI of the CSV file to be imported.")
		supplychain_createBillOfMaterialsImportJobCmd.MarkFlagRequired("instance-id")
		supplychain_createBillOfMaterialsImportJobCmd.MarkFlagRequired("s3uri")
	})
	supplychainCmd.AddCommand(supplychain_createBillOfMaterialsImportJobCmd)
}

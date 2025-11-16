package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_batchGrantPermissionsCmd = &cobra.Command{
	Use:   "batch-grant-permissions",
	Short: "Batch operation to grant permissions to the principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_batchGrantPermissionsCmd).Standalone()

	lakeformation_batchGrantPermissionsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_batchGrantPermissionsCmd.Flags().String("entries", "", "A list of up to 20 entries for resource permissions to be granted by batch operation to the principal.")
	lakeformation_batchGrantPermissionsCmd.MarkFlagRequired("entries")
	lakeformationCmd.AddCommand(lakeformation_batchGrantPermissionsCmd)
}

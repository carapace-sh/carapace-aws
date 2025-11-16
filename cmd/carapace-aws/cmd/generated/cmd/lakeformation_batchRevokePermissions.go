package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_batchRevokePermissionsCmd = &cobra.Command{
	Use:   "batch-revoke-permissions",
	Short: "Batch operation to revoke permissions from the principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_batchRevokePermissionsCmd).Standalone()

	lakeformation_batchRevokePermissionsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_batchRevokePermissionsCmd.Flags().String("entries", "", "A list of up to 20 entries for resource permissions to be revoked by batch operation to the principal.")
	lakeformation_batchRevokePermissionsCmd.MarkFlagRequired("entries")
	lakeformationCmd.AddCommand(lakeformation_batchRevokePermissionsCmd)
}

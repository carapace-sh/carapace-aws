package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_revokePermissionsCmd = &cobra.Command{
	Use:   "revoke-permissions",
	Short: "Revokes permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_revokePermissionsCmd).Standalone()

	lakeformation_revokePermissionsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_revokePermissionsCmd.Flags().String("condition", "", "")
	lakeformation_revokePermissionsCmd.Flags().String("permissions", "", "The permissions revoked to the principal on the resource.")
	lakeformation_revokePermissionsCmd.Flags().String("permissions-with-grant-option", "", "Indicates a list of permissions for which to revoke the grant option allowing the principal to pass permissions to other principals.")
	lakeformation_revokePermissionsCmd.Flags().String("principal", "", "The principal to be revoked permissions on the resource.")
	lakeformation_revokePermissionsCmd.Flags().String("resource", "", "The resource to which permissions are to be revoked.")
	lakeformation_revokePermissionsCmd.MarkFlagRequired("permissions")
	lakeformation_revokePermissionsCmd.MarkFlagRequired("principal")
	lakeformation_revokePermissionsCmd.MarkFlagRequired("resource")
	lakeformationCmd.AddCommand(lakeformation_revokePermissionsCmd)
}

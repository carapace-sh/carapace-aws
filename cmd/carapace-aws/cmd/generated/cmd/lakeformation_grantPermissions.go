package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_grantPermissionsCmd = &cobra.Command{
	Use:   "grant-permissions",
	Short: "Grants permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_grantPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_grantPermissionsCmd).Standalone()

		lakeformation_grantPermissionsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_grantPermissionsCmd.Flags().String("condition", "", "")
		lakeformation_grantPermissionsCmd.Flags().String("permissions", "", "The permissions granted to the principal on the resource.")
		lakeformation_grantPermissionsCmd.Flags().String("permissions-with-grant-option", "", "Indicates a list of the granted permissions that the principal may pass to other users.")
		lakeformation_grantPermissionsCmd.Flags().String("principal", "", "The principal to be granted the permissions on the resource.")
		lakeformation_grantPermissionsCmd.Flags().String("resource", "", "The resource to which permissions are to be granted.")
		lakeformation_grantPermissionsCmd.MarkFlagRequired("permissions")
		lakeformation_grantPermissionsCmd.MarkFlagRequired("principal")
		lakeformation_grantPermissionsCmd.MarkFlagRequired("resource")
	})
	lakeformationCmd.AddCommand(lakeformation_grantPermissionsCmd)
}

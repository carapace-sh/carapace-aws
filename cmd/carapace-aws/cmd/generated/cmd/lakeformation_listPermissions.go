package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listPermissionsCmd = &cobra.Command{
	Use:   "list-permissions",
	Short: "Returns a list of the principal permissions on the resource, filtered by the permissions of the caller.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listPermissionsCmd).Standalone()

	lakeformation_listPermissionsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_listPermissionsCmd.Flags().String("include-related", "", "Indicates that related permissions should be included in the results.")
	lakeformation_listPermissionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	lakeformation_listPermissionsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve this list.")
	lakeformation_listPermissionsCmd.Flags().String("principal", "", "Specifies a principal to filter the permissions returned.")
	lakeformation_listPermissionsCmd.Flags().String("resource", "", "A resource where you will get a list of the principal permissions.")
	lakeformation_listPermissionsCmd.Flags().String("resource-type", "", "Specifies a resource type to filter the permissions returned.")
	lakeformationCmd.AddCommand(lakeformation_listPermissionsCmd)
}

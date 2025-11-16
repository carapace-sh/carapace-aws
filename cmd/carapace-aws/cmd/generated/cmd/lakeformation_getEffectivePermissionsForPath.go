package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getEffectivePermissionsForPathCmd = &cobra.Command{
	Use:   "get-effective-permissions-for-path",
	Short: "Returns the Lake Formation permissions for a specified table or database resource located at a path in Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getEffectivePermissionsForPathCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getEffectivePermissionsForPathCmd).Standalone()

		lakeformation_getEffectivePermissionsForPathCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_getEffectivePermissionsForPathCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		lakeformation_getEffectivePermissionsForPathCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve this list.")
		lakeformation_getEffectivePermissionsForPathCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to get permissions.")
		lakeformation_getEffectivePermissionsForPathCmd.MarkFlagRequired("resource-arn")
	})
	lakeformationCmd.AddCommand(lakeformation_getEffectivePermissionsForPathCmd)
}

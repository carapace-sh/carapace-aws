package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeResourcePermissionsCmd = &cobra.Command{
	Use:   "describe-resource-permissions",
	Short: "Describes the permissions of a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeResourcePermissionsCmd).Standalone()

	workdocs_describeResourcePermissionsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_describeResourcePermissionsCmd.Flags().String("limit", "", "The maximum number of items to return with this call.")
	workdocs_describeResourcePermissionsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	workdocs_describeResourcePermissionsCmd.Flags().String("principal-id", "", "The ID of the principal to filter permissions by.")
	workdocs_describeResourcePermissionsCmd.Flags().String("resource-id", "", "The ID of the resource.")
	workdocs_describeResourcePermissionsCmd.MarkFlagRequired("resource-id")
	workdocsCmd.AddCommand(workdocs_describeResourcePermissionsCmd)
}

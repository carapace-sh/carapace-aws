package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_listPermissionsCmd = &cobra.Command{
	Use:   "list-permissions",
	Short: "Lists the users and groups who have the Grafana `Admin` and `Editor` roles in this workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_listPermissionsCmd).Standalone()

	grafana_listPermissionsCmd.Flags().String("group-id", "", "(Optional) Limits the results to only the group that matches this ID.")
	grafana_listPermissionsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	grafana_listPermissionsCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
	grafana_listPermissionsCmd.Flags().String("user-id", "", "(Optional) Limits the results to only the user that matches this ID.")
	grafana_listPermissionsCmd.Flags().String("user-type", "", "(Optional) If you specify `SSO_USER`, then only the permissions of IAM Identity Center users are returned.")
	grafana_listPermissionsCmd.Flags().String("workspace-id", "", "The ID of the workspace to list permissions for.")
	grafana_listPermissionsCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_listPermissionsCmd)
}

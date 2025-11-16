package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_updateQappPermissionsCmd = &cobra.Command{
	Use:   "update-qapp-permissions",
	Short: "Updates read permissions for a Amazon Q App in Amazon Q Business application environment instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_updateQappPermissionsCmd).Standalone()

	qapps_updateQappPermissionsCmd.Flags().String("app-id", "", "The unique identifier of the Amazon Q App for which permissions are being updated.")
	qapps_updateQappPermissionsCmd.Flags().String("grant-permissions", "", "The list of permissions to grant for the Amazon Q App.")
	qapps_updateQappPermissionsCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_updateQappPermissionsCmd.Flags().String("revoke-permissions", "", "The list of permissions to revoke for the Amazon Q App.")
	qapps_updateQappPermissionsCmd.MarkFlagRequired("app-id")
	qapps_updateQappPermissionsCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_updateQappPermissionsCmd)
}

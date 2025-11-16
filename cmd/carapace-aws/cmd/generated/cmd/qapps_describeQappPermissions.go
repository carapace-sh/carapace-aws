package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_describeQappPermissionsCmd = &cobra.Command{
	Use:   "describe-qapp-permissions",
	Short: "Describes read permissions for a Amazon Q App in Amazon Q Business application environment instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_describeQappPermissionsCmd).Standalone()

	qapps_describeQappPermissionsCmd.Flags().String("app-id", "", "The unique identifier of the Amazon Q App for which to retrieve permissions.")
	qapps_describeQappPermissionsCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_describeQappPermissionsCmd.MarkFlagRequired("app-id")
	qapps_describeQappPermissionsCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_describeQappPermissionsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getAppsListCmd = &cobra.Command{
	Use:   "get-apps-list",
	Short: "Returns information about the specified Firewall Manager applications list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getAppsListCmd).Standalone()

	fms_getAppsListCmd.Flags().Bool("default-list", false, "Specifies whether the list to retrieve is a default list owned by Firewall Manager.")
	fms_getAppsListCmd.Flags().String("list-id", "", "The ID of the Firewall Manager applications list that you want the details for.")
	fms_getAppsListCmd.Flags().Bool("no-default-list", false, "Specifies whether the list to retrieve is a default list owned by Firewall Manager.")
	fms_getAppsListCmd.MarkFlagRequired("list-id")
	fms_getAppsListCmd.Flag("no-default-list").Hidden = true
	fmsCmd.AddCommand(fms_getAppsListCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_deleteAppsListCmd = &cobra.Command{
	Use:   "delete-apps-list",
	Short: "Permanently deletes an Firewall Manager applications list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_deleteAppsListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_deleteAppsListCmd).Standalone()

		fms_deleteAppsListCmd.Flags().String("list-id", "", "The ID of the applications list that you want to delete.")
		fms_deleteAppsListCmd.MarkFlagRequired("list-id")
	})
	fmsCmd.AddCommand(fms_deleteAppsListCmd)
}

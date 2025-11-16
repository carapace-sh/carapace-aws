package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_putAppsListCmd = &cobra.Command{
	Use:   "put-apps-list",
	Short: "Creates an Firewall Manager applications list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_putAppsListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_putAppsListCmd).Standalone()

		fms_putAppsListCmd.Flags().String("apps-list", "", "The details of the Firewall Manager applications list to be created.")
		fms_putAppsListCmd.Flags().String("tag-list", "", "The tags associated with the resource.")
		fms_putAppsListCmd.MarkFlagRequired("apps-list")
	})
	fmsCmd.AddCommand(fms_putAppsListCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateServiceSyncBlockerCmd = &cobra.Command{
	Use:   "update-service-sync-blocker",
	Short: "Update the service sync blocker by resolving it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateServiceSyncBlockerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateServiceSyncBlockerCmd).Standalone()

		proton_updateServiceSyncBlockerCmd.Flags().String("id", "", "The ID of the service sync blocker.")
		proton_updateServiceSyncBlockerCmd.Flags().String("resolved-reason", "", "The reason the service sync blocker was resolved.")
		proton_updateServiceSyncBlockerCmd.MarkFlagRequired("id")
		proton_updateServiceSyncBlockerCmd.MarkFlagRequired("resolved-reason")
	})
	protonCmd.AddCommand(proton_updateServiceSyncBlockerCmd)
}

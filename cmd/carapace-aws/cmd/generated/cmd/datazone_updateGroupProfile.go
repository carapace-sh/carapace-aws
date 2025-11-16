package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateGroupProfileCmd = &cobra.Command{
	Use:   "update-group-profile",
	Short: "Updates the specified group profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateGroupProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateGroupProfileCmd).Standalone()

		datazone_updateGroupProfileCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a group profile is updated.")
		datazone_updateGroupProfileCmd.Flags().String("group-identifier", "", "The identifier of the group profile that is updated.")
		datazone_updateGroupProfileCmd.Flags().String("status", "", "The status of the group profile that is updated.")
		datazone_updateGroupProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_updateGroupProfileCmd.MarkFlagRequired("group-identifier")
		datazone_updateGroupProfileCmd.MarkFlagRequired("status")
	})
	datazoneCmd.AddCommand(datazone_updateGroupProfileCmd)
}

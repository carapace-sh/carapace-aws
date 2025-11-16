package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteMlconfigurationCmd = &cobra.Command{
	Use:   "delete-mlconfiguration",
	Short: "Deletes a ML modeling configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteMlconfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteMlconfigurationCmd).Standalone()

		cleanroomsml_deleteMlconfigurationCmd.Flags().String("membership-identifier", "", "The membership ID of the of the member that is deleting the ML modeling configuration.")
		cleanroomsml_deleteMlconfigurationCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteMlconfigurationCmd)
}

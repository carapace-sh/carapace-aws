package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubConfig_deleteHomeRegionControlCmd = &cobra.Command{
	Use:   "delete-home-region-control",
	Short: "This operation deletes the home region configuration for the calling account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubConfig_deleteHomeRegionControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubConfig_deleteHomeRegionControlCmd).Standalone()

		migrationhubConfig_deleteHomeRegionControlCmd.Flags().String("control-id", "", "A unique identifier that's generated for each home region control.")
		migrationhubConfig_deleteHomeRegionControlCmd.MarkFlagRequired("control-id")
	})
	migrationhubConfigCmd.AddCommand(migrationhubConfig_deleteHomeRegionControlCmd)
}

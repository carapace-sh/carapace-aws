package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubConfig_createHomeRegionControlCmd = &cobra.Command{
	Use:   "create-home-region-control",
	Short: "This API sets up the home region for the calling account only.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubConfig_createHomeRegionControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubConfig_createHomeRegionControlCmd).Standalone()

		migrationhubConfig_createHomeRegionControlCmd.Flags().String("dry-run", "", "Optional Boolean flag to indicate whether any effect should take place.")
		migrationhubConfig_createHomeRegionControlCmd.Flags().String("home-region", "", "The name of the home region of the calling account.")
		migrationhubConfig_createHomeRegionControlCmd.Flags().String("target", "", "The account for which this command sets up a home region control.")
		migrationhubConfig_createHomeRegionControlCmd.MarkFlagRequired("home-region")
		migrationhubConfig_createHomeRegionControlCmd.MarkFlagRequired("target")
	})
	migrationhubConfigCmd.AddCommand(migrationhubConfig_createHomeRegionControlCmd)
}

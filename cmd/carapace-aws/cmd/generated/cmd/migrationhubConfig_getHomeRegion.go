package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubConfig_getHomeRegionCmd = &cobra.Command{
	Use:   "get-home-region",
	Short: "Returns the calling account’s home region, if configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubConfig_getHomeRegionCmd).Standalone()

	migrationhubConfigCmd.AddCommand(migrationhubConfig_getHomeRegionCmd)
}

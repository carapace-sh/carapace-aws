package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubConfig_describeHomeRegionControlsCmd = &cobra.Command{
	Use:   "describe-home-region-controls",
	Short: "This API permits filtering on the `ControlId` and `HomeRegion` fields.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubConfig_describeHomeRegionControlsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubConfig_describeHomeRegionControlsCmd).Standalone()

		migrationhubConfig_describeHomeRegionControlsCmd.Flags().String("control-id", "", "The `ControlID` is a unique identifier string of your `HomeRegionControl` object.")
		migrationhubConfig_describeHomeRegionControlsCmd.Flags().String("home-region", "", "The name of the home region you'd like to view.")
		migrationhubConfig_describeHomeRegionControlsCmd.Flags().String("max-results", "", "The maximum number of filtering results to display per page.")
		migrationhubConfig_describeHomeRegionControlsCmd.Flags().String("next-token", "", "If a `NextToken` was returned by a previous call, more results are available.")
		migrationhubConfig_describeHomeRegionControlsCmd.Flags().String("target", "", "The target parameter specifies the identifier to which the home region is applied, which is always of type `ACCOUNT`.")
	})
	migrationhubConfigCmd.AddCommand(migrationhubConfig_describeHomeRegionControlsCmd)
}

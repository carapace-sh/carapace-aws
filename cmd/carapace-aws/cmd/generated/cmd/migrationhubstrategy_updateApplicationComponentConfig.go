package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_updateApplicationComponentConfigCmd = &cobra.Command{
	Use:   "update-application-component-config",
	Short: "Updates the configuration of an application component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_updateApplicationComponentConfigCmd).Standalone()

	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().String("app-type", "", "The type of known component.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().String("application-component-id", "", "The ID of the application component.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().Bool("configure-only", false, "Update the configuration request of an application component.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().String("inclusion-status", "", "Indicates whether the application component has been included for server recommendation or not.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().Bool("no-configure-only", false, "Update the configuration request of an application component.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().String("secrets-manager-key", "", "Database credentials.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().String("source-code-list", "", "The list of source code configurations to update for the application component.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flags().String("strategy-option", "", "The preferred strategy options for the application component.")
	migrationhubstrategy_updateApplicationComponentConfigCmd.MarkFlagRequired("application-component-id")
	migrationhubstrategy_updateApplicationComponentConfigCmd.Flag("no-configure-only").Hidden = true
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_updateApplicationComponentConfigCmd)
}

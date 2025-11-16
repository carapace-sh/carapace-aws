package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_notifyApplicationStateCmd = &cobra.Command{
	Use:   "notify-application-state",
	Short: "Sets the migration state of an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_notifyApplicationStateCmd).Standalone()

	mgh_notifyApplicationStateCmd.Flags().String("application-id", "", "The configurationId in Application Discovery Service that uniquely identifies the grouped application.")
	mgh_notifyApplicationStateCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
	mgh_notifyApplicationStateCmd.Flags().String("status", "", "Status of the application - Not Started, In-Progress, Complete.")
	mgh_notifyApplicationStateCmd.Flags().String("update-date-time", "", "The timestamp when the application state changed.")
	mgh_notifyApplicationStateCmd.MarkFlagRequired("application-id")
	mgh_notifyApplicationStateCmd.MarkFlagRequired("status")
	mghCmd.AddCommand(mgh_notifyApplicationStateCmd)
}

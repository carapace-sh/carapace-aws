package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteAlarmCmd = &cobra.Command{
	Use:   "delete-alarm",
	Short: "Deletes an alarm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteAlarmCmd).Standalone()

	lightsail_deleteAlarmCmd.Flags().String("alarm-name", "", "The name of the alarm to delete.")
	lightsail_deleteAlarmCmd.MarkFlagRequired("alarm-name")
	lightsailCmd.AddCommand(lightsail_deleteAlarmCmd)
}

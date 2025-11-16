package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_describeViewCmd = &cobra.Command{
	Use:   "describe-view",
	Short: "Retrieves the view for the specified view token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_describeViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectparticipant_describeViewCmd).Standalone()

		connectparticipant_describeViewCmd.Flags().String("connection-token", "", "The connection token.")
		connectparticipant_describeViewCmd.Flags().String("view-token", "", "An encrypted token originating from the interactive message of a ShowView block operation.")
		connectparticipant_describeViewCmd.MarkFlagRequired("connection-token")
		connectparticipant_describeViewCmd.MarkFlagRequired("view-token")
	})
	connectparticipantCmd.AddCommand(connectparticipant_describeViewCmd)
}

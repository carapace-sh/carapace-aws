package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_listRoomsCmd = &cobra.Command{
	Use:   "list-rooms",
	Short: "Gets summary information about all your rooms in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_listRoomsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_listRoomsCmd).Standalone()

		ivschat_listRoomsCmd.Flags().String("logging-configuration-identifier", "", "Logging-configuration identifier.")
		ivschat_listRoomsCmd.Flags().String("max-results", "", "Maximum number of rooms to return.")
		ivschat_listRoomsCmd.Flags().String("message-review-handler-uri", "", "Filters the list to match the specified message review handler URI.")
		ivschat_listRoomsCmd.Flags().String("name", "", "Filters the list to match the specified room name.")
		ivschat_listRoomsCmd.Flags().String("next-token", "", "The first room to retrieve.")
	})
	ivschatCmd.AddCommand(ivschat_listRoomsCmd)
}

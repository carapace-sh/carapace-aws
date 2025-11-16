package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_deleteRoomCmd = &cobra.Command{
	Use:   "delete-room",
	Short: "Deletes the specified room.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_deleteRoomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_deleteRoomCmd).Standalone()

		ivschat_deleteRoomCmd.Flags().String("identifier", "", "Identifier of the room to be deleted.")
		ivschat_deleteRoomCmd.MarkFlagRequired("identifier")
	})
	ivschatCmd.AddCommand(ivschat_deleteRoomCmd)
}

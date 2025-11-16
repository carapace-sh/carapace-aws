package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteRemoteAccessSessionCmd = &cobra.Command{
	Use:   "delete-remote-access-session",
	Short: "Deletes a completed remote access session and its results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteRemoteAccessSessionCmd).Standalone()

	devicefarm_deleteRemoteAccessSessionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the session for which you want to delete remote access.")
	devicefarm_deleteRemoteAccessSessionCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_deleteRemoteAccessSessionCmd)
}

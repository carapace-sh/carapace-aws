package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_stopRemoteAccessSessionCmd = &cobra.Command{
	Use:   "stop-remote-access-session",
	Short: "Ends a specified remote access session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_stopRemoteAccessSessionCmd).Standalone()

	devicefarm_stopRemoteAccessSessionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the remote access session to stop.")
	devicefarm_stopRemoteAccessSessionCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_stopRemoteAccessSessionCmd)
}

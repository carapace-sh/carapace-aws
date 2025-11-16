package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getRemoteAccessSessionCmd = &cobra.Command{
	Use:   "get-remote-access-session",
	Short: "Returns a link to a currently running remote access session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getRemoteAccessSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getRemoteAccessSessionCmd).Standalone()

		devicefarm_getRemoteAccessSessionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the remote access session about which you want to get session information.")
		devicefarm_getRemoteAccessSessionCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_getRemoteAccessSessionCmd)
}

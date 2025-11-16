package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_installToRemoteAccessSessionCmd = &cobra.Command{
	Use:   "install-to-remote-access-session",
	Short: "Installs an application to the device in a remote access session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_installToRemoteAccessSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_installToRemoteAccessSessionCmd).Standalone()

		devicefarm_installToRemoteAccessSessionCmd.Flags().String("app-arn", "", "The ARN of the app about which you are requesting information.")
		devicefarm_installToRemoteAccessSessionCmd.Flags().String("remote-access-session-arn", "", "The Amazon Resource Name (ARN) of the remote access session about which you are requesting information.")
		devicefarm_installToRemoteAccessSessionCmd.MarkFlagRequired("app-arn")
		devicefarm_installToRemoteAccessSessionCmd.MarkFlagRequired("remote-access-session-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_installToRemoteAccessSessionCmd)
}

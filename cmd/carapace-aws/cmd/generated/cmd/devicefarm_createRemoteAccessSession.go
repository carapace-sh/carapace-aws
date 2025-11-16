package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createRemoteAccessSessionCmd = &cobra.Command{
	Use:   "create-remote-access-session",
	Short: "Specifies and starts a remote access session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createRemoteAccessSessionCmd).Standalone()

	devicefarm_createRemoteAccessSessionCmd.Flags().String("app-arn", "", "The Amazon Resource Name (ARN) of the app to create the remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("client-id", "", "Unique identifier for the client.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("configuration", "", "The configuration information for the remote access session request.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("device-arn", "", "The ARN of the device for which you want to create a remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("instance-arn", "", "The Amazon Resource Name (ARN) of the device instance for which you want to create a remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("interaction-mode", "", "The interaction mode of the remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("name", "", "The name of the remote access session to create.")
	devicefarm_createRemoteAccessSessionCmd.Flags().Bool("no-remote-debug-enabled", false, "Set to `true` if you want to access devices remotely for debugging in your remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().Bool("no-remote-record-enabled", false, "Set to `true` to enable remote recording for the remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().Bool("no-skip-app-resign", false, "When set to `true`, for private devices, Device Farm does not sign your app again.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("project-arn", "", "The Amazon Resource Name (ARN) of the project for which you want to create a remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().Bool("remote-debug-enabled", false, "Set to `true` if you want to access devices remotely for debugging in your remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("remote-record-app-arn", "", "The Amazon Resource Name (ARN) for the app to be recorded in the remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().Bool("remote-record-enabled", false, "Set to `true` to enable remote recording for the remote access session.")
	devicefarm_createRemoteAccessSessionCmd.Flags().Bool("skip-app-resign", false, "When set to `true`, for private devices, Device Farm does not sign your app again.")
	devicefarm_createRemoteAccessSessionCmd.Flags().String("ssh-public-key", "", "Ignored.")
	devicefarm_createRemoteAccessSessionCmd.MarkFlagRequired("device-arn")
	devicefarm_createRemoteAccessSessionCmd.Flag("no-remote-debug-enabled").Hidden = true
	devicefarm_createRemoteAccessSessionCmd.Flag("no-remote-record-enabled").Hidden = true
	devicefarm_createRemoteAccessSessionCmd.Flag("no-skip-app-resign").Hidden = true
	devicefarm_createRemoteAccessSessionCmd.MarkFlagRequired("project-arn")
	devicefarmCmd.AddCommand(devicefarm_createRemoteAccessSessionCmd)
}

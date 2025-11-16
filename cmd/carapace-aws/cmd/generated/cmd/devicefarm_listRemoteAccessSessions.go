package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listRemoteAccessSessionsCmd = &cobra.Command{
	Use:   "list-remote-access-sessions",
	Short: "Returns a list of all currently running remote access sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listRemoteAccessSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listRemoteAccessSessionsCmd).Standalone()

		devicefarm_listRemoteAccessSessionsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the project about which you are requesting information.")
		devicefarm_listRemoteAccessSessionsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
		devicefarm_listRemoteAccessSessionsCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listRemoteAccessSessionsCmd)
}

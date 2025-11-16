package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_getDeviceCmd = &cobra.Command{
	Use:   "get-device",
	Short: "Returns information for a thin client device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_getDeviceCmd).Standalone()

	workspacesThinClient_getDeviceCmd.Flags().String("id", "", "The ID of the device for which to return information.")
	workspacesThinClient_getDeviceCmd.MarkFlagRequired("id")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_getDeviceCmd)
}

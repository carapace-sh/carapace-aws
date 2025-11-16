package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listNetworkProfilesCmd = &cobra.Command{
	Use:   "list-network-profiles",
	Short: "Returns the list of available network profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listNetworkProfilesCmd).Standalone()

	devicefarm_listNetworkProfilesCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the project for which you want to list network profiles.")
	devicefarm_listNetworkProfilesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarm_listNetworkProfilesCmd.Flags().String("type", "", "The type of network profile to return information about.")
	devicefarm_listNetworkProfilesCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_listNetworkProfilesCmd)
}

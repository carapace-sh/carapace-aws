package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listEnvironmentAccountConnectionsCmd = &cobra.Command{
	Use:   "list-environment-account-connections",
	Short: "View a list of environment account connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listEnvironmentAccountConnectionsCmd).Standalone()

	proton_listEnvironmentAccountConnectionsCmd.Flags().String("environment-name", "", "The environment name that's associated with each listed environment account connection.")
	proton_listEnvironmentAccountConnectionsCmd.Flags().String("max-results", "", "The maximum number of environment account connections to list.")
	proton_listEnvironmentAccountConnectionsCmd.Flags().String("next-token", "", "A token that indicates the location of the next environment account connection in the array of environment account connections, after the list of environment account connections that was previously requested.")
	proton_listEnvironmentAccountConnectionsCmd.Flags().String("requested-by", "", "The type of account making the `ListEnvironmentAccountConnections` request.")
	proton_listEnvironmentAccountConnectionsCmd.Flags().String("statuses", "", "The status details for each listed environment account connection.")
	proton_listEnvironmentAccountConnectionsCmd.MarkFlagRequired("requested-by")
	protonCmd.AddCommand(proton_listEnvironmentAccountConnectionsCmd)
}

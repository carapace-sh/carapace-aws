package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_createEnvironmentHostCmd = &cobra.Command{
	Use:   "create-environment-host",
	Short: "Creates an ESXi host and adds it to an Amazon EVS environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_createEnvironmentHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_createEnvironmentHostCmd).Standalone()

		evs_createEnvironmentHostCmd.Flags().String("client-token", "", "This parameter is not used in Amazon EVS currently.")
		evs_createEnvironmentHostCmd.Flags().String("environment-id", "", "A unique ID for the environment that the host is added to.")
		evs_createEnvironmentHostCmd.Flags().String("host", "", "The host that is created and added to the environment.")
		evs_createEnvironmentHostCmd.MarkFlagRequired("environment-id")
		evs_createEnvironmentHostCmd.MarkFlagRequired("host")
	})
	evsCmd.AddCommand(evs_createEnvironmentHostCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_deleteEnvironmentHostCmd = &cobra.Command{
	Use:   "delete-environment-host",
	Short: "Deletes a host from an Amazon EVS environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_deleteEnvironmentHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_deleteEnvironmentHostCmd).Standalone()

		evs_deleteEnvironmentHostCmd.Flags().String("client-token", "", "This parameter is not used in Amazon EVS currently.")
		evs_deleteEnvironmentHostCmd.Flags().String("environment-id", "", "A unique ID for the host's environment.")
		evs_deleteEnvironmentHostCmd.Flags().String("host-name", "", "The DNS hostname associated with the host to be deleted.")
		evs_deleteEnvironmentHostCmd.MarkFlagRequired("environment-id")
		evs_deleteEnvironmentHostCmd.MarkFlagRequired("host-name")
	})
	evsCmd.AddCommand(evs_deleteEnvironmentHostCmd)
}

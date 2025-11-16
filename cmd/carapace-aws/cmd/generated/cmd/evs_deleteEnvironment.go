package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes an Amazon EVS environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_deleteEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_deleteEnvironmentCmd).Standalone()

		evs_deleteEnvironmentCmd.Flags().String("client-token", "", "This parameter is not used in Amazon EVS currently.")
		evs_deleteEnvironmentCmd.Flags().String("environment-id", "", "A unique ID associated with the environment to be deleted.")
		evs_deleteEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	evsCmd.AddCommand(evs_deleteEnvironmentCmd)
}

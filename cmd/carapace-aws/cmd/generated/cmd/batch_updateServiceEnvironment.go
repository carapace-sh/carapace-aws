package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_updateServiceEnvironmentCmd = &cobra.Command{
	Use:   "update-service-environment",
	Short: "Updates a service environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_updateServiceEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_updateServiceEnvironmentCmd).Standalone()

		batch_updateServiceEnvironmentCmd.Flags().String("capacity-limits", "", "The capacity limits for the service environment.")
		batch_updateServiceEnvironmentCmd.Flags().String("service-environment", "", "The name or ARN of the service environment to update.")
		batch_updateServiceEnvironmentCmd.Flags().String("state", "", "The state of the service environment.")
		batch_updateServiceEnvironmentCmd.MarkFlagRequired("service-environment")
	})
	batchCmd.AddCommand(batch_updateServiceEnvironmentCmd)
}

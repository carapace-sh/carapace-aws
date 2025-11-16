package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_createServiceEnvironmentCmd = &cobra.Command{
	Use:   "create-service-environment",
	Short: "Creates a service environment for running service jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_createServiceEnvironmentCmd).Standalone()

	batch_createServiceEnvironmentCmd.Flags().String("capacity-limits", "", "The capacity limits for the service environment.")
	batch_createServiceEnvironmentCmd.Flags().String("service-environment-name", "", "The name for the service environment.")
	batch_createServiceEnvironmentCmd.Flags().String("service-environment-type", "", "The type of service environment.")
	batch_createServiceEnvironmentCmd.Flags().String("state", "", "The state of the service environment.")
	batch_createServiceEnvironmentCmd.Flags().String("tags", "", "The tags that you apply to the service environment to help you categorize and organize your resources.")
	batch_createServiceEnvironmentCmd.MarkFlagRequired("capacity-limits")
	batch_createServiceEnvironmentCmd.MarkFlagRequired("service-environment-name")
	batch_createServiceEnvironmentCmd.MarkFlagRequired("service-environment-type")
	batchCmd.AddCommand(batch_createServiceEnvironmentCmd)
}

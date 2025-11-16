package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_deleteServiceEnvironmentCmd = &cobra.Command{
	Use:   "delete-service-environment",
	Short: "Deletes a Service environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_deleteServiceEnvironmentCmd).Standalone()

	batch_deleteServiceEnvironmentCmd.Flags().String("service-environment", "", "The name or ARN of the service environment to delete.")
	batch_deleteServiceEnvironmentCmd.MarkFlagRequired("service-environment")
	batchCmd.AddCommand(batch_deleteServiceEnvironmentCmd)
}

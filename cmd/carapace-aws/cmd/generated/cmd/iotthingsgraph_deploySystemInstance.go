package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_deploySystemInstanceCmd = &cobra.Command{
	Use:   "deploy-system-instance",
	Short: "**Greengrass and Cloud Deployments**\n\nDeploys the system instance to the target specified in `CreateSystemInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_deploySystemInstanceCmd).Standalone()

	iotthingsgraph_deploySystemInstanceCmd.Flags().String("id", "", "The ID of the system instance.")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_deploySystemInstanceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_undeploySystemInstanceCmd = &cobra.Command{
	Use:   "undeploy-system-instance",
	Short: "Removes a system instance from its target (Cloud or Greengrass).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_undeploySystemInstanceCmd).Standalone()

	iotthingsgraph_undeploySystemInstanceCmd.Flags().String("id", "", "The ID of the system instance to remove from its target.")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_undeploySystemInstanceCmd)
}

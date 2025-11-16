package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_startDeleteMonitorDeploymentCmd = &cobra.Command{
	Use:   "start-delete-monitor-deployment",
	Short: "Initiates a deployment to delete the monitor of the specified signal map.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_startDeleteMonitorDeploymentCmd).Standalone()

	medialive_startDeleteMonitorDeploymentCmd.Flags().String("identifier", "", "A signal map's identifier.")
	medialive_startDeleteMonitorDeploymentCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_startDeleteMonitorDeploymentCmd)
}

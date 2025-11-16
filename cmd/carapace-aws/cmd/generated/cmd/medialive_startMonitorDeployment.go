package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_startMonitorDeploymentCmd = &cobra.Command{
	Use:   "start-monitor-deployment",
	Short: "Initiates a deployment to deploy the latest monitor of the specified signal map.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_startMonitorDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_startMonitorDeploymentCmd).Standalone()

		medialive_startMonitorDeploymentCmd.Flags().String("dry-run", "", "")
		medialive_startMonitorDeploymentCmd.Flags().String("identifier", "", "A signal map's identifier.")
		medialive_startMonitorDeploymentCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_startMonitorDeploymentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_skipWaitTimeForInstanceTerminationCmd = &cobra.Command{
	Use:   "skip-wait-time-for-instance-termination",
	Short: "In a blue/green deployment, overrides any specified wait time and starts terminating instances immediately after the traffic routing is complete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_skipWaitTimeForInstanceTerminationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_skipWaitTimeForInstanceTerminationCmd).Standalone()

		codedeploy_skipWaitTimeForInstanceTerminationCmd.Flags().String("deployment-id", "", "The unique ID of a blue/green deployment for which you want to skip the instance termination wait time.")
	})
	codedeployCmd.AddCommand(codedeploy_skipWaitTimeForInstanceTerminationCmd)
}

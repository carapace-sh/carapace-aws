package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_terminateTargetInstancesCmd = &cobra.Command{
	Use:   "terminate-target-instances",
	Short: "Starts a job that terminates specific launched EC2 Test and Cutover instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_terminateTargetInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_terminateTargetInstancesCmd).Standalone()

		mgn_terminateTargetInstancesCmd.Flags().String("account-id", "", "Terminate Target instance by Account ID")
		mgn_terminateTargetInstancesCmd.Flags().String("source-server-ids", "", "Terminate Target instance by Source Server IDs.")
		mgn_terminateTargetInstancesCmd.Flags().String("tags", "", "Terminate Target instance by Tags.")
		mgn_terminateTargetInstancesCmd.MarkFlagRequired("source-server-ids")
	})
	mgnCmd.AddCommand(mgn_terminateTargetInstancesCmd)
}

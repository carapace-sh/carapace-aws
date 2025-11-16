package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_terminateRecoveryInstancesCmd = &cobra.Command{
	Use:   "terminate-recovery-instances",
	Short: "Initiates a Job for terminating the EC2 resources associated with the specified Recovery Instances, and then will delete the Recovery Instances from the Elastic Disaster Recovery service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_terminateRecoveryInstancesCmd).Standalone()

	drs_terminateRecoveryInstancesCmd.Flags().String("recovery-instance-ids", "", "The IDs of the Recovery Instances that should be terminated.")
	drs_terminateRecoveryInstancesCmd.MarkFlagRequired("recovery-instance-ids")
	drsCmd.AddCommand(drs_terminateRecoveryInstancesCmd)
}

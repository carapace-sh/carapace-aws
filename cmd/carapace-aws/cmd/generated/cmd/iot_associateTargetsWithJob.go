package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_associateTargetsWithJobCmd = &cobra.Command{
	Use:   "associate-targets-with-job",
	Short: "Associates a group with a continuous job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_associateTargetsWithJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_associateTargetsWithJobCmd).Standalone()

		iot_associateTargetsWithJobCmd.Flags().String("comment", "", "An optional comment string describing why the job was associated with the targets.")
		iot_associateTargetsWithJobCmd.Flags().String("job-id", "", "The unique identifier you assigned to this job when it was created.")
		iot_associateTargetsWithJobCmd.Flags().String("namespace-id", "", "The namespace used to indicate that a job is a customer-managed job.")
		iot_associateTargetsWithJobCmd.Flags().String("targets", "", "A list of thing group ARNs that define the targets of the job.")
		iot_associateTargetsWithJobCmd.MarkFlagRequired("job-id")
		iot_associateTargetsWithJobCmd.MarkFlagRequired("targets")
	})
	iotCmd.AddCommand(iot_associateTargetsWithJobCmd)
}

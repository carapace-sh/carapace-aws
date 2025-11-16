package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeJobCmd = &cobra.Command{
	Use:   "describe-job",
	Short: "Describes a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeJobCmd).Standalone()

		iot_describeJobCmd.Flags().String("before-substitution", "", "Provides a view of the job document before and after the substitution parameters have been resolved with their exact values.")
		iot_describeJobCmd.Flags().String("job-id", "", "The unique identifier you assigned to this job when it was created.")
		iot_describeJobCmd.MarkFlagRequired("job-id")
	})
	iotCmd.AddCommand(iot_describeJobCmd)
}

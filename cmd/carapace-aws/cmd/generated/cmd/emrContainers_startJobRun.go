package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_startJobRunCmd = &cobra.Command{
	Use:   "start-job-run",
	Short: "Starts a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_startJobRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_startJobRunCmd).Standalone()

		emrContainers_startJobRunCmd.Flags().String("client-token", "", "The client idempotency token of the job run request.")
		emrContainers_startJobRunCmd.Flags().String("configuration-overrides", "", "The configuration overrides for the job run.")
		emrContainers_startJobRunCmd.Flags().String("execution-role-arn", "", "The execution role ARN for the job run.")
		emrContainers_startJobRunCmd.Flags().String("job-driver", "", "The job driver for the job run.")
		emrContainers_startJobRunCmd.Flags().String("job-template-id", "", "The job template ID to be used to start the job run.")
		emrContainers_startJobRunCmd.Flags().String("job-template-parameters", "", "The values of job template parameters to start a job run.")
		emrContainers_startJobRunCmd.Flags().String("name", "", "The name of the job run.")
		emrContainers_startJobRunCmd.Flags().String("release-label", "", "The Amazon EMR release version to use for the job run.")
		emrContainers_startJobRunCmd.Flags().String("retry-policy-configuration", "", "The retry policy configuration for the job run.")
		emrContainers_startJobRunCmd.Flags().String("tags", "", "The tags assigned to job runs.")
		emrContainers_startJobRunCmd.Flags().String("virtual-cluster-id", "", "The virtual cluster ID for which the job run request is submitted.")
		emrContainers_startJobRunCmd.MarkFlagRequired("client-token")
		emrContainers_startJobRunCmd.MarkFlagRequired("virtual-cluster-id")
	})
	emrContainersCmd.AddCommand(emrContainers_startJobRunCmd)
}

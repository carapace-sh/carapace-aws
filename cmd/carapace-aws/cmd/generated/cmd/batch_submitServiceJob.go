package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_submitServiceJobCmd = &cobra.Command{
	Use:   "submit-service-job",
	Short: "Submits a service job to a specified job queue to run on SageMaker AI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_submitServiceJobCmd).Standalone()

	batch_submitServiceJobCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	batch_submitServiceJobCmd.Flags().String("job-name", "", "The name of the service job.")
	batch_submitServiceJobCmd.Flags().String("job-queue", "", "The job queue into which the service job is submitted.")
	batch_submitServiceJobCmd.Flags().String("retry-strategy", "", "The retry strategy to use for failed service jobs that are submitted with this service job request.")
	batch_submitServiceJobCmd.Flags().String("scheduling-priority", "", "The scheduling priority of the service job.")
	batch_submitServiceJobCmd.Flags().String("service-job-type", "", "The type of service job.")
	batch_submitServiceJobCmd.Flags().String("service-request-payload", "", "The request, in JSON, for the service that the SubmitServiceJob operation is queueing.")
	batch_submitServiceJobCmd.Flags().String("share-identifier", "", "The share identifier for the service job.")
	batch_submitServiceJobCmd.Flags().String("tags", "", "The tags that you apply to the service job request.")
	batch_submitServiceJobCmd.Flags().String("timeout-config", "", "The timeout configuration for the service job.")
	batch_submitServiceJobCmd.MarkFlagRequired("job-name")
	batch_submitServiceJobCmd.MarkFlagRequired("job-queue")
	batch_submitServiceJobCmd.MarkFlagRequired("service-job-type")
	batch_submitServiceJobCmd.MarkFlagRequired("service-request-payload")
	batchCmd.AddCommand(batch_submitServiceJobCmd)
}

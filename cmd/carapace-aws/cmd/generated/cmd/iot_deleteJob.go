package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteJobCmd = &cobra.Command{
	Use:   "delete-job",
	Short: "Deletes a job and its related job executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteJobCmd).Standalone()

	iot_deleteJobCmd.Flags().String("force", "", "(Optional) When true, you can delete a job which is \"IN\\_PROGRESS\".")
	iot_deleteJobCmd.Flags().String("job-id", "", "The ID of the job to be deleted.")
	iot_deleteJobCmd.Flags().String("namespace-id", "", "The namespace used to indicate that a job is a customer-managed job.")
	iot_deleteJobCmd.MarkFlagRequired("job-id")
	iotCmd.AddCommand(iot_deleteJobCmd)
}

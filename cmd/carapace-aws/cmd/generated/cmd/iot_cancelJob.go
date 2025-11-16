package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "Cancels a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_cancelJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_cancelJobCmd).Standalone()

		iot_cancelJobCmd.Flags().String("comment", "", "An optional comment string describing why the job was canceled.")
		iot_cancelJobCmd.Flags().String("force", "", "(Optional) If `true` job executions with status \"IN\\_PROGRESS\" and \"QUEUED\" are canceled, otherwise only job executions with status \"QUEUED\" are canceled.")
		iot_cancelJobCmd.Flags().String("job-id", "", "The unique identifier you assigned to this job when it was created.")
		iot_cancelJobCmd.Flags().String("reason-code", "", "(Optional)A reason code string that explains why the job was canceled.")
		iot_cancelJobCmd.MarkFlagRequired("job-id")
	})
	iotCmd.AddCommand(iot_cancelJobCmd)
}

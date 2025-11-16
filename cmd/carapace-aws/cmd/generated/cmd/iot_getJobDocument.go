package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getJobDocumentCmd = &cobra.Command{
	Use:   "get-job-document",
	Short: "Gets a job document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getJobDocumentCmd).Standalone()

	iot_getJobDocumentCmd.Flags().String("before-substitution", "", "Provides a view of the job document before and after the substitution parameters have been resolved with their exact values.")
	iot_getJobDocumentCmd.Flags().String("job-id", "", "The unique identifier you assigned to this job when it was created.")
	iot_getJobDocumentCmd.MarkFlagRequired("job-id")
	iotCmd.AddCommand(iot_getJobDocumentCmd)
}

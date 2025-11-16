package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteJobTemplateCmd = &cobra.Command{
	Use:   "delete-job-template",
	Short: "Deletes the specified job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteJobTemplateCmd).Standalone()

	iot_deleteJobTemplateCmd.Flags().String("job-template-id", "", "The unique identifier of the job template to delete.")
	iot_deleteJobTemplateCmd.MarkFlagRequired("job-template-id")
	iotCmd.AddCommand(iot_deleteJobTemplateCmd)
}

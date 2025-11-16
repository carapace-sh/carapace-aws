package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeJobTemplateCmd = &cobra.Command{
	Use:   "describe-job-template",
	Short: "Returns information about a job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeJobTemplateCmd).Standalone()

	iot_describeJobTemplateCmd.Flags().String("job-template-id", "", "The unique identifier of the job template.")
	iot_describeJobTemplateCmd.MarkFlagRequired("job-template-id")
	iotCmd.AddCommand(iot_describeJobTemplateCmd)
}

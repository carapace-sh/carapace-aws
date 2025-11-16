package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_createJobTemplateCmd = &cobra.Command{
	Use:   "create-job-template",
	Short: "Creates a job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_createJobTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_createJobTemplateCmd).Standalone()

		emrContainers_createJobTemplateCmd.Flags().String("client-token", "", "The client token of the job template.")
		emrContainers_createJobTemplateCmd.Flags().String("job-template-data", "", "The job template data which holds values of StartJobRun API request.")
		emrContainers_createJobTemplateCmd.Flags().String("kms-key-arn", "", "The KMS key ARN used to encrypt the job template.")
		emrContainers_createJobTemplateCmd.Flags().String("name", "", "The specified name of the job template.")
		emrContainers_createJobTemplateCmd.Flags().String("tags", "", "The tags that are associated with the job template.")
		emrContainers_createJobTemplateCmd.MarkFlagRequired("client-token")
		emrContainers_createJobTemplateCmd.MarkFlagRequired("job-template-data")
		emrContainers_createJobTemplateCmd.MarkFlagRequired("name")
	})
	emrContainersCmd.AddCommand(emrContainers_createJobTemplateCmd)
}

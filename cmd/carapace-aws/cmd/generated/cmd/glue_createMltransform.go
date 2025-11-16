package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createMltransformCmd = &cobra.Command{
	Use:   "create-mltransform",
	Short: "Creates an Glue machine learning transform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createMltransformCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createMltransformCmd).Standalone()

		glue_createMltransformCmd.Flags().String("description", "", "A description of the machine learning transform that is being defined.")
		glue_createMltransformCmd.Flags().String("glue-version", "", "This value determines which version of Glue this machine learning transform is compatible with.")
		glue_createMltransformCmd.Flags().String("input-record-tables", "", "A list of Glue table definitions used by the transform.")
		glue_createMltransformCmd.Flags().String("max-capacity", "", "The number of Glue data processing units (DPUs) that are allocated to task runs for this transform.")
		glue_createMltransformCmd.Flags().String("max-retries", "", "The maximum number of times to retry a task for this transform after a task run fails.")
		glue_createMltransformCmd.Flags().String("name", "", "The unique name that you give the transform when you create it.")
		glue_createMltransformCmd.Flags().String("number-of-workers", "", "The number of workers of a defined `workerType` that are allocated when this task runs.")
		glue_createMltransformCmd.Flags().String("parameters", "", "The algorithmic parameters that are specific to the transform type used.")
		glue_createMltransformCmd.Flags().String("role", "", "The name or Amazon Resource Name (ARN) of the IAM role with the required permissions.")
		glue_createMltransformCmd.Flags().String("tags", "", "The tags to use with this machine learning transform.")
		glue_createMltransformCmd.Flags().String("timeout", "", "The timeout of the task run for this transform in minutes.")
		glue_createMltransformCmd.Flags().String("transform-encryption", "", "The encryption-at-rest settings of the transform that apply to accessing user data.")
		glue_createMltransformCmd.Flags().String("worker-type", "", "The type of predefined worker that is allocated when this task runs.")
		glue_createMltransformCmd.MarkFlagRequired("input-record-tables")
		glue_createMltransformCmd.MarkFlagRequired("name")
		glue_createMltransformCmd.MarkFlagRequired("parameters")
		glue_createMltransformCmd.MarkFlagRequired("role")
	})
	glueCmd.AddCommand(glue_createMltransformCmd)
}

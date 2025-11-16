package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateMltransformCmd = &cobra.Command{
	Use:   "update-mltransform",
	Short: "Updates an existing machine learning transform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateMltransformCmd).Standalone()

	glue_updateMltransformCmd.Flags().String("description", "", "A description of the transform.")
	glue_updateMltransformCmd.Flags().String("glue-version", "", "This value determines which version of Glue this machine learning transform is compatible with.")
	glue_updateMltransformCmd.Flags().String("max-capacity", "", "The number of Glue data processing units (DPUs) that are allocated to task runs for this transform.")
	glue_updateMltransformCmd.Flags().String("max-retries", "", "The maximum number of times to retry a task for this transform after a task run fails.")
	glue_updateMltransformCmd.Flags().String("name", "", "The unique name that you gave the transform when you created it.")
	glue_updateMltransformCmd.Flags().String("number-of-workers", "", "The number of workers of a defined `workerType` that are allocated when this task runs.")
	glue_updateMltransformCmd.Flags().String("parameters", "", "The configuration parameters that are specific to the transform type (algorithm) used.")
	glue_updateMltransformCmd.Flags().String("role", "", "The name or Amazon Resource Name (ARN) of the IAM role with the required permissions.")
	glue_updateMltransformCmd.Flags().String("timeout", "", "The timeout for a task run for this transform in minutes.")
	glue_updateMltransformCmd.Flags().String("transform-id", "", "A unique identifier that was generated when the transform was created.")
	glue_updateMltransformCmd.Flags().String("worker-type", "", "The type of predefined worker that is allocated when this task runs.")
	glue_updateMltransformCmd.MarkFlagRequired("transform-id")
	glueCmd.AddCommand(glue_updateMltransformCmd)
}

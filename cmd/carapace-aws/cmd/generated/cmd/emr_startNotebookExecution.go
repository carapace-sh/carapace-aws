package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_startNotebookExecutionCmd = &cobra.Command{
	Use:   "start-notebook-execution",
	Short: "Starts a notebook execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_startNotebookExecutionCmd).Standalone()

	emr_startNotebookExecutionCmd.Flags().String("editor-id", "", "The unique identifier of the Amazon EMR Notebook to use for notebook execution.")
	emr_startNotebookExecutionCmd.Flags().String("environment-variables", "", "The environment variables associated with the notebook execution.")
	emr_startNotebookExecutionCmd.Flags().String("execution-engine", "", "Specifies the execution engine (cluster) that runs the notebook execution.")
	emr_startNotebookExecutionCmd.Flags().String("notebook-execution-name", "", "An optional name for the notebook execution.")
	emr_startNotebookExecutionCmd.Flags().String("notebook-instance-security-group-id", "", "The unique identifier of the Amazon EC2 security group to associate with the Amazon EMR Notebook for this notebook execution.")
	emr_startNotebookExecutionCmd.Flags().String("notebook-params", "", "Input parameters in JSON format passed to the Amazon EMR Notebook at runtime for execution.")
	emr_startNotebookExecutionCmd.Flags().String("notebook-s3-location", "", "The Amazon S3 location for the notebook execution input.")
	emr_startNotebookExecutionCmd.Flags().String("output-notebook-format", "", "The output format for the notebook execution.")
	emr_startNotebookExecutionCmd.Flags().String("output-notebook-s3-location", "", "The Amazon S3 location for the notebook execution output.")
	emr_startNotebookExecutionCmd.Flags().String("relative-path", "", "The path and file name of the notebook file for this execution, relative to the path specified for the Amazon EMR Notebook.")
	emr_startNotebookExecutionCmd.Flags().String("service-role", "", "The name or ARN of the IAM role that is used as the service role for Amazon EMR (the Amazon EMR role) for the notebook execution.")
	emr_startNotebookExecutionCmd.Flags().String("tags", "", "A list of tags associated with a notebook execution.")
	emr_startNotebookExecutionCmd.MarkFlagRequired("execution-engine")
	emr_startNotebookExecutionCmd.MarkFlagRequired("service-role")
	emrCmd.AddCommand(emr_startNotebookExecutionCmd)
}

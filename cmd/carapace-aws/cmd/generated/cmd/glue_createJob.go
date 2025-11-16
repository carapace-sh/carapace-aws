package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Creates a new job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createJobCmd).Standalone()

	glue_createJobCmd.Flags().String("allocated-capacity", "", "This parameter is deprecated.")
	glue_createJobCmd.Flags().String("code-gen-configuration-nodes", "", "The representation of a directed acyclic graph on which both the Glue Studio visual component and Glue Studio code generation is based.")
	glue_createJobCmd.Flags().String("command", "", "The `JobCommand` that runs this job.")
	glue_createJobCmd.Flags().String("connections", "", "The connections used for this job.")
	glue_createJobCmd.Flags().String("default-arguments", "", "The default arguments for every run of this job, specified as name-value pairs.")
	glue_createJobCmd.Flags().String("description", "", "Description of the job being defined.")
	glue_createJobCmd.Flags().String("execution-class", "", "Indicates whether the job is run with a standard or flexible execution class.")
	glue_createJobCmd.Flags().String("execution-property", "", "An `ExecutionProperty` specifying the maximum number of concurrent runs allowed for this job.")
	glue_createJobCmd.Flags().String("glue-version", "", "In Spark jobs, `GlueVersion` determines the versions of Apache Spark and Python that Glue available in a job.")
	glue_createJobCmd.Flags().String("job-mode", "", "A mode that describes how a job was created.")
	glue_createJobCmd.Flags().String("job-run-queuing-enabled", "", "Specifies whether job run queuing is enabled for the job runs for this job.")
	glue_createJobCmd.Flags().String("log-uri", "", "This field is reserved for future use.")
	glue_createJobCmd.Flags().String("maintenance-window", "", "This field specifies a day of the week and hour for a maintenance window for streaming jobs.")
	glue_createJobCmd.Flags().String("max-capacity", "", "For Glue version 1.0 or earlier jobs, using the standard worker type, the number of Glue data processing units (DPUs) that can be allocated when this job runs.")
	glue_createJobCmd.Flags().String("max-retries", "", "The maximum number of times to retry this job if it fails.")
	glue_createJobCmd.Flags().String("name", "", "The name you assign to this job definition.")
	glue_createJobCmd.Flags().String("non-overridable-arguments", "", "Arguments for this job that are not overridden when providing job arguments in a job run, specified as name-value pairs.")
	glue_createJobCmd.Flags().String("notification-property", "", "Specifies configuration properties of a job notification.")
	glue_createJobCmd.Flags().String("number-of-workers", "", "The number of workers of a defined `workerType` that are allocated when a job runs.")
	glue_createJobCmd.Flags().String("role", "", "The name or Amazon Resource Name (ARN) of the IAM role associated with this job.")
	glue_createJobCmd.Flags().String("security-configuration", "", "The name of the `SecurityConfiguration` structure to be used with this job.")
	glue_createJobCmd.Flags().String("source-control-details", "", "The details for a source control configuration for a job, allowing synchronization of job artifacts to or from a remote repository.")
	glue_createJobCmd.Flags().String("tags", "", "The tags to use with this job.")
	glue_createJobCmd.Flags().String("timeout", "", "The job timeout in minutes.")
	glue_createJobCmd.Flags().String("worker-type", "", "The type of predefined worker that is allocated when a job runs.")
	glue_createJobCmd.MarkFlagRequired("command")
	glue_createJobCmd.MarkFlagRequired("name")
	glue_createJobCmd.MarkFlagRequired("role")
	glueCmd.AddCommand(glue_createJobCmd)
}

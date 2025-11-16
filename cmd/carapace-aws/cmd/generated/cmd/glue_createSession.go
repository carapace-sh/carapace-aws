package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createSessionCmd = &cobra.Command{
	Use:   "create-session",
	Short: "Creates a new session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createSessionCmd).Standalone()

	glue_createSessionCmd.Flags().String("command", "", "The `SessionCommand` that runs the job.")
	glue_createSessionCmd.Flags().String("connections", "", "The number of connections to use for the session.")
	glue_createSessionCmd.Flags().String("default-arguments", "", "A map array of key-value pairs.")
	glue_createSessionCmd.Flags().String("description", "", "The description of the session.")
	glue_createSessionCmd.Flags().String("glue-version", "", "The Glue version determines the versions of Apache Spark and Python that Glue supports.")
	glue_createSessionCmd.Flags().String("id", "", "The ID of the session request.")
	glue_createSessionCmd.Flags().String("idle-timeout", "", "The number of minutes when idle before session times out.")
	glue_createSessionCmd.Flags().String("max-capacity", "", "The number of Glue data processing units (DPUs) that can be allocated when the job runs.")
	glue_createSessionCmd.Flags().String("number-of-workers", "", "The number of workers of a defined `WorkerType` to use for the session.")
	glue_createSessionCmd.Flags().String("request-origin", "", "The origin of the request.")
	glue_createSessionCmd.Flags().String("role", "", "The IAM Role ARN")
	glue_createSessionCmd.Flags().String("security-configuration", "", "The name of the SecurityConfiguration structure to be used with the session")
	glue_createSessionCmd.Flags().String("tags", "", "The map of key value pairs (tags) belonging to the session.")
	glue_createSessionCmd.Flags().String("timeout", "", "The number of minutes before session times out.")
	glue_createSessionCmd.Flags().String("worker-type", "", "The type of predefined worker that is allocated when a job runs.")
	glue_createSessionCmd.MarkFlagRequired("command")
	glue_createSessionCmd.MarkFlagRequired("id")
	glue_createSessionCmd.MarkFlagRequired("role")
	glueCmd.AddCommand(glue_createSessionCmd)
}

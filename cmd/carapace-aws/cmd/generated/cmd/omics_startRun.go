package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_startRunCmd = &cobra.Command{
	Use:   "start-run",
	Short: "Starts a new run and returns details about the run, or duplicates an existing run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_startRunCmd).Standalone()

	omics_startRunCmd.Flags().String("cache-behavior", "", "The cache behavior for the run.")
	omics_startRunCmd.Flags().String("cache-id", "", "Identifier of the cache associated with this run.")
	omics_startRunCmd.Flags().String("log-level", "", "A log level for the run.")
	omics_startRunCmd.Flags().String("name", "", "A name for the run.")
	omics_startRunCmd.Flags().String("output-uri", "", "An output S3 URI for the run.")
	omics_startRunCmd.Flags().String("parameters", "", "Parameters for the run.")
	omics_startRunCmd.Flags().String("priority", "", "Use the run priority (highest: 1) to establish the order of runs in a run group when you start a run.")
	omics_startRunCmd.Flags().String("request-id", "", "An idempotency token used to dedupe retry requests so that duplicate runs are not created.")
	omics_startRunCmd.Flags().String("retention-mode", "", "The retention mode for the run.")
	omics_startRunCmd.Flags().String("role-arn", "", "A service role for the run.")
	omics_startRunCmd.Flags().String("run-group-id", "", "The run's group ID.")
	omics_startRunCmd.Flags().String("run-id", "", "The ID of a run to duplicate.")
	omics_startRunCmd.Flags().String("storage-capacity", "", "The `STATIC` storage capacity (in gibibytes, GiB) for this run.")
	omics_startRunCmd.Flags().String("storage-type", "", "The storage type for the run.")
	omics_startRunCmd.Flags().String("tags", "", "Tags for the run.")
	omics_startRunCmd.Flags().String("workflow-id", "", "The run's workflow ID.")
	omics_startRunCmd.Flags().String("workflow-owner-id", "", "The 12-digit account ID of the workflow owner that is used for running a shared workflow.")
	omics_startRunCmd.Flags().String("workflow-type", "", "The run's workflow type.")
	omics_startRunCmd.Flags().String("workflow-version-name", "", "The name of the workflow version.")
	omics_startRunCmd.MarkFlagRequired("output-uri")
	omics_startRunCmd.MarkFlagRequired("request-id")
	omics_startRunCmd.MarkFlagRequired("role-arn")
	omicsCmd.AddCommand(omics_startRunCmd)
}

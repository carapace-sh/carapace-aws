package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_startLoaderJobCmd = &cobra.Command{
	Use:   "start-loader-job",
	Short: "Starts a Neptune bulk loader job to load data from an Amazon S3 bucket into a Neptune DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_startLoaderJobCmd).Standalone()

	neptunedata_startLoaderJobCmd.Flags().String("dependencies", "", "This is an optional parameter that can make a queued load request contingent on the successful completion of one or more previous jobs in the queue.")
	neptunedata_startLoaderJobCmd.Flags().Bool("fail-on-error", false, "**`failOnError`** \u00a0 – \u00a0 A flag to toggle a complete stop on an error.")
	neptunedata_startLoaderJobCmd.Flags().String("format", "", "The format of the data.")
	neptunedata_startLoaderJobCmd.Flags().String("iam-role-arn", "", "The Amazon Resource Name (ARN) for an IAM role to be assumed by the Neptune DB instance for access to the S3 bucket.")
	neptunedata_startLoaderJobCmd.Flags().String("mode", "", "The load job mode.")
	neptunedata_startLoaderJobCmd.Flags().Bool("no-fail-on-error", false, "**`failOnError`** \u00a0 – \u00a0 A flag to toggle a complete stop on an error.")
	neptunedata_startLoaderJobCmd.Flags().Bool("no-queue-request", false, "This is an optional flag parameter that indicates whether the load request can be queued up or not.")
	neptunedata_startLoaderJobCmd.Flags().Bool("no-update-single-cardinality-properties", false, "`updateSingleCardinalityProperties` is an optional parameter that controls how the bulk loader treats a new value for single-cardinality vertex or edge properties.")
	neptunedata_startLoaderJobCmd.Flags().Bool("no-user-provided-edge-ids", false, "This parameter is required only when loading openCypher data that contains relationship IDs.")
	neptunedata_startLoaderJobCmd.Flags().String("parallelism", "", "The optional `parallelism` parameter can be set to reduce the number of threads used by the bulk load process.")
	neptunedata_startLoaderJobCmd.Flags().String("parser-configuration", "", "**`parserConfiguration`** \u00a0 – \u00a0 An optional object with additional parser configuration values.")
	neptunedata_startLoaderJobCmd.Flags().Bool("queue-request", false, "This is an optional flag parameter that indicates whether the load request can be queued up or not.")
	neptunedata_startLoaderJobCmd.Flags().String("s3-bucket-region", "", "The Amazon region of the S3 bucket.")
	neptunedata_startLoaderJobCmd.Flags().String("source", "", "The `source` parameter accepts an S3 URI that identifies a single file, multiple files, a folder, or multiple folders.")
	neptunedata_startLoaderJobCmd.Flags().Bool("update-single-cardinality-properties", false, "`updateSingleCardinalityProperties` is an optional parameter that controls how the bulk loader treats a new value for single-cardinality vertex or edge properties.")
	neptunedata_startLoaderJobCmd.Flags().Bool("user-provided-edge-ids", false, "This parameter is required only when loading openCypher data that contains relationship IDs.")
	neptunedata_startLoaderJobCmd.MarkFlagRequired("format")
	neptunedata_startLoaderJobCmd.MarkFlagRequired("iam-role-arn")
	neptunedata_startLoaderJobCmd.Flag("no-fail-on-error").Hidden = true
	neptunedata_startLoaderJobCmd.Flag("no-queue-request").Hidden = true
	neptunedata_startLoaderJobCmd.Flag("no-update-single-cardinality-properties").Hidden = true
	neptunedata_startLoaderJobCmd.Flag("no-user-provided-edge-ids").Hidden = true
	neptunedata_startLoaderJobCmd.MarkFlagRequired("s3-bucket-region")
	neptunedata_startLoaderJobCmd.MarkFlagRequired("source")
	neptunedataCmd.AddCommand(neptunedata_startLoaderJobCmd)
}

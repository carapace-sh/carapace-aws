package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createRunCacheCmd = &cobra.Command{
	Use:   "create-run-cache",
	Short: "Creates a run cache to store and reference task outputs from completed private runs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createRunCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createRunCacheCmd).Standalone()

		omics_createRunCacheCmd.Flags().String("cache-behavior", "", "Default cache behavior for runs that use this cache.")
		omics_createRunCacheCmd.Flags().String("cache-bucket-owner-id", "", "The Amazon Web Services account ID of the expected owner of the S3 bucket for the run cache.")
		omics_createRunCacheCmd.Flags().String("cache-s3-location", "", "Specify the S3 location for storing the cached task outputs.")
		omics_createRunCacheCmd.Flags().String("description", "", "Enter a description of the run cache.")
		omics_createRunCacheCmd.Flags().String("name", "", "Enter a user-friendly name for the run cache.")
		omics_createRunCacheCmd.Flags().String("request-id", "", "A unique request token, to ensure idempotency.")
		omics_createRunCacheCmd.Flags().String("tags", "", "Specify one or more tags to associate with this run cache.")
		omics_createRunCacheCmd.MarkFlagRequired("cache-s3-location")
		omics_createRunCacheCmd.MarkFlagRequired("request-id")
	})
	omicsCmd.AddCommand(omics_createRunCacheCmd)
}

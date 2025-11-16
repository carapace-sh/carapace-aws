package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeJobLogItemsCmd = &cobra.Command{
	Use:   "describe-job-log-items",
	Short: "Retrieves a detailed Job log with pagination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeJobLogItemsCmd).Standalone()

	drs_describeJobLogItemsCmd.Flags().String("job-id", "", "The ID of the Job for which Job log items will be retrieved.")
	drs_describeJobLogItemsCmd.Flags().String("max-results", "", "Maximum number of Job log items to retrieve.")
	drs_describeJobLogItemsCmd.Flags().String("next-token", "", "The token of the next Job log items to retrieve.")
	drs_describeJobLogItemsCmd.MarkFlagRequired("job-id")
	drsCmd.AddCommand(drs_describeJobLogItemsCmd)
}

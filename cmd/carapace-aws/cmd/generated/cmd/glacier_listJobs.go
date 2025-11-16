package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "This operation lists jobs for a vault, including jobs that are in-progress and jobs that have recently finished.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_listJobsCmd).Standalone()

	glacier_listJobsCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_listJobsCmd.Flags().String("completed", "", "The state of the jobs to return.")
	glacier_listJobsCmd.Flags().String("limit", "", "The maximum number of jobs to be returned.")
	glacier_listJobsCmd.Flags().String("marker", "", "An opaque string used for pagination.")
	glacier_listJobsCmd.Flags().String("statuscode", "", "The type of job status to return.")
	glacier_listJobsCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_listJobsCmd.MarkFlagRequired("account-id")
	glacier_listJobsCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_listJobsCmd)
}

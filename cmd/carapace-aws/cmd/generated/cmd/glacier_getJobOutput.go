package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_getJobOutputCmd = &cobra.Command{
	Use:   "get-job-output",
	Short: "This operation downloads the output of the job you initiated using [InitiateJob]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_getJobOutputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_getJobOutputCmd).Standalone()

		glacier_getJobOutputCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_getJobOutputCmd.Flags().String("job-id", "", "The job ID whose data is downloaded.")
		glacier_getJobOutputCmd.Flags().String("range", "", "The range of bytes to retrieve from the output.")
		glacier_getJobOutputCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_getJobOutputCmd.MarkFlagRequired("account-id")
		glacier_getJobOutputCmd.MarkFlagRequired("job-id")
		glacier_getJobOutputCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_getJobOutputCmd)
}

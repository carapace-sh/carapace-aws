package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_describeJobCmd = &cobra.Command{
	Use:   "describe-job",
	Short: "This operation returns information about a job you previously initiated, including the job initiation date, the user who initiated the job, the job status code/message and the Amazon SNS topic to notify after Amazon S3 Glacier (Glacier) completes the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_describeJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_describeJobCmd).Standalone()

		glacier_describeJobCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_describeJobCmd.Flags().String("job-id", "", "The ID of the job to describe.")
		glacier_describeJobCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_describeJobCmd.MarkFlagRequired("account-id")
		glacier_describeJobCmd.MarkFlagRequired("job-id")
		glacier_describeJobCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_describeJobCmd)
}

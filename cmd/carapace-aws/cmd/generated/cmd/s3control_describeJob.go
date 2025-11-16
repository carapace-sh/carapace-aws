package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_describeJobCmd = &cobra.Command{
	Use:   "describe-job",
	Short: "Retrieves the configuration parameters and status for a Batch Operations job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_describeJobCmd).Standalone()

	s3control_describeJobCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the S3 Batch Operations job.")
	s3control_describeJobCmd.Flags().String("job-id", "", "The ID for the job whose information you want to retrieve.")
	s3control_describeJobCmd.MarkFlagRequired("account-id")
	s3control_describeJobCmd.MarkFlagRequired("job-id")
	s3controlCmd.AddCommand(s3control_describeJobCmd)
}

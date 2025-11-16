package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_startActivityStreamCmd = &cobra.Command{
	Use:   "start-activity-stream",
	Short: "Starts a database activity stream to monitor activity on the database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_startActivityStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_startActivityStreamCmd).Standalone()

		rds_startActivityStreamCmd.Flags().String("apply-immediately", "", "Specifies whether or not the database activity stream is to start as soon as possible, regardless of the maintenance window for the database.")
		rds_startActivityStreamCmd.Flags().String("engine-native-audit-fields-included", "", "Specifies whether the database activity stream includes engine-native audit fields.")
		rds_startActivityStreamCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for encrypting messages in the database activity stream.")
		rds_startActivityStreamCmd.Flags().String("mode", "", "Specifies the mode of the database activity stream.")
		rds_startActivityStreamCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the DB cluster, for example, `arn:aws:rds:us-east-1:12345667890:cluster:das-cluster`.")
		rds_startActivityStreamCmd.MarkFlagRequired("kms-key-id")
		rds_startActivityStreamCmd.MarkFlagRequired("mode")
		rds_startActivityStreamCmd.MarkFlagRequired("resource-arn")
	})
	rdsCmd.AddCommand(rds_startActivityStreamCmd)
}

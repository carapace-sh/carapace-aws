package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyActivityStreamCmd = &cobra.Command{
	Use:   "modify-activity-stream",
	Short: "Changes the audit policy state of a database activity stream to either locked (default) or unlocked.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyActivityStreamCmd).Standalone()

	rds_modifyActivityStreamCmd.Flags().String("audit-policy-state", "", "The audit policy state.")
	rds_modifyActivityStreamCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the RDS for Oracle or Microsoft SQL Server DB instance.")
	rdsCmd.AddCommand(rds_modifyActivityStreamCmd)
}

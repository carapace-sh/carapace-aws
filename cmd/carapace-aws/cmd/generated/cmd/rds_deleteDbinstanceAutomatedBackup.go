package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbinstanceAutomatedBackupCmd = &cobra.Command{
	Use:   "delete-dbinstance-automated-backup",
	Short: "Deletes automated backups using the `DbiResourceId` value of the source DB instance or the Amazon Resource Name (ARN) of the automated backups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbinstanceAutomatedBackupCmd).Standalone()

	rds_deleteDbinstanceAutomatedBackupCmd.Flags().String("dbi-resource-id", "", "The identifier for the source DB instance, which can't be changed and which is unique to an Amazon Web Services Region.")
	rds_deleteDbinstanceAutomatedBackupCmd.Flags().String("dbinstance-automated-backups-arn", "", "The Amazon Resource Name (ARN) of the automated backups to delete, for example, `arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE`.")
	rdsCmd.AddCommand(rds_deleteDbinstanceAutomatedBackupCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_restoreTableFromRecoveryPointCmd = &cobra.Command{
	Use:   "restore-table-from-recovery-point",
	Short: "Restores a table from a recovery point to your Amazon Redshift Serverless instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_restoreTableFromRecoveryPointCmd).Standalone()

	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().Bool("activate-case-sensitive-identifier", false, "Indicates whether name identifiers for database, schema, and table are case sensitive.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("namespace-name", "", "Namespace of the recovery point to restore from.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("new-table-name", "", "The name of the table to create from the restore operation.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().Bool("no-activate-case-sensitive-identifier", false, "Indicates whether name identifiers for database, schema, and table are case sensitive.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("recovery-point-id", "", "The ID of the recovery point to restore the table from.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("source-database-name", "", "The name of the source database that contains the table being restored.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("source-schema-name", "", "The name of the source schema that contains the table being restored.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("source-table-name", "", "The name of the source table being restored.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("target-database-name", "", "The name of the database to restore the table to.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("target-schema-name", "", "The name of the schema to restore the table to.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flags().String("workgroup-name", "", "The workgroup to restore the table to.")
	redshiftServerless_restoreTableFromRecoveryPointCmd.MarkFlagRequired("namespace-name")
	redshiftServerless_restoreTableFromRecoveryPointCmd.MarkFlagRequired("new-table-name")
	redshiftServerless_restoreTableFromRecoveryPointCmd.Flag("no-activate-case-sensitive-identifier").Hidden = true
	redshiftServerless_restoreTableFromRecoveryPointCmd.MarkFlagRequired("recovery-point-id")
	redshiftServerless_restoreTableFromRecoveryPointCmd.MarkFlagRequired("source-database-name")
	redshiftServerless_restoreTableFromRecoveryPointCmd.MarkFlagRequired("source-table-name")
	redshiftServerless_restoreTableFromRecoveryPointCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_restoreTableFromRecoveryPointCmd)
}

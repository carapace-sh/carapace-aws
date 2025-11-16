package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_createDbinstanceCmd = &cobra.Command{
	Use:   "create-dbinstance",
	Short: "Creates a new instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_createDbinstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_createDbinstanceCmd).Standalone()

		docdb_createDbinstanceCmd.Flags().String("auto-minor-version-upgrade", "", "This parameter does not apply to Amazon DocumentDB.")
		docdb_createDbinstanceCmd.Flags().String("availability-zone", "", "The Amazon EC2 Availability Zone that the instance is created in.")
		docdb_createDbinstanceCmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the DB instance's server certificate.")
		docdb_createDbinstanceCmd.Flags().String("copy-tags-to-snapshot", "", "A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance.")
		docdb_createDbinstanceCmd.Flags().String("dbcluster-identifier", "", "The identifier of the cluster that the instance will belong to.")
		docdb_createDbinstanceCmd.Flags().String("dbinstance-class", "", "The compute and memory capacity of the instance; for example, `db.r5.large`.")
		docdb_createDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The instance identifier.")
		docdb_createDbinstanceCmd.Flags().String("enable-performance-insights", "", "A value that indicates whether to enable Performance Insights for the DB Instance.")
		docdb_createDbinstanceCmd.Flags().String("engine", "", "The name of the database engine to be used for this instance.")
		docdb_createDbinstanceCmd.Flags().String("performance-insights-kmskey-id", "", "The KMS key identifier for encryption of Performance Insights data.")
		docdb_createDbinstanceCmd.Flags().String("preferred-maintenance-window", "", "The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		docdb_createDbinstanceCmd.Flags().String("promotion-tier", "", "A value that specifies the order in which an Amazon DocumentDB replica is promoted to the primary instance after a failure of the existing primary instance.")
		docdb_createDbinstanceCmd.Flags().String("tags", "", "The tags to be assigned to the instance.")
		docdb_createDbinstanceCmd.MarkFlagRequired("dbcluster-identifier")
		docdb_createDbinstanceCmd.MarkFlagRequired("dbinstance-class")
		docdb_createDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
		docdb_createDbinstanceCmd.MarkFlagRequired("engine")
	})
	docdbCmd.AddCommand(docdb_createDbinstanceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_modifyDbinstanceCmd = &cobra.Command{
	Use:   "modify-dbinstance",
	Short: "Modifies settings for an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_modifyDbinstanceCmd).Standalone()

	docdb_modifyDbinstanceCmd.Flags().Bool("apply-immediately", false, "Specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the instance.")
	docdb_modifyDbinstanceCmd.Flags().String("auto-minor-version-upgrade", "", "This parameter does not apply to Amazon DocumentDB.")
	docdb_modifyDbinstanceCmd.Flags().String("cacertificate-identifier", "", "Indicates the certificate that needs to be associated with the instance.")
	docdb_modifyDbinstanceCmd.Flags().String("certificate-rotation-restart", "", "Specifies whether the DB instance is restarted when you rotate your SSL/TLS certificate.")
	docdb_modifyDbinstanceCmd.Flags().String("copy-tags-to-snapshot", "", "A value that indicates whether to copy all tags from the DB instance to snapshots of the DB instance.")
	docdb_modifyDbinstanceCmd.Flags().String("dbinstance-class", "", "The new compute and memory capacity of the instance; for example, `db.r5.large`. Not all instance classes are available in all Amazon Web Services Regions.")
	docdb_modifyDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The instance identifier.")
	docdb_modifyDbinstanceCmd.Flags().String("enable-performance-insights", "", "A value that indicates whether to enable Performance Insights for the DB Instance.")
	docdb_modifyDbinstanceCmd.Flags().String("new-dbinstance-identifier", "", "The new instance identifier for the instance when renaming an instance.")
	docdb_modifyDbinstanceCmd.Flags().Bool("no-apply-immediately", false, "Specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the instance.")
	docdb_modifyDbinstanceCmd.Flags().String("performance-insights-kmskey-id", "", "The KMS key identifier for encryption of Performance Insights data.")
	docdb_modifyDbinstanceCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range (in UTC) during which system maintenance can occur, which might result in an outage.")
	docdb_modifyDbinstanceCmd.Flags().String("promotion-tier", "", "A value that specifies the order in which an Amazon DocumentDB replica is promoted to the primary instance after a failure of the existing primary instance.")
	docdb_modifyDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	docdb_modifyDbinstanceCmd.Flag("no-apply-immediately").Hidden = true
	docdbCmd.AddCommand(docdb_modifyDbinstanceCmd)
}

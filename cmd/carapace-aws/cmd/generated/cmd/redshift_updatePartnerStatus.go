package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_updatePartnerStatusCmd = &cobra.Command{
	Use:   "update-partner-status",
	Short: "Updates the status of a partner integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_updatePartnerStatusCmd).Standalone()

	redshift_updatePartnerStatusCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the cluster.")
	redshift_updatePartnerStatusCmd.Flags().String("cluster-identifier", "", "The cluster identifier of the cluster whose partner integration status is being updated.")
	redshift_updatePartnerStatusCmd.Flags().String("database-name", "", "The name of the database whose partner integration status is being updated.")
	redshift_updatePartnerStatusCmd.Flags().String("partner-name", "", "The name of the partner whose integration status is being updated.")
	redshift_updatePartnerStatusCmd.Flags().String("status", "", "The value of the updated status.")
	redshift_updatePartnerStatusCmd.Flags().String("status-message", "", "The status message provided by the partner.")
	redshift_updatePartnerStatusCmd.MarkFlagRequired("account-id")
	redshift_updatePartnerStatusCmd.MarkFlagRequired("cluster-identifier")
	redshift_updatePartnerStatusCmd.MarkFlagRequired("database-name")
	redshift_updatePartnerStatusCmd.MarkFlagRequired("partner-name")
	redshift_updatePartnerStatusCmd.MarkFlagRequired("status")
	redshiftCmd.AddCommand(redshift_updatePartnerStatusCmd)
}

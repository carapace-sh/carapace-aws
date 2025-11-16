package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deletePartnerCmd = &cobra.Command{
	Use:   "delete-partner",
	Short: "Deletes a partner integration from a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deletePartnerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deletePartnerCmd).Standalone()

		redshift_deletePartnerCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the cluster.")
		redshift_deletePartnerCmd.Flags().String("cluster-identifier", "", "The cluster identifier of the cluster that receives data from the partner.")
		redshift_deletePartnerCmd.Flags().String("database-name", "", "The name of the database that receives data from the partner.")
		redshift_deletePartnerCmd.Flags().String("partner-name", "", "The name of the partner that is authorized to send data.")
		redshift_deletePartnerCmd.MarkFlagRequired("account-id")
		redshift_deletePartnerCmd.MarkFlagRequired("cluster-identifier")
		redshift_deletePartnerCmd.MarkFlagRequired("database-name")
		redshift_deletePartnerCmd.MarkFlagRequired("partner-name")
	})
	redshiftCmd.AddCommand(redshift_deletePartnerCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_addPartnerCmd = &cobra.Command{
	Use:   "add-partner",
	Short: "Adds a partner integration to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_addPartnerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_addPartnerCmd).Standalone()

		redshift_addPartnerCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the cluster.")
		redshift_addPartnerCmd.Flags().String("cluster-identifier", "", "The cluster identifier of the cluster that receives data from the partner.")
		redshift_addPartnerCmd.Flags().String("database-name", "", "The name of the database that receives data from the partner.")
		redshift_addPartnerCmd.Flags().String("partner-name", "", "The name of the partner that is authorized to send data.")
		redshift_addPartnerCmd.MarkFlagRequired("account-id")
		redshift_addPartnerCmd.MarkFlagRequired("cluster-identifier")
		redshift_addPartnerCmd.MarkFlagRequired("database-name")
		redshift_addPartnerCmd.MarkFlagRequired("partner-name")
	})
	redshiftCmd.AddCommand(redshift_addPartnerCmd)
}

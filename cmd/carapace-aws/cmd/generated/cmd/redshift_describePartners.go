package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describePartnersCmd = &cobra.Command{
	Use:   "describe-partners",
	Short: "Returns information about the partner integrations defined for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describePartnersCmd).Standalone()

	redshift_describePartnersCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the cluster.")
	redshift_describePartnersCmd.Flags().String("cluster-identifier", "", "The cluster identifier of the cluster whose partner integration is being described.")
	redshift_describePartnersCmd.Flags().String("database-name", "", "The name of the database whose partner integration is being described.")
	redshift_describePartnersCmd.Flags().String("partner-name", "", "The name of the partner that is being described.")
	redshift_describePartnersCmd.MarkFlagRequired("account-id")
	redshift_describePartnersCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_describePartnersCmd)
}

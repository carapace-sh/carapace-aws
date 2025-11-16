package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_failoverPrimaryComputeCmd = &cobra.Command{
	Use:   "failover-primary-compute",
	Short: "Fails over the primary compute unit of the specified Multi-AZ cluster to another Availability Zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_failoverPrimaryComputeCmd).Standalone()

	redshift_failoverPrimaryComputeCmd.Flags().String("cluster-identifier", "", "The unique identifier of the cluster for which the primary compute unit will be failed over to another Availability Zone.")
	redshift_failoverPrimaryComputeCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_failoverPrimaryComputeCmd)
}

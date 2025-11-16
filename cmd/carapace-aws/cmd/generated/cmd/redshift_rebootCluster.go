package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_rebootClusterCmd = &cobra.Command{
	Use:   "reboot-cluster",
	Short: "Reboots a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_rebootClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_rebootClusterCmd).Standalone()

		redshift_rebootClusterCmd.Flags().String("cluster-identifier", "", "The cluster identifier.")
		redshift_rebootClusterCmd.MarkFlagRequired("cluster-identifier")
	})
	redshiftCmd.AddCommand(redshift_rebootClusterCmd)
}

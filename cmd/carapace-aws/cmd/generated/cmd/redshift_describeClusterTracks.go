package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeClusterTracksCmd = &cobra.Command{
	Use:   "describe-cluster-tracks",
	Short: "Returns a list of all the available maintenance tracks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeClusterTracksCmd).Standalone()

	redshift_describeClusterTracksCmd.Flags().String("maintenance-track-name", "", "The name of the maintenance track.")
	redshift_describeClusterTracksCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeClusterTracksCmd.Flags().String("max-records", "", "An integer value for the maximum number of maintenance tracks to return.")
	redshiftCmd.AddCommand(redshift_describeClusterTracksCmd)
}

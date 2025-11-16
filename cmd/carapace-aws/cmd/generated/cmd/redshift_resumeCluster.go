package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_resumeClusterCmd = &cobra.Command{
	Use:   "resume-cluster",
	Short: "Resumes a paused cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_resumeClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_resumeClusterCmd).Standalone()

		redshift_resumeClusterCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster to be resumed.")
		redshift_resumeClusterCmd.MarkFlagRequired("cluster-identifier")
	})
	redshiftCmd.AddCommand(redshift_resumeClusterCmd)
}

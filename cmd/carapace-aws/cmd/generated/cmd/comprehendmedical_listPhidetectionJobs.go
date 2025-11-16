package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_listPhidetectionJobsCmd = &cobra.Command{
	Use:   "list-phidetection-jobs",
	Short: "Gets a list of protected health information (PHI) detection jobs you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_listPhidetectionJobsCmd).Standalone()

	comprehendmedical_listPhidetectionJobsCmd.Flags().String("filter", "", "Filters the jobs that are returned.")
	comprehendmedical_listPhidetectionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in each page.")
	comprehendmedical_listPhidetectionJobsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendmedicalCmd.AddCommand(comprehendmedical_listPhidetectionJobsCmd)
}

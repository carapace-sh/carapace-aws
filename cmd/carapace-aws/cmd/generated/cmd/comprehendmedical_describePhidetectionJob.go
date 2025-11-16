package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_describePhidetectionJobCmd = &cobra.Command{
	Use:   "describe-phidetection-job",
	Short: "Gets the properties associated with a protected health information (PHI) detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_describePhidetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_describePhidetectionJobCmd).Standalone()

		comprehendmedical_describePhidetectionJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend Medical generated for the job.")
		comprehendmedical_describePhidetectionJobCmd.MarkFlagRequired("job-id")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_describePhidetectionJobCmd)
}

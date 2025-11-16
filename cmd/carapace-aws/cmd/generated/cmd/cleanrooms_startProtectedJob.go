package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_startProtectedJobCmd = &cobra.Command{
	Use:   "start-protected-job",
	Short: "Creates a protected job that is started by Clean Rooms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_startProtectedJobCmd).Standalone()

	cleanrooms_startProtectedJobCmd.Flags().String("compute-configuration", "", "The compute configuration for the protected job.")
	cleanrooms_startProtectedJobCmd.Flags().String("job-parameters", "", "The job parameters.")
	cleanrooms_startProtectedJobCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership to run this job against.")
	cleanrooms_startProtectedJobCmd.Flags().String("result-configuration", "", "The details needed to write the job results.")
	cleanrooms_startProtectedJobCmd.Flags().String("type", "", "The type of protected job to start.")
	cleanrooms_startProtectedJobCmd.MarkFlagRequired("job-parameters")
	cleanrooms_startProtectedJobCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_startProtectedJobCmd.MarkFlagRequired("type")
	cleanroomsCmd.AddCommand(cleanrooms_startProtectedJobCmd)
}

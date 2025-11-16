package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_getDashboardForJobRunCmd = &cobra.Command{
	Use:   "get-dashboard-for-job-run",
	Short: "Creates and returns a URL that you can use to access the application UIs for a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_getDashboardForJobRunCmd).Standalone()

	emrServerless_getDashboardForJobRunCmd.Flags().Bool("access-system-profile-logs", false, "Allows access to system profile logs for Lake Formation-enabled jobs.")
	emrServerless_getDashboardForJobRunCmd.Flags().String("application-id", "", "The ID of the application.")
	emrServerless_getDashboardForJobRunCmd.Flags().String("attempt", "", "An optimal parameter that indicates the amount of attempts for the job.")
	emrServerless_getDashboardForJobRunCmd.Flags().String("job-run-id", "", "The ID of the job run.")
	emrServerless_getDashboardForJobRunCmd.Flags().Bool("no-access-system-profile-logs", false, "Allows access to system profile logs for Lake Formation-enabled jobs.")
	emrServerless_getDashboardForJobRunCmd.MarkFlagRequired("application-id")
	emrServerless_getDashboardForJobRunCmd.MarkFlagRequired("job-run-id")
	emrServerless_getDashboardForJobRunCmd.Flag("no-access-system-profile-logs").Hidden = true
	emrServerlessCmd.AddCommand(emrServerless_getDashboardForJobRunCmd)
}

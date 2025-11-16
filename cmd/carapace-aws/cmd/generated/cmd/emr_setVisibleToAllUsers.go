package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_setVisibleToAllUsersCmd = &cobra.Command{
	Use:   "set-visible-to-all-users",
	Short: "The SetVisibleToAllUsers parameter is no longer supported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_setVisibleToAllUsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_setVisibleToAllUsersCmd).Standalone()

		emr_setVisibleToAllUsersCmd.Flags().String("job-flow-ids", "", "The unique identifier of the job flow (cluster).")
		emr_setVisibleToAllUsersCmd.Flags().Bool("no-visible-to-all-users", false, "A value of `true` indicates that an IAM principal in the Amazon Web Services account can perform Amazon EMR actions on the cluster that the IAM policies attached to the principal allow.")
		emr_setVisibleToAllUsersCmd.Flags().Bool("visible-to-all-users", false, "A value of `true` indicates that an IAM principal in the Amazon Web Services account can perform Amazon EMR actions on the cluster that the IAM policies attached to the principal allow.")
		emr_setVisibleToAllUsersCmd.MarkFlagRequired("job-flow-ids")
		emr_setVisibleToAllUsersCmd.Flag("no-visible-to-all-users").Hidden = true
		emr_setVisibleToAllUsersCmd.MarkFlagRequired("no-visible-to-all-users")
		emr_setVisibleToAllUsersCmd.MarkFlagRequired("visible-to-all-users")
	})
	emrCmd.AddCommand(emr_setVisibleToAllUsersCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listProfileNotificationsCmd = &cobra.Command{
	Use:   "list-profile-notifications",
	Short: "List profile notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listProfileNotificationsCmd).Standalone()

	wellarchitected_listProfileNotificationsCmd.Flags().String("max-results", "", "")
	wellarchitected_listProfileNotificationsCmd.Flags().String("next-token", "", "")
	wellarchitected_listProfileNotificationsCmd.Flags().String("workload-id", "", "")
	wellarchitectedCmd.AddCommand(wellarchitected_listProfileNotificationsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteJobCmd = &cobra.Command{
	Use:   "delete-job",
	Short: "Deletes a single Job by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_deleteJobCmd).Standalone()

		mgn_deleteJobCmd.Flags().String("account-id", "", "Request to delete Job from service by Account ID.")
		mgn_deleteJobCmd.Flags().String("job-id", "", "Request to delete Job from service by Job ID.")
		mgn_deleteJobCmd.MarkFlagRequired("job-id")
	})
	mgnCmd.AddCommand(mgn_deleteJobCmd)
}

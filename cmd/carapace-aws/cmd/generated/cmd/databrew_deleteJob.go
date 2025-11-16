package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_deleteJobCmd = &cobra.Command{
	Use:   "delete-job",
	Short: "Deletes the specified DataBrew job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_deleteJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_deleteJobCmd).Standalone()

		databrew_deleteJobCmd.Flags().String("name", "", "The name of the job to be deleted.")
		databrew_deleteJobCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_deleteJobCmd)
}

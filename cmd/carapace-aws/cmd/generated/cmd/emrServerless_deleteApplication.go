package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrServerless_deleteApplicationCmd).Standalone()

		emrServerless_deleteApplicationCmd.Flags().String("application-id", "", "The ID of the application that will be deleted.")
		emrServerless_deleteApplicationCmd.MarkFlagRequired("application-id")
	})
	emrServerlessCmd.AddCommand(emrServerless_deleteApplicationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_deleteApplicationCmd).Standalone()

		qbusiness_deleteApplicationCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application.")
		qbusiness_deleteApplicationCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_deleteApplicationCmd)
}

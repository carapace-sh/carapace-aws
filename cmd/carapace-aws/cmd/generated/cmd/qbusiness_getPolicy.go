package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Retrieves the current permission policy for a Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getPolicyCmd).Standalone()

		qbusiness_getPolicyCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
		qbusiness_getPolicyCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getPolicyCmd)
}

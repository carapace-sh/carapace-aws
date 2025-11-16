package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Gets information about an existing Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getApplicationCmd).Standalone()

	qbusiness_getApplicationCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application.")
	qbusiness_getApplicationCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_getApplicationCmd)
}

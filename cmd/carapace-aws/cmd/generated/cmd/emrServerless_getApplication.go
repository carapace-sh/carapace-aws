package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Displays detailed information about a specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_getApplicationCmd).Standalone()

	emrServerless_getApplicationCmd.Flags().String("application-id", "", "The ID of the application that will be described.")
	emrServerless_getApplicationCmd.MarkFlagRequired("application-id")
	emrServerlessCmd.AddCommand(emrServerless_getApplicationCmd)
}

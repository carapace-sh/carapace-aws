package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listSuitesCmd = &cobra.Command{
	Use:   "list-suites",
	Short: "Gets information about test suites for a given job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listSuitesCmd).Standalone()

	devicefarm_listSuitesCmd.Flags().String("arn", "", "The job's Amazon Resource Name (ARN).")
	devicefarm_listSuitesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarm_listSuitesCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_listSuitesCmd)
}

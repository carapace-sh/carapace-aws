package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listSamplesCmd = &cobra.Command{
	Use:   "list-samples",
	Short: "Gets information about samples, given an AWS Device Farm job ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listSamplesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listSamplesCmd).Standalone()

		devicefarm_listSamplesCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the job used to list samples.")
		devicefarm_listSamplesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
		devicefarm_listSamplesCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listSamplesCmd)
}

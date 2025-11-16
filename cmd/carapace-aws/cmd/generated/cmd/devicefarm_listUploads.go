package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listUploadsCmd = &cobra.Command{
	Use:   "list-uploads",
	Short: "Gets information about uploads, given an AWS Device Farm project ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listUploadsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listUploadsCmd).Standalone()

		devicefarm_listUploadsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the project for which you want to list uploads.")
		devicefarm_listUploadsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
		devicefarm_listUploadsCmd.Flags().String("type", "", "The type of upload.")
		devicefarm_listUploadsCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listUploadsCmd)
}

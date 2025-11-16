package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_configureLogsCmd = &cobra.Command{
	Use:   "configure-logs",
	Short: "Changes the Channel's properities to configure log subscription",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_configureLogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_configureLogsCmd).Standalone()

		mediapackage_configureLogsCmd.Flags().String("egress-access-logs", "", "")
		mediapackage_configureLogsCmd.Flags().String("id", "", "The ID of the channel to log subscription.")
		mediapackage_configureLogsCmd.Flags().String("ingress-access-logs", "", "")
		mediapackage_configureLogsCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_configureLogsCmd)
}

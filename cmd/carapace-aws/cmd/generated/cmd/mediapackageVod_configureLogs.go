package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_configureLogsCmd = &cobra.Command{
	Use:   "configure-logs",
	Short: "Changes the packaging group's properities to configure log subscription",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_configureLogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_configureLogsCmd).Standalone()

		mediapackageVod_configureLogsCmd.Flags().String("egress-access-logs", "", "")
		mediapackageVod_configureLogsCmd.Flags().String("id", "", "The ID of a MediaPackage VOD PackagingGroup resource.")
		mediapackageVod_configureLogsCmd.MarkFlagRequired("id")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_configureLogsCmd)
}

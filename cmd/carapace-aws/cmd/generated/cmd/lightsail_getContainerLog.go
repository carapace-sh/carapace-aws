package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContainerLogCmd = &cobra.Command{
	Use:   "get-container-log",
	Short: "Returns the log events of a container of your Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContainerLogCmd).Standalone()

	lightsail_getContainerLogCmd.Flags().String("container-name", "", "The name of the container that is either running or previously ran on the container service for which to return a log.")
	lightsail_getContainerLogCmd.Flags().String("end-time", "", "The end of the time interval for which to get log data.")
	lightsail_getContainerLogCmd.Flags().String("filter-pattern", "", "The pattern to use to filter the returned log events to a specific term.")
	lightsail_getContainerLogCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsail_getContainerLogCmd.Flags().String("service-name", "", "The name of the container service for which to get a container log.")
	lightsail_getContainerLogCmd.Flags().String("start-time", "", "The start of the time interval for which to get log data.")
	lightsail_getContainerLogCmd.MarkFlagRequired("container-name")
	lightsail_getContainerLogCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_getContainerLogCmd)
}

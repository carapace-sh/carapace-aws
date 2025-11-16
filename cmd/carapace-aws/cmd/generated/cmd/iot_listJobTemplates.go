package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listJobTemplatesCmd = &cobra.Command{
	Use:   "list-job-templates",
	Short: "Returns a list of job templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listJobTemplatesCmd).Standalone()

	iot_listJobTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return in the list.")
	iot_listJobTemplatesCmd.Flags().String("next-token", "", "The token to use to return the next set of results in the list.")
	iotCmd.AddCommand(iot_listJobTemplatesCmd)
}

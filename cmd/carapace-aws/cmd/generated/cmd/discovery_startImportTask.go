package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_startImportTaskCmd = &cobra.Command{
	Use:   "start-import-task",
	Short: "Starts an import task, which allows you to import details of your on-premises environment directly into Amazon Web Services Migration Hub without having to use the Amazon Web Services Application Discovery Service (Application Discovery Service) tools such as the Amazon Web Services Application Discovery Service Agentless Collector or Application Discovery Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_startImportTaskCmd).Standalone()

	discovery_startImportTaskCmd.Flags().String("client-request-token", "", "Optional.")
	discovery_startImportTaskCmd.Flags().String("import-url", "", "The URL for your import file that you've uploaded to Amazon S3.")
	discovery_startImportTaskCmd.Flags().String("name", "", "A descriptive name for this request.")
	discovery_startImportTaskCmd.MarkFlagRequired("import-url")
	discovery_startImportTaskCmd.MarkFlagRequired("name")
	discoveryCmd.AddCommand(discovery_startImportTaskCmd)
}

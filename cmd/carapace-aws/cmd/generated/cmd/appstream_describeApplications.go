package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeApplicationsCmd = &cobra.Command{
	Use:   "describe-applications",
	Short: "Retrieves a list that describes one or more applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeApplicationsCmd).Standalone()

		appstream_describeApplicationsCmd.Flags().String("arns", "", "The ARNs for the applications.")
		appstream_describeApplicationsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
		appstream_describeApplicationsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	appstreamCmd.AddCommand(appstream_describeApplicationsCmd)
}

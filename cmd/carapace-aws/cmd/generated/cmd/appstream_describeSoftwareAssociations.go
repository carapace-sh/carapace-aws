package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeSoftwareAssociationsCmd = &cobra.Command{
	Use:   "describe-software-associations",
	Short: "Retrieves license included application associations for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeSoftwareAssociationsCmd).Standalone()

	appstream_describeSoftwareAssociationsCmd.Flags().String("associated-resource", "", "The ARN of the resource to describe software associations.")
	appstream_describeSoftwareAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	appstream_describeSoftwareAssociationsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	appstream_describeSoftwareAssociationsCmd.MarkFlagRequired("associated-resource")
	appstreamCmd.AddCommand(appstream_describeSoftwareAssociationsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "DEPRECATED - The `DescribeTags` action is deprecated and not maintained.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_describeTagsCmd).Standalone()

		efs_describeTagsCmd.Flags().String("file-system-id", "", "The ID of the file system whose tag set you want to retrieve.")
		efs_describeTagsCmd.Flags().String("marker", "", "(Optional) An opaque pagination token returned from a previous `DescribeTags` operation (String).")
		efs_describeTagsCmd.Flags().String("max-items", "", "(Optional) The maximum number of file system tags to return in the response.")
		efs_describeTagsCmd.MarkFlagRequired("file-system-id")
	})
	efsCmd.AddCommand(efs_describeTagsCmd)
}

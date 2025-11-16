package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_createResourceShareCmd = &cobra.Command{
	Use:   "create-resource-share",
	Short: "Create a new resource share request for MediaConvert resources with AWS Support.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_createResourceShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_createResourceShareCmd).Standalone()

		mediaconvert_createResourceShareCmd.Flags().String("job-id", "", "Specify MediaConvert Job ID or ARN to share")
		mediaconvert_createResourceShareCmd.Flags().String("support-case-id", "", "AWS Support case identifier")
		mediaconvert_createResourceShareCmd.MarkFlagRequired("job-id")
		mediaconvert_createResourceShareCmd.MarkFlagRequired("support-case-id")
	})
	mediaconvertCmd.AddCommand(mediaconvert_createResourceShareCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3outposts_listOutpostsWithS3Cmd = &cobra.Command{
	Use:   "list-outposts-with-s3",
	Short: "Lists the Outposts with S3 on Outposts capacity for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3outposts_listOutpostsWithS3Cmd).Standalone()

	s3outposts_listOutpostsWithS3Cmd.Flags().String("max-results", "", "The maximum number of Outposts to return.")
	s3outposts_listOutpostsWithS3Cmd.Flags().String("next-token", "", "When you can get additional results from the `ListOutpostsWithS3` call, a `NextToken` parameter is returned in the output.")
	s3outpostsCmd.AddCommand(s3outposts_listOutpostsWithS3Cmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2Cmd = &cobra.Command{
	Use:   "sesv2",
	Short: "Amazon SES API v2\n\n[Amazon SES](http://aws.amazon.com/ses) is an Amazon Web Services service that you can use to send email messages to your customers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2Cmd).Standalone()

	rootCmd.AddCommand(sesv2Cmd)
}

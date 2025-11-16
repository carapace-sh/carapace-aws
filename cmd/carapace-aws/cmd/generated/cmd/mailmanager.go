package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanagerCmd = &cobra.Command{
	Use:   "mailmanager",
	Short: "Amazon SES Mail Manager API\n\nThe Amazon SES Mail Manager API contains operations and data types that comprise the Mail Manager feature of [Amazon Simple Email Service (SES)](http://aws.amazon.com/ses).\n\nMail Manager is a set of Amazon SES email gateway features designed to help you strengthen your organization's email infrastructure, simplify email workflow management, and streamline email compliance control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanagerCmd).Standalone()

	rootCmd.AddCommand(mailmanagerCmd)
}

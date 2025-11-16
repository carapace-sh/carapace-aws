package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesCmd = &cobra.Command{
	Use:   "ses",
	Short: "Amazon Simple Email Service\n\nThis document contains reference information for the [Amazon Simple Email Service](https://aws.amazon.com/ses/) (Amazon SES) API, version 2010-12-01.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesCmd).Standalone()

	rootCmd.AddCommand(sesCmd)
}

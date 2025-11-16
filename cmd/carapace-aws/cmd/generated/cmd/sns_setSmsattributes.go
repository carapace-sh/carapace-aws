package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_setSmsattributesCmd = &cobra.Command{
	Use:   "set-smsattributes",
	Short: "Use this request to set the default settings for sending SMS messages and receiving daily SMS usage reports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_setSmsattributesCmd).Standalone()

	sns_setSmsattributesCmd.Flags().String("attributes", "", "The default settings for sending SMS messages from your Amazon Web Services account.")
	sns_setSmsattributesCmd.MarkFlagRequired("attributes")
	snsCmd.AddCommand(sns_setSmsattributesCmd)
}

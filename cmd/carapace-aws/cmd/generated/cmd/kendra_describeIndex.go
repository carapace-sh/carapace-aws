package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeIndexCmd = &cobra.Command{
	Use:   "describe-index",
	Short: "Gets information about an Amazon Kendra index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeIndexCmd).Standalone()

	kendra_describeIndexCmd.Flags().String("id", "", "The identifier of the index you want to get information on.")
	kendra_describeIndexCmd.MarkFlagRequired("id")
	kendraCmd.AddCommand(kendra_describeIndexCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listExperiencesCmd = &cobra.Command{
	Use:   "list-experiences",
	Short: "Lists one or more Amazon Kendra experiences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listExperiencesCmd).Standalone()

	kendra_listExperiencesCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
	kendra_listExperiencesCmd.Flags().String("max-results", "", "The maximum number of returned Amazon Kendra experiences.")
	kendra_listExperiencesCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_listExperiencesCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_listExperiencesCmd)
}

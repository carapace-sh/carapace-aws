package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listEntityPersonasCmd = &cobra.Command{
	Use:   "list-entity-personas",
	Short: "Lists specific permissions of users and groups with access to your Amazon Kendra experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listEntityPersonasCmd).Standalone()

	kendra_listEntityPersonasCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience.")
	kendra_listEntityPersonasCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
	kendra_listEntityPersonasCmd.Flags().String("max-results", "", "The maximum number of returned users or groups.")
	kendra_listEntityPersonasCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_listEntityPersonasCmd.MarkFlagRequired("id")
	kendra_listEntityPersonasCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_listEntityPersonasCmd)
}

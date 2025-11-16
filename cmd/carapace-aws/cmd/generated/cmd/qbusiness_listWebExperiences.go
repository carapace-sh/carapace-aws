package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listWebExperiencesCmd = &cobra.Command{
	Use:   "list-web-experiences",
	Short: "Lists one or more Amazon Q Business Web Experiences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listWebExperiencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listWebExperiencesCmd).Standalone()

		qbusiness_listWebExperiencesCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application linked to the listed web experiences.")
		qbusiness_listWebExperiencesCmd.Flags().String("max-results", "", "The maximum number of Amazon Q Business Web Experiences to return.")
		qbusiness_listWebExperiencesCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
		qbusiness_listWebExperiencesCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_listWebExperiencesCmd)
}

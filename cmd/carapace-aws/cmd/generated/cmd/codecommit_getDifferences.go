package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getDifferencesCmd = &cobra.Command{
	Use:   "get-differences",
	Short: "Returns information about the differences in a valid commit specifier (such as a branch, tag, HEAD, commit ID, or other fully qualified reference).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getDifferencesCmd).Standalone()

	codecommit_getDifferencesCmd.Flags().String("after-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit.")
	codecommit_getDifferencesCmd.Flags().String("after-path", "", "The file path in which to check differences.")
	codecommit_getDifferencesCmd.Flags().String("before-commit-specifier", "", "The branch, tag, HEAD, or other fully qualified reference used to identify a commit (for example, the full commit ID).")
	codecommit_getDifferencesCmd.Flags().String("before-path", "", "The file path in which to check for differences.")
	codecommit_getDifferencesCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codecommit_getDifferencesCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codecommit_getDifferencesCmd.Flags().String("repository-name", "", "The name of the repository where you want to get differences.")
	codecommit_getDifferencesCmd.MarkFlagRequired("after-commit-specifier")
	codecommit_getDifferencesCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getDifferencesCmd)
}

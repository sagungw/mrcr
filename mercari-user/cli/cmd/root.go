package cmd

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
)

func init() {
	RootCmd.AddCommand(ServerCmd)
}

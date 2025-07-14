package cmd

import (
	"fmt"
	"os"

	config "Backend-POS/configs"
	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "migrate",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return config.Open(cmd.Context())
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return config.Close(cmd.Context())
		},
		Run: func(cmd *cobra.Command, args []string) {
			migrateUp().Run(cmd, args)
		},
	}
	cmd.AddCommand(migrateUp())
	cmd.AddCommand(migrateDown())
	cmd.AddCommand(migrateRefresh())
	return cmd
}

func migrateUp() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "up",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			db := config.Database()
			if err := modelUp(db); err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			os.Exit(0)
		},
	}
	return cmd
}

func migrateDown() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "down",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			db := config.Database()
			if err := modelDown(db); err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			os.Exit(0)

		},
	}
	return cmd
}

func migrateRefresh() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "refresh",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			db := config.Database()
			if err := modelDown(db); err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			if err := modelUp(db); err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			os.Exit(0)

		},
	}
	return cmd
}

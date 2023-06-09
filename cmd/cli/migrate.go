package main

func doMigrate(arg2, arg3 string) error {
	//dsn := getDSN()
	checkForDB()

	tx, err := swg.PopConnect()
	if err != nil {
		exitGracefully(err)
	}
	defer tx.Close()

	// run the migration command
	switch arg2 {
	case "up":
		err := swg.RunPopMigrations(tx)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" {
			err := swg.PopMigrateDown(tx, -1)
			if err != nil {
				return err
			}
		} else {
			err := swg.PopMigrateDown(tx, 1)
			if err != nil {
				return err
			}
		}

	case "reset":
		err := swg.PopMigrateReset(tx)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}

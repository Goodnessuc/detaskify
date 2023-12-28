package db

func (d *Database) MigrateDatabase() error {
	err := d.Client.AutoMigrate(&User{}, &Follower{})
	if err != nil {
		return err
	}
	return nil
}

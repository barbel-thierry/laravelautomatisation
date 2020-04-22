package laravelreset

// CreateOrReset freshen the database up.
func CreateOrReset(dbContainer string) {
	artisanExists()
	createDb(getDbName(), dbContainer)
	migrate()
}
package laravelreset

// CreateOrReset freshen the database up.
func CreateOrReset(dbContainer string) {
	ArtisanExists()
	CreateDb(GetDbName(), dbContainer)
	Migrate()
}
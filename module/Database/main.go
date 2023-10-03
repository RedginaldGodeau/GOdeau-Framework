package Database

func Init(driver string, source string) Db {

	return Db{
		Driver: driver,
		Source: source,
	}
}
